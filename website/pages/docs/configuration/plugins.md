# Plugins

CloudQuery has a pluggable architecture and relies on two types of plugins:

- Source plugins - which are responsible for extracting and transforming configuration from cloud-providers, SaaS apps, and other APIs.
- Destination plugins - which are responsible for writing the data from the source plugins to various destinations suchs as databases, message queues and storage.

## Capabilities

The core responsibility of a provider is to define resources and tables.

Every resource can define one or more tables this resource configuration is transformed to.

## Hosting

Providers are hosted on GitHub while [hub.cloudquery.io](https://hub.cloudquery.io) serves as a central registry for discovery, validation & verification.

Currently CloudQuery Hub contains two types of providers:

- **official** - Owned and maintained by CloudQuery team
- **community** - Owned and maintained by either individuals or third-party vendors.
