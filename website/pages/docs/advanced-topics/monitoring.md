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

ELT workloads can be long running and sometimes it is necessary to better understand what calls are taking the most time; to potentially optimize those on the plugin side, ignore them or split them to a different workload.

CloudQuery supports [OpenTelemetry](https://opentelemetry.io/) tracing out of the box and can be enabled easily via [configuration](/docs/reference/source-spec).

To collect traces you need a collector that supports OpenTelemetry protocol, for example [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/). For example you can use [Jaeger](https://opentelemetry.io/docs/instrumentation/go/exporters/#jaeger) to visualize and analyze traces.

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
  version: "VERSION_SOURCE_AWS"
  tables: ["aws_s3_buckets"]
  destinations: ["postgresql"]
  otel_endpoint: "localhost:4318"
  otel_endpoint_insecure: true # this is only in development when running local jaeger
  spec:
```

After that you can open [http://localhost:16686](http://localhost:16686) and see the traces:

![jaeger](/images/docs/jaeger.png)


