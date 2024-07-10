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

### OpenTelemetry and Datadog

In this example we will show how to send OpenTelemetry traces from the CLI directly to Datadog.

First, you will need to [setup OpenTelemetry in Datadog](https://docs.datadoghq.com/opentelemetry/).
You can chose either to send data directly to a Datadog agent or use the OpenTelemetry collector, follow the instructions in the link above and chose what's best for you.

Once you have the agent or collector ready, you can specify the endpoint in the source spec:

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

Once ingestion starts you should be able to start seeing the traces in Datadog under APM->Traces->Explorer.

![Datadog](/images/docs/monitoring/cq_otel_datadog.png)

We also provide a Datadog dashboard you can download from [here](/assets/datadog-dashboard.json) and import it into your Datadog account.

![Datadog](/images/docs/monitoring/cq_otel_datadog_dashboard.png)
