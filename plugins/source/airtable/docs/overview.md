The CloudQuery Airtable plugin pulls data from [Airtable](https://airtable.com/) and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](https://hub.cloudquery.io/plugins/destination)).

The plugin discover all bases and tables in your account and syncs them to the destination.

## Example Configuration

This example syncs from Airtable to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

:configuration

## Authentication

:authentication

## Tables naming convention

In Airtable base names are not unique, and table names are unique within a base.
The plugin uses the following naming convention for tables: `<lowercase_base_id>__<snake_case_base_name>__<snake_case_table_name>`.
This ensures table names are unique across all bases (note the separator between name parts is `__`).

To sync only specific base(s) or table(s) you can use wildcard matching for the `tables` option, for example:

```yaml
kind: source
spec:
  name: airtable
  registry: docker
  path: docker.cloudquery.io/cloudquery/source-airtable:VERSION_SOURCE_AIRTABLE
  tables:
    # Sync all tables under bases matching the name `base_name_to_sync`
    - "*__<base_name_to_sync>__*"
    # Sync all tables matching the name `table_name_to_sync`
    - "*__*__<table_name_to_sync>"
    # Sync all tables matching the name `table_name_to_sync` under bases matching the name `base_name_to_sync`
    - "*__<base_name_to_sync>__<table_name_to_sync>"
  ...
```

## Configuration Reference

This is the (nested) spec used by the Airtable source plugin:

- `access_token` (`string`) (required)

  Your Airtable API [personal access token](https://airtable.com/developers/web/guides/personal-access-tokens).

- `endpoint_url` (`string`) (optional) (default: `https://api.airtable.com`)

  The endpoint URL to fetch data from.

- `concurrency` (`integer`) (optional) (default: `10000`)

  Best effort maximum number of tables to sync concurrently.

## Running in a Containerized Environment

To run the Airtable plugin in a containerized environment, see the [Using CloudQuery Docker Registry Plugins Inside a Containerized Environment](https://docs.cloudquery.io/docs/advanced-topics/using-cloud-query-docker-registry-plugins-inside-a-containerized-environment) guide.
