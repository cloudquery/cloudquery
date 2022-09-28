# Plugins

CloudQuery has a pluggable architecture and relies on two types of plugins:

- **Source plugin** - Responsible for extracting and transforming configuration from cloud-providers, SaaS apps, and other APIs ([Available source plugins](../plugins/sources)).
- **Destination plugin** - Responsible for writing the data from the source plugins to various destinations such as databases, message queues and storage (Full list available [here](../plugins/destinations)).

All plugins are split to official (maintained by CloudQuery) and community (maintained by memebers of the community in their own repositories).

## Source Plugin

The core responsibilities of a source plugin:

- Define the schema (tables).
- Authenticate with the supported APIs, SaaS services and/or cloud providers.
- Extracting data from the supported APIs and transform them into the defined schema.
- Send the data as JSON to the CLI for further processing and storage at the defined destination plugins.

### Configuration

Following is all available configuration for source plugins:

```yaml
kind: source
spec:
  ## Required. name of the plugin (should match any name available in list of source plugins)
  name: SOURCE_PLUGIN_NAME

  ## Optional. Default: latest. It is highly recommended to pin to a specific version in production.
  # version: latest

  ## Optional. Default: github. Available: local, grpc.
  ## By default it will search for plugins hosted on github.
  # registry: github

  ## Optional. Default: "cloudquery/SOURCE_PLUGIN_NAME".
  ## For community plugin name and path will be the same i.e: ""
  # path: cloudquery/SOURCE_PLUGIN_NAME
```

## Destination Plugin

The core responsibilities of a destination plugin:

- Authenticate with the destination (such as database, message queue, storage).
- Auto-migrate the schemas defined by the source plugins.
- Save each incoming json object in the appropriate table.
