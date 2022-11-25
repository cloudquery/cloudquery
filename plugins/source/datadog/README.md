# Datadog Plugin

The CloudQuery DataDog plugin extracts your Datadog information and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Links

- [Tables](./docs/tables/README.md)

## Authentication

CloudQuery requires only `DD_CLIENT_API_KEY` and `DD_CLIENT_APP_KEY`. Follow [this guide](https://docs.datadoghq.com/account_management/api-app-keys/) for how to create an API key and app key for CloudQuery.

CloudQuery requires only *read* permissions (we will never make any changes to your Datadog account),
so, following the principle of the least privilege, it's recommended to grant it read-only permissions.

## Configuration

To configure CloudQuery to extract from Datadog, create a `.yml` file in your CloudQuery configuration directory.
For example, the following configuration will extract information from the `cloudquery` organization, and connect it to a `postgresql` destination plugin

```yml
kind: source
spec:
  # Source spec section
  name: datadog
  path: cloudquery/datadog
  version: "0.0.1" # latest version of datadog plugin
  tables: ["*"]
  destinations: ["postgresql"]
  spec:
    accounts:
      - name:
        api_key: <DD_CLIENT_API_KEY> # Required. API key
        app_key: <DD_CLIENT_APP_KEY> # Required. app key
```
