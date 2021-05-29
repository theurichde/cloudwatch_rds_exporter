package enhanced

import (
	"context"
	"github.com/theurichde/cloudwatch_rds_exporter/internal/sessions"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

// Collector collects enhanced RDS metrics by utilizing several scrapers.
type Collector struct {
	sessions *sessions.Sessions
	logger   log.Logger

	rw      sync.RWMutex
	metrics map[string][]prometheus.Metric
}

// Maximal and minimal metrics update interval.
const (
	maxInterval = 30 * time.Second
	minInterval = 10 * time.Second
)

// NewCollector creates new collector and starts scrapers.
func NewCollector(sessions *sessions.Sessions) *Collector {
	c := &Collector{
		sessions: sessions,
		logger:   log.With("component", "enhanced"),
		metrics:  make(map[string][]prometheus.Metric),
	}

	for session, instances := range sessions.AllSessions() {
		s := newScraper(session, instances)

		interval := maxInterval
		for _, instance := range instances {
			if instance.EnhancedMonitoringInterval > 0 && instance.EnhancedMonitoringInterval < interval {
				interval = instance.EnhancedMonitoringInterval
			}
		}
		if interval < minInterval {
			interval = minInterval
		}
		s.logger.Infof("Updating enhanced metrics every %s.", interval)

		// perform first scrapes synchronously so returned collector has all metric descriptions
		c.setMetrics(s.scrape(context.TODO()))

		ch := make(chan map[string][]prometheus.Metric)
		go func() {
			for m := range ch {
				c.setMetrics(m)
			}
		}()
		go s.start(context.TODO(), interval, ch)
	}

	return c
}

// setMetrics saves latest scraped metrics.
func (c *Collector) setMetrics(m map[string][]prometheus.Metric) {
	c.rw.Lock()
	for id, metrics := range m {
		c.metrics[id] = metrics
	}
	c.rw.Unlock()
}

// Describe implements prometheus.Collector.
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	for _, metrics := range c.metrics {
		for _, m := range metrics {
			ch <- m.Desc()
		}
	}
}

// Collect implements prometheus.Collector.
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	for _, metrics := range c.metrics {
		for _, m := range metrics {
			ch <- m
		}
	}
}

// check interfaces
var (
	_ prometheus.Collector = (*Collector)(nil)
)
