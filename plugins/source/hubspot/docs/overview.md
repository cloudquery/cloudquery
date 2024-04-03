# HubSpot Source Plugin

The CloudQuery HubSpot plugin extracts HubSpot information and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)). It is based on the [HubSpot API](https://developers.hubspot.com/docs/api/overview) and the [github.com/clarkmcc/go-hubspot](https://github.com/clarkmcc/go-hubspot) library.

:::callout{type="warning"}
The HubSpot REST API is [rate limited](https://developers.hubspot.com/docs/api/usage-details#rate-limits).
By default, CloudQuery makes up to 5 API requests per second - but you may need to increase/decrease this value depending on your HubSpot subscription.
You can use the `max_requests_per_second` configuration option to change this value (see below).
:::

## Authentication

In Order for CloudQuery to sync resources from your HubSpot setup, you will need to authenticate with your HubSpot account. You will need to create a [HubSpot Private App](https://developers.hubspot.com/docs/api/private-apps), and copy the App Token to the spec.
If not specified `HUBSPOT_APP_TOKEN` environment variable will be used instead.

```bash copy
export HUBSPOT_APP_TOKEN=<your_app_token> # optional, if not using spec configuration
```

## Configuration

The following example sets up the HubSpot plugin and connects it to a postgresql destination. Incremental syncing is enabled and will be saved to the `cq_state_hubspot` table.

```yaml copy
kind: source
spec:
  name: "hubspot"
  path: cloudquery/hubspot
  registry: cloudquery
  version: "VERSION_SOURCE_HUBSPOT"
  destinations: ["DESTINATION_NAME"]
  backend_options:
    table_name: "cq_state_hubspot"
    connection: "@@plugins.DESTINATION_NAME.connection"
  tables: ["*"]
  spec:
    # required, unless the HUBSPOT_APP_TOKEN environment variable is set
    app_token: "${HUBSPOT_APP_TOKEN}"
    # optional, default is 5.
    # See https://developers.hubspot.com/docs/api/usage-details#rate-limits
    # max_requests_per_second: 5
```

:::callout{type="info"}
Note that if `backend_options` is omitted, by default no backend will be used.
This will result in all items being fetched on every sync.

For more information about managing state for incremental tables, see [Managing Incremental Tables](https://cloudquery.io/docs/advanced-topics/managing-incremental-tables).
:::

### HubSpot Spec

This is the specs that can be used by the HubSpot source Plugin.

- `app_token` (`string`)  (optional) (default: `HUBSPOT_APP_TOKEN` environment variable value)
  The HubSpot App Token to use for authentication. This can also be set with the `HUBSPOT_APP_TOKEN` environment variable. 

- `concurrency` (`integer`) (optional) (default: `1000`)

  A best effort maximum number of Go routines to use. Lower this number to reduce memory usage.

- `max_requests_per_second` (`integer`) (optional) (default: `5`)

  Rate limit per second for requests done HubSpot API, this will depend on your HubSpot plan (https://developers.hubspot.com/docs/api/usage-details#rate-limits)

- `table_options` (map[string][TableOptions](#table-options) spec) (optional) (default: `empty`)

  Table Options for HubSpot entities that will be synced.

### Table Options

- `associations` (`[]string`) (optional) (default: empty)

  Additional associations to be retrieved from HubSpot when syncing the table entity

- `properties` (`[]string`) (optional) (default: empty)

  Additional properties to be retrieved from HubSpot when syncing the table entity
