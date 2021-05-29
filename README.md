# CloudWatch Prometheus Exporter

![Go Build](https://github.com/theurichde/cloudwatch_rds_exporter/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/percona/rds_exporter)](https://goreportcard.com/report/github.com/theurichde/cloudwatch_rds_exporter)

An [AWS RDS](https://aws.amazon.com/ru/rds/) exporter for [Prometheus](https://github.com/prometheus/prometheus).
It gets metrics from both [basic CloudWatch Metrics](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MonitoringOverview.html)
and [RDS Enhanced Monitoring via CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html).

Based on [Technofy/cloudwatch_exporter](https://github.com/Technofy/cloudwatch_exporter) and [percona/rds_exporter](https://github.com/percona/rds_exporter).

## Quick start

Create a configuration file `config.yml`:

```yaml
---
instances:
  - instance: rds-aurora1
    region: us-east-1
  - instance: rds-mysql57
    region: us-east-1
    aws_access_key: AKIAIOSFODNN7EXAMPLE
    aws_secret_key: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

If `aws_access_key` and `aws_secret_key` are present, they are used for that instance.
Otherwise, the [default credential provider chain](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials)
is used, which includes `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables, `~/.aws/credentials` file,
and IAM role for EC2.


Start the exporter by running:
```
cloudwatch_rds_exporter --config.file="config.yaml"
```

To see all flags run:
```
cloudwatch_rds_exporter --help
```

Configure Prometheus:

```yaml
---
scrape_configs:
  - job_name: rds-basic
    scrape_interval: 60s
    scrape_timeout: 55s
    metrics_path: /basic
    honor_labels: true
    static_configs:
      - targets:
        - 127.0.0.1:9042

  - job_name: rds-enhanced
    scrape_interval: 10s
    scrape_timeout: 9s
    metrics_path: /enhanced
    honor_labels: true
    static_configs:
      - targets:
        - 127.0.0.1:9042
```

`honor_labels: true` is important because the exporter returns metrics with `instance` label set.

