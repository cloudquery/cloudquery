# Providers

CloudQuery has a pluggable architecture and relies on plugins called "providers" which are responsible for extracting and transforming configuration from cloud-providers, SaaS apps, and other APIs.

> If you are familiar with Terraform, we use the [same pluggable architecture](https://www.terraform.io/language/providers). CloudQuery providers focus on extracting & transforming configuration from infrastructure, instead of provisioning it.

## Capabilities

The core responsibility of a provider is to define resources and tables.

Every resource can define one or more tables this resource configuration is transformed to.

## Hosting

Providers are hosted on GitHub while [hub.cloudquery.io](https://hub.cloudquery.io) serves as a central registry for discovery, validation & verification.

Currently CloudQuery Hub contains two types of providers:

- **official** - Owned and maintained by CloudQuery team
- **community** - Owned and maintained by either individuals or third-party vendors.
