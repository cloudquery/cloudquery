---
title: Monitoring CloudQuery
description: Short walkthrough on how to use OpenTelemetry to monitor CloudQuery Syncs
---

# Monitoring

Monitoring CloudQuery can be done in a number of main ways:

- Logging
- OpenTelemetry

## Logging

CloudQuery utilizes structured [logging](../reference/cli/cloudquery) (in plain and JSON formats) which can be analyzed by local tools such as `jq`, `grep` and remote aggregations tools like `loki`, `datadog` or any other popular log aggregation that supports structured logging.

## OpenTelemetry (Preview)

ELT workloads can be long running and sometimes it is necessary to better understand what calls are taking the most time; to potentially optimize those on the plugin side, ignore them or split them to a different workload. Plugins come with an OpenTelemetry library built in, but it is up to the plugin author to instrument the most important parts--usually the API calls--this way it is possible to see what calls take the longest time, or where throttling and errors are happening.

CloudQuery supports [OpenTelemetry](https://opentelemetry.io/) tracing out of the box and can be enabled easily via [configuration](/docs/reference/source-spec).

To collect traces you need a [backend](https://opentelemetry.io/docs/concepts/components/#exporters) that supports OpenTelemetry protocol. For example you can use [Jaeger](https://opentelemetry.io/docs/instrumentation/go/exporters/#jaeger) to visualize and analyze traces.

To start Jaeger locally you can use Docker:

```bash
docker run -d \
  -e COLLECTOR_OTLP_ENABLED=true \
  -p 16686:16686 \
  -p 4318:4318 \
  jaegertracing/all-in-one:latest
```

and then specify in the source spec the following:

```yaml
kind: source
spec:
  name: "aws"
  path: "cloudquery/aws"
  registry: "cloudquery"
  version: "VERSION_SOURCE_AWS"
  tables: ["aws_s3_buckets"]
  destinations: ["postgresql"]
  otel_endpoint: "localhost:4318"
  otel_endpoint_insecure: true # this is only in development when running local jaeger
  spec:
```

After that you can open [http://localhost:16686](http://localhost:16686) and see the traces:

![jaeger](/images/docs/jaeger.png)

In production, it is common to use an OpenTelemetry [collector](https://opentelemetry.io/docs/concepts/components/#collector) that runs locally or as a gateway to batch the traces and forward it to the final backend. This helps with performance, fault-tolerance and decoupling of the backend in case the tracing backend changes.

### OpenTelemetry and Datadog

In this quick example we will show how to connect an open telemetry collector to Datadog via OpenTelemetry exporter.

Firstly, you will need to have an [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/) running either locally or as a gateway. Here is an example of running it locally with docker:

```bash
docker run  -p 4319:4319 -v $(pwd)/config.yml:/etc/otelcol-contrib/config.yaml otel/opentelemetry-collector-contrib:0.91.0
```

following is an example for OTEL collector `config.yml` to receive traces locally on 4318 and export them to Datadog:

```yaml
receivers:
  otlp:
    protocols:
      http:
        endpoint: 0.0.0.0:4318
exporters:
  datadog:
    api:
      site: "datadoghq.com" # or your tenant site https://docs.datadoghq.com/getting_started/site/
      key: "<DATADOG_API_KEY>"
```

Once ingestion starts you should be able to start seeing the traces in Datadog under ServiceCatalog and Traces with ability to view average p95 latency, error rate, total duration and other useful information you can query to either split the workload better or improve the plugin scheduling if you are the plugin author:

![Datadog](/images/docs/monitoring/cq_otel_datadog.png)

![Datadog](/images/docs/monitoring/cq_otel_datadog_traces.png)

