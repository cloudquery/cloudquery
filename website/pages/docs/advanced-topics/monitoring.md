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

CloudQuery supports [OpenTelemetry](https://opentelemetry.io/) traces, metrics and logs out of the box and can be enabled easily via [configuration](/docs/reference/source-spec).

To collect OpenTelemetry data you need a [backend](https://opentelemetry.io/docs/concepts/components/#exporters) that supports OpenTelemetry protocol. For example you can use [Jaeger](https://opentelemetry.io/docs/instrumentation/go/exporters/#jaeger) to visualize and analyze traces.

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

## OpenTelemetry and Datadog

In this example we will show how to send OpenTelemetry traces, metrics and logs directly to Datadog.
First, you will need to [setup OpenTelemetry with Datadog](https://docs.datadoghq.com/opentelemetry/).
There are multiple ways to configure OpenTelemetry with Datadog. We'll show only a subset of them here, and you can find more information in the link above.

### Option 1: Using an OpenTelemetry collector

To config an OpenTelemetry collector with Datadog, you need to create a configuration file, for example `otel_collector_config.yaml` with the content below:

```yaml
receivers:
  otlp:
    protocols:
      http:
        endpoint: "0.0.0.0:4318"

processors:
  batch/datadog:
    send_batch_max_size: 1000
    send_batch_size: 100
    timeout: 10s

exporters:
  datadog:
    api:
      site: ${env:DATADOG_SITE}
      key: ${env:DATADOG_API_KEY}

service:
  pipelines:
    metrics:
      receivers: [otlp]
      processors: [batch/datadog]
      exporters: [datadog]
    traces:
      receivers: [otlp]
      processors: [batch/datadog]
      exporters: [datadog]
    logs:
      receivers: [otlp]
      processors: [batch/datadog]
      exporters: [datadog]
```

Then run the collector with the following command (replacing `DATADOG_SITE` and `DATADOG_API_KEY` with your own values):

```bash
docker run \
    -p 4318:4318 \
    -e DATADOG_SITE=$DATADOG_SITE \
    -e DATADOG_API_KEY=$DATADOG_API_KEY \
    --hostname $(hostname) \
    -v $(pwd)/otel_collector_config.yaml:/etc/otelcol-contrib/config.yaml \
    otel/opentelemetry-collector-contrib:0.104.0
```

> For additional ways to run the collector, please refer to the [official documentation](https://docs.datadoghq.com/opentelemetry/collector_exporter/deployment#running-the-collector).

### Option 2: Direct OTEL Ingestion by the Datadog Agent via a configuration file

[Locate](https://docs.datadoghq.com/agent/configuration/agent-configuration-files/) your `datadog.yaml` file and add the following configuration:

```yaml
otlp_config:
  receiver:
    protocols:
      http:
        endpoint: 0.0.0.0:4318
  logs:
    enabled: true
logs_enabled: true
```

[Restart](https://docs.datadoghq.com/agent/configuration/agent-commands/#restart-the-agent) the Datadog agent for the change to take effect.

### Option 3: Direct OTEL ingestion by the Datadog Agent via environment variables

Pass the `DD_OTLP_CONFIG_RECEIVER_PROTOCOLS_HTTP_ENDPOINT` environment variable to the Datadog agent with a value of `0.0.0.0:4318`.
If you're using Docker compose, you can find an example below:

```yaml
version: "3.0"
services:
  agent:
    image: gcr.io/datadoghq/agent:7
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - /proc/:/host/proc/:ro
      - /sys/fs/cgroup/:/host/sys/fs/cgroup:ro
    environment:
      DD_API_KEY: redacted
      DD_SITE: "datadoghq.eu"
      DD_OTLP_CONFIG_RECEIVER_PROTOCOLS_HTTP_ENDPOINT: "0.0.0.0:4318"
      DD_LOGS_ENABLED: "true"
      DD_OTLP_CONFIG_LOGS_ENABLED: "true"
    ports:
      - "4318:4318"
```

[Restart](https://docs.datadoghq.com/agent/configuration/agent-commands/#restart-the-agent) the Datadog agent for the change to take effect.

> For additional ways to configure the Datadog agent, please refer to the [official documentation](https://docs.datadoghq.com/opentelemetry/interoperability/otlp_ingest_in_the_agent#enabling-otlp-ingestion-on-the-datadog-agent).

### Start CloudQuery Configured with Datadog

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

We also provide a Datadog dashboard you can download from [here](/assets/datadog-dashboard.json) and import it into your Datadog account:
1. Click "New Dashboard"
2. In the name field, type "CloudQuery Sync Dashboard", then click "New Dashboard"
3. Click "Configure" -> "Import dashboard JSON…"
4. Drag the JSON file into the window, or copy-paste the contents.

![Datadog](/images/docs/monitoring/cq_otel_datadog_dashboard.png)
