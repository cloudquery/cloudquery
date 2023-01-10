# Plugins

CloudQuery has a pluggable architecture and relies on two types of plugins:

- **Source plugin** - Responsible for extracting and transforming configuration from cloud-providers, SaaS apps, and other APIs ([Available source plugins](/docs/plugins/sources/overview)).
- **Destination plugin** - Responsible for writing the data from the source plugins to various destinations such as databases, message queues and storage ([Available destination plugins](/docs/plugins/destinations/overview)).

All plugins are split to official (maintained by CloudQuery) and community (maintained by members of the community in their own repositories).

## Source Plugin

The core responsibilities of a source plugin:

- Define the schema (tables).
- Authenticate with the supported APIs, SaaS services and/or cloud providers.
- Extracting data from the supported APIs and transform them into the defined schema.
- Send the data as JSON to the CLI for further processing and storage at the defined destination plugins.

See [Configuration Reference](../reference/source-spec)

## Destination Plugin

The core responsibilities of a destination plugin:

- Authenticate with the destination (such as database, message queue, storage).
- Auto-migrate the schemas defined by the source plugins.
- Save each incoming JSON object in the appropriate table.

See [Configuration Reference](../reference/destination-spec)
