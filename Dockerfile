# syntax=docker/dockerfile:1
FROM golang:1.16 AS builder
WORKDIR /go/src/github.com/theurichde/cloudwatch_rds_exporter
COPY . .
RUN go build

FROM golang:1.16-buster AS app
RUN apt-get install -y ca-certificates
RUN groupadd -r cloudwatch_rds_exporter --gid 1000 && useradd --uid 1000 -g cloudwatch_rds_exporter -r -s /bin/bash cloudwatch_rds_exporter
WORKDIR /opt/cloudwatch_rds_exporter
COPY --from=builder /go/src/github.com/theurichde/cloudwatch_rds_exporter/cloudwatch_rds_exporter /usr/local/bin/cloudwatch_rds_exporter
RUN chown cloudwatch_rds_exporter:cloudwatch_rds_exporter /usr/local/bin/cloudwatch_rds_exporter

USER cloudwatch_rds_exporter

EXPOSE 9042
CMD ["cloudwatch_rds_exporter", "--config.file", "/opt/cloudwatch_rds_exporter/config.yaml"]
