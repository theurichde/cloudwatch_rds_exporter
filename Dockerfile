# syntax=docker/dockerfile:1
FROM golang:1.16 AS builder
WORKDIR /go/src/github.com/hellofresh/rds_exporter
COPY . .
RUN go build

FROM golang:1.16-buster AS app
RUN apt-get install -y ca-certificates
RUN groupadd -r rds_exporter --gid 1000 && useradd --uid 1000 -g rds_exporter -r -s /bin/bash rds_exporter
WORKDIR /opt/rds_exporter
COPY --from=builder /go/src/github.com/hellofresh/rds_exporter/rds_exporter /usr/local/bin/rds_exporter
RUN chown rds_exporter:rds_exporter /usr/local/bin/rds_exporter

USER rds_exporter

EXPOSE 9042
CMD ["rds_exporter", "--config.file", "/opt/rds_exporter/config.yaml"]
