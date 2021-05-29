package basic

import (
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/theurichde/cloudwatch_rds_exporter/internal/config"
)

var (
	// Period is period in seconds
	Period = 60 * time.Second
	// Delay is delay in seconds
	Delay = 600 * time.Second
	// Range is range in seconds
	Range = 600 * time.Second
)

// Scraper is an entity used to describe a scraper
type Scraper struct {
	// params
	instance *config.Instance
	exporter *Exporter
	ch       chan<- prometheus.Metric

	// internal
	svc    *cloudwatch.CloudWatch
	labels []string
}

// NewScraper is a function used to initialize a scraper
func NewScraper(instance *config.Instance, exporter *Exporter, ch chan<- prometheus.Metric) *Scraper {
	// Create CloudWatch client
	sess, _ := exporter.sessions.GetSession(instance.Region, instance.Instance)
	svc := cloudwatch.New(sess)

	// Create labels for all metrics
	labels := []string{
		instance.Instance,
		instance.Region,
	}

	return &Scraper{
		// params
		instance: instance,
		exporter: exporter,
		ch:       ch,

		// internal
		svc:    svc,
		labels: labels,
	}
}

func getLatestDatapoint(datapoints []*cloudwatch.Datapoint) *cloudwatch.Datapoint {
	var latest *cloudwatch.Datapoint

	for dp := range datapoints {
		if latest == nil || latest.Timestamp.Before(*datapoints[dp].Timestamp) {
			latest = datapoints[dp]
		}
	}

	return latest
}

// Scrape makes the required calls to AWS CloudWatch by using the parameters in the Exporter.
// Once converted into Prometheus format, the metrics are pushed on the ch channel.
func (s *Scraper) Scrape() {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	wg.Add(len(s.exporter.metrics))
	for _, metric := range s.exporter.metrics {
		go func(metric Metric) {
			err := s.scrapeMetric(metric)
			if err != nil {
				s.exporter.l.With("metric", metric.Name).Error(err)
			}

			wg.Done()
		}(metric)
	}
}

func (s *Scraper) scrapeMetric(metric Metric) error {
	now := time.Now()
	end := now.Add(-Delay)

	params := &cloudwatch.GetMetricStatisticsInput{
		EndTime:   aws.Time(end),
		StartTime: aws.Time(end.Add(-Range)),

		Period:     aws.Int64(int64(Period.Seconds())),
		MetricName: aws.String(metric.Name),
		Namespace:  aws.String("AWS/RDS"),
		Dimensions: []*cloudwatch.Dimension{},
		Statistics: aws.StringSlice([]string{"Average"}),
		Unit:       nil,
	}

	params.Dimensions = append(params.Dimensions, &cloudwatch.Dimension{
		Name:  aws.String("DBInstanceIdentifier"),
		Value: aws.String(s.instance.Instance),
	})

	// Call CloudWatch to gather the datapoints
	resp, err := s.svc.GetMetricStatistics(params)
	if err != nil {
		return err
	}

	// There's nothing in there, don't publish the metric
	if len(resp.Datapoints) == 0 {
		return nil
	}

	// Pick the latest datapoint
	dp := getLatestDatapoint(resp.Datapoints)

	// Get the metric.
	v := aws.Float64Value(dp.Average)
	switch metric.Name {
	case "EngineUptime":
		// "Fake EngineUptime -> node_boot_time with time.Now().Unix() - EngineUptime."
		v = float64(time.Now().Unix() - int64(v))
	}

	// Send metric.
	s.ch <- prometheus.MustNewConstMetric(metric.Desc, prometheus.GaugeValue, v, s.labels...)

	return nil
}
