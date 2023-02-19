# Okta Plugin

The CloudQuery Okta plugin extracts Okta resources configurations and loads them into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](https://www.cloudquery.io/docs/plugins/destinations/overview)).

## Links

- [Tables](./docs/tables/README.md)

## Authentication

To [authenticate](https://developer.okta.com/docs/guides/create-an-api-token/overview/) CloudQuery with your Okta account you need to set an `OKTA_API_TOKEN` environment variable or add it the configuration.

## Configuration

In order to get started with the Okta plugin, you need to create a YAML file in your CloudQuery configuration directory (e.g. named `okta.yml`).

The following example sets up the Okta plugin, and connects it to a postgresql destination:

```yaml
kind: source
spec:
  # Source spec section
  name: okta
  path: cloudquery/okta
  version: "v1.2.0" # latest version of okta plugin
  tables: ["*"]
  destinations: ["postgresql"]
  spec:
    # Required. Your Okta domain name
    domain: "https://<YOUR_OKTA_DOMAIN>.okta.com/"

    # Optional. Okta Token to access API, you can set this with OKTA_API_TOKEN environment variable
    # ⚠️ Warning - Your token should be kept secret and not committed to source control
    # token: "<YOUR_OKTA_TOKEN>"
```

- `domain` (Required) - Specify the Okta domain you are fetching from. [Visit this link](https://developer.okta.com/docs/guides/find-your-domain/findorg/) to find your Okta domain
- `token` (Optional) - Okta Token to access the API. You can set this with an `OKTA_API_TOKEN` environment variable
