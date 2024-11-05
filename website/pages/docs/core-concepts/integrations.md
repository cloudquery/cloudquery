---
title: Integrations
description: "CloudQuery has a integration-based architecture and relies on two types of integrations: source integrations and destination integrations."
---

# Integrations

CloudQuery has a integration-based architecture, with integrations communicating over [gRPC](https://github.com/cloudquery/plugin-pb). An integration can be implemented to be a source, destination or both.

- **Source integration** - Responsible for extracting and transforming configuration from cloud-providers, SaaS apps, and other APIs ([Available source integrations](https://hub.cloudquery.io/plugins/source)).
- **Destination integration** - Responsible for writing the data from the source integrations to various destinations such as databases, message queues and storage ([Available destination integrations](https://hub.cloudquery.io/plugins/destination)).

All integrations are split to official (maintained by CloudQuery) and community (maintained by members of the community in their own repositories).

## Source Integration

The core responsibilities of a source integration:

- Define the schema (tables).
- Authenticate with the supported APIs, SaaS services and/or cloud providers.
- Extracting data from the supported APIs and transform them into the defined schema.
- Send the data via [protobuf](https://github.com/cloudquery/plugin-sdk/tree/main/internal/pb) to the CLI for further processing and storage at the defined destination integrations.

See [Configuration Reference](../reference/source-spec)

## Destination Integration

The core responsibilities of a destination integration:

- Authenticate with the destination (such as database, message queue, storage).
- Auto-migrate the schemas defined by the source integrations.
- Save each incoming object in the appropriate table.

See [Configuration Reference](../reference/destination-spec)
