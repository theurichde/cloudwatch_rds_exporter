# CloudWatch Prometheus Exporter

![Go Build & Test](https://github.com/theurichde/cloudwatch_rds_exporter/actions/workflows/go.yml/badge.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/percona/rds_exporter)](https://goreportcard.com/report/github.com/theurichde/cloudwatch_rds_exporter)

This project contains an [AWS RDS](https://aws.amazon.com/rds) exporter for [Prometheus](https://github.com/prometheus/prometheus).
It retrieves metrics from both [basic CloudWatch Metrics](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MonitoringOverview.html)
and [RDS Enhanced Monitoring via CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html).

* Based on [Technofy/cloudwatch_exporter](https://github.com/Technofy/cloudwatch_exporter) and [percona/rds_exporter](https://github.com/percona/rds_exporter).

## Quick Start

* Build the project from sources from the root folder

`go build -o cloudwatch_rds_exporter cmd/main.go`

* Create a configuration file `config.yml`:

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

* If `aws_access_key` and `aws_secret_key` are present, they are used for that instance.
Otherwise, the [default credential provider chain](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials)
is used, which includes `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables, `~/.aws/credentials` file,
and IAM role for EC2.


* Start the exporter by running:
    ```
    cloudwatch_rds_exporter --config.file="config.yaml"
    ```

* To see all flags run:
    ```
    cloudwatch_rds_exporter --help
    ```

## Run the Cloudwatch RDS Exporter via Docker
* Create your config.yaml
* Run the docker image
  * Mount your config.yaml
  * Either set your AWS Key and Secret via config.yaml (don't commit it anywhere) or via environment variables
    
```
$ docker run --rm -v $(pwd)/config.yaml:/opt/cloudwatch_rds_exporter/config.yaml theurichde/cloudwatch_rds_exporter
```

```
$ docker run --rm --env AWS_ACCESS_KEY_ID=myKeyID --env AWS_SECRET_ACCESS_KEY=mySecretAccessKey -v $(pwd)/config.yaml:/opt/cloudwatch_rds_exporter/config.yaml theurichde/cloudwatch_rds_exporter
```