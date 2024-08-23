---
title: OpenTelemetry and Grafana
description: Walkthrough on how to use OpenTelemetry to monitor CloudQuery Syncs with Grafana
---

# OpenTelemetry and Grafana

In this example we will show how to visualize [CloudQuery][CloudQuery] [OpenTelemetry][OpenTelemetry] traces, metrics and logs with [Grafana][Grafana].
We will use [Docker Compose][Compose] to run [Grafana][Grafana] and related services, so make sure you it installed on your machine.

## Step 1: Creating a Docker Compose file

We will use [Tempo][Tempo] for ingesting traces, [Loki][Loki] for logs, [Prometheus][Prometheus] for metrics, and the [OpenTelemetry collector][Collector] for collecting and forwarding the data to each service.

Create a file named `docker-compose.yml` with the following content:

```yaml
version: "3.8"
services:
  tempo:
    image: grafana/tempo:latest
    command: ["-config.file=/etc/tempo.yaml"]
    volumes:
      - tempo_data:/tmp
      - ./tempo/tempo.yaml:/etc/tempo.yaml
    ports:
      - "3200"
      - "4318"
  loki:
    image: grafana/loki:latest
    ports:
      - "3100"
    command: -config.file=/etc/loki/local-config.yaml
  collector:
    image: otel/opentelemetry-collector-contrib:latest
    ports:
      - "4318:4318" # 4318 needs to be exposed to the host for the collector to ingest data
      - "8090"
    volumes:
      - ./collector/collector.yaml:/etc/otelcol-contrib/config.yaml
  prometheus:
    image: prom/prometheus:latest
    command:
      - "--enable-feature=remote-write-receiver"
      - "--config.file=/etc/prometheus/prometheus.yaml"
    ports:
      - "9090"
    volumes:
      - prometheus:/prometheus
      - ./prometheus/prometheus.yaml:/etc/prometheus/prometheus.yaml
  grafana:
    image: grafana/grafana-enterprise
    volumes:
      - grafana_data:/var/lib/grafana
      - ./grafana/datasources.yaml:/etc/grafana/provisioning/datasources/datasources.yaml
      - ./grafana/dashboards.yaml:/etc/grafana/provisioning/dashboards/dashboards.yaml
      - ./grafana/cloudquery-dashboard.json:/var/lib/grafana/dashboards/cloudquery-dashboard.json
    environment:
      GF_FEATURE_TOGGLES_ENABLE: "tempoApmTable"
    ports:
      - "3000:3000" # 3000 needs to be exposed to the host for the Grafana UI
volumes:
  prometheus:
    driver: local
  grafana_data:
    driver: local
  tempo_data:
    driver: local
```

This Docker Compose file configures [Prometheus][Prometheus], [Tempo][Tempo], an [OpenTelemetry collector][Collector] and Grafana with a custom configuration, and [Loki][Loki] with the default configuration.

## Step 2: Configure Prometheus

Create a file with the path `prometheus/prometheus.yaml` with the following content:

```yaml
global:
  scrape_interval: 15s
scrape_configs:
  - job_name: "opentelemetry"
    static_configs:
      - targets: ["collector:8090"]
```

This configuration will tell [Prometheus][Prometheus] to scrape the [OpenTelemetry][OpenTelemetry] collector every 15 seconds.

## Step 3: Configure Tempo

Create a file with the path `tempo/tempo.yaml` with the following content:

```yaml
server:
  http_listen_port: 3200
distributor:
  receivers:
    otlp:
      protocols:
        http:
storage:
  trace:
    backend: local
    wal:
      path: /tmp/tempo/wal
    local:
      path: /tmp/tempo/blocks
# Needed for aggregation functions, e.g. quantile_over_time
# Visit https://grafana.com/docs/tempo/latest/traceql/metrics-queries/ for more information
query_frontend:
  search:
    max_duration: 0
  metrics:
    max_duration: 0
overrides:
  metrics_generator_processors: ["local-blocks"]
metrics_generator:
  processor:
    local_blocks:
      filter_server_spans: false
  storage:
    path: /var/tempo/generator/wal
  traces_storage:
    path: /var/tempo/generator/traces
```

This configuration will tell [Tempo][Tempo] to listen on port 3200 and receive [OpenTelemetry][OpenTelemetry] traces via HTTP on the default port of 4318.

## Step 4: Configure the OpenTelemetry collector

Create a file with the path `collector/collector.yaml` with the following content:

```yaml
receivers:
  otlp:
    protocols:
      http:
        endpoint: "0.0.0.0:4318"
processors:
  batch:
exporters:
  prometheus:
    endpoint: collector:8090
  otlphttp:
    endpoint: http://tempo:4318
  loki:
    endpoint: http://loki:3100/loki/api/v1/push
service:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlphttp]
    metrics:
      receivers: [otlp]
      processors: [batch]
      exporters: [prometheus]
    logs:
      receivers: [otlp]
      exporters: [loki]
```

This configuration will tell the [OpenTelemetry collector][Collector] to receive traces, metrics, and logs and forward them to [Tempo][Tempo], [Prometheus][Prometheus] and [Loki][Loki], respectively.

## Step 5: Configure Grafana Data Sources

Create a file with the path `grafana/datasources.yaml` with the following content:

```yaml
apiVersion: 1
datasources:
  - name: Prometheus
    type: prometheus
    access: proxy
    orgId: 1
    url: http://prometheus:9090
    basicAuth: false
    isDefault: false
    version: 1
    editable: true
    uid: prometheus
  - name: Loki
    type: loki
    access: proxy
    orgId: 1
    url: http://loki:3100
    basicAuth: false
    isDefault: false
    version: 1
    editable: true
    uid: loki
  - name: Tempo
    type: tempo
    access: proxy
    orgId: 1
    url: http://tempo:3200
    basicAuth: false
    isDefault: true
    version: 1
    editable: true
    apiVersion: 1
    uid: tempo
```

This configuration will tell [Grafana][Grafana] to use [Prometheus][Prometheus], [Loki][Loki], and [Tempo][Tempo] as data sources.

## Step 6: Download the CloudQuery Grafana Dashboard

Create a file with the path `grafana/cloudquery-dashboard.json` with the content from [here](/assets/grafana-dashboard.json).

> If you'd like to import the dashboard to an existing Grafana instance, you can download an external version of it from [here](/assets/grafana-dashboard-external.json).

## Step 7: Configure Grafana with the CloudQuery Dashboard

Create a file with the path `grafana/dashboards.yaml` with the following content:

```yaml
apiVersion: 1

providers:
  - name: CloudQuery
    folder: CloudQuery
    type: file
    allowUiUpdates: true
    options:
      path: /var/lib/grafana/dashboards
```

This configuration will tell [Grafana][Grafana] to load the [CloudQuery][CloudQuery] dashboard.

## Step 8: Start the services

Run `docker-compose up` to start the services. Once the services are up and running, you should be able to access [Grafana][Grafana] at [http://localhost:3000](http://localhost:3000) with the default credentials `admin:admin`.

## Step 9: Configure a Source Plugin with OpenTelemetry

You can use the example source configuration below to start a sync with [OpenTelemetry][OpenTelemetry] enabled:

```yaml
kind: source
spec:
  name: "aws"
  path: "cloudquery/aws"
  registry: "cloudquery"
  version: "VERSION_SOURCE_AWS"
  tables: ["aws_s3_buckets"]
  destinations: ["postgresql"]
  otel_endpoint: "0.0.0.0:4318"
  otel_endpoint_insecure: true
  spec:
```

## Step 10: Run the sync

Run `cloudquery sync spec.yml --log-level debug`.

:::callout{type="info"}
Running with `--log-level debug` is recommended to get more detailed logs about requests retries and errors.
:::

After ingestion starts, you can access the [dashboard](http://localhost:3000/d/6_bNYpGVz/cloudquery-dashboard?orgId=1) to see sync insights, traces, metrics, and logs.


[CloudQuery]: https://www.cloudquery.io/
[OpenTelemetry]: https://opentelemetry.io/
[Compose]: https://docs.docker.com/compose/install/
[Grafana]: https://grafana.com/
[Tempo]: https://grafana.com/oss/tempo/
[Loki]: https://grafana.com/oss/loki/
[Prometheus]: https://prometheus.io/
[Collector]: https://opentelemetry.io/docs/collector
