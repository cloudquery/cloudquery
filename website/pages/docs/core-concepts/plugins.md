# Plugins

CloudQuery has a pluggable architecture and relies on two types of plugins:

- **Source plugins** - Responsible for extracting and transforming configuration from cloud-providers, SaaS apps, and other APIs (Full List available [here](../plugins/sources.md)).
- **Destination plugins** - Responsible for writing the data from the source plugins to various destinations suchs as databases, message queues and storage (Full list available [here](../plugins/destinations.md)).

All plugins are split to official (maintained by CloudQuery) and community (maintained by memebers of the community in their own repositories).

## Source Plugin

The core responsibilities of a source plugin:

- Define the schema (tables).
- Authenticate with the supported APIs, SaaS services and/or cloud providers.
- Extract all data from the supported APIs and transform them into the defined schema.
- Send the data as JSON to the CLI for further processing and storage at the defined destination plugins.

### Configuration

Following is all available configuration for source plugins:

```yaml
kind: source
spec:
  ## Required. name of the plugin (should match any name available in list of source plugins)
  name: SOURCE_PLUGIN_NAME

  ## Optional. Default: latest. It is highly recommended to pin to a specific version in production.
  # verison: latest

  ## Optional. Default: github. Available: local, grpc.
  ## By default it will search for plugins hosted on github.
  # registry: github

  ## Optional. Default: "cloudquery/SOURCE_PLUGIN_NAME".
  ## For community plugin name and path will be the same i.e: ""
  # path: cloudquery/SOURCE_PLUGIN_NAME
```

## Destination Plugin

The core responsabilities of a destination plugin:

- Authenticate with the destination (such as database, message queue, storage).
- Auto-migrate the schemas defined by the source plugins.
- Save each incoming json object in the appropriate table.
