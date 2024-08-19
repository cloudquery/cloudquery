---
title: OpenTelemetry and Datadog
description: Walkthrough on how to use OpenTelemetry to monitor CloudQuery Syncs with Datadog
---

# OpenTelemetry and Datadog

In this example we will show how to send OpenTelemetry traces, metrics and logs directly to Datadog.
First, you will need to [setup OpenTelemetry with Datadog](https://docs.datadoghq.com/opentelemetry/).
There are multiple ways to configure OpenTelemetry with Datadog. We'll show only a subset of them here, and you can find more information in the link above.

## Option 1: Using an OpenTelemetry collector

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

## Option 2: Direct OTEL Ingestion by the Datadog Agent via a configuration file

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

## Option 3: Direct OTEL ingestion by the Datadog Agent via environment variables

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

## Start CloudQuery Configured with Datadog

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

Run `cloudquery sync spec.yml --log-level debug`.

:::callout{type="info"}
Running with `--log-level debug` is recommended to get more detailed logs about requests retries and errors.
:::

After ingestion starts, you should start seeing traces in the Datadog [**APM Traces Explorer**](https://app.datadoghq.com/apm/traces).
You can also validate metrics and logs in the [**Metrics Summary**](https://app.datadoghq.com/metric/summary) and [**Log Explorer**](https://app.datadoghq.com/logs).

![Datadog](/images/docs/monitoring/cq_otel_datadog.png)

We also provide a Datadog dashboard you can download from [here](/assets/datadog-dashboard.json) and import it into your Datadog account:
1. Click "New Dashboard"
2. In the name field, type "CloudQuery Sync Dashboard", then click "New Dashboard"
3. Click "Configure" -> "Import dashboard JSON…"
4. Drag the JSON file into the window, or copy-paste the contents.

![Datadog](/images/docs/monitoring/cq_otel_datadog_dashboard.png)