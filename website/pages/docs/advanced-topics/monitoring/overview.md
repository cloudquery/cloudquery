---
title: Overview
description: Walkthrough on how to use OpenTelemetry to monitor CloudQuery Syncs
---

# Overview

Monitoring CloudQuery can be done in a number of ways:

- [Logging](#logging)
- [OpenTelemetry](#opentelemetry-preview)
- [Datadog](/docs/advanced-topics/monitoring/otel-datadog)
- [Grafana](/docs/advanced-topics/monitoring/otel-grafana)

## Logging

CloudQuery utilizes structured [logging](/docs/reference/cli/cloudquery) (in plain and JSON formats) which can be analyzed by local tools such as `jq`, `grep` and remote aggregations tools like `loki`, `datadog` or any other popular log aggregation that supports structured logging.

## OpenTelemetry (Preview)

ELT workloads can be long running and sometimes it is necessary to better understand what calls are taking the most time, to optimize those on the integration side, ignore them or split them to a different workload.
CloudQuery supports [OpenTelemetry](https://opentelemetry.io/) traces, metrics and logs out of the box and can be enabled easily via [configuration](/docs/reference/source-spec).

To collect OpenTelemetry data you need a [backend](https://opentelemetry.io/docs/concepts/components/#exporters) that supports the OpenTelemetry protocol. For example you can use [Jaeger](https://opentelemetry.io/docs/instrumentation/go/exporters/#jaeger) to visualize and analyze traces.

To start Jaeger locally you can use Docker:

```bash
docker run -d \
  -e COLLECTOR_OTLP_ENABLED=true \
  -p 16686:16686 \
  -p 4318:4318 \
  jaegertracing/all-in-one:1.58
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
