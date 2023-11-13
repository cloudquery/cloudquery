---
name: Mixpanel
stage: GA (Premium)
title: Mixpanel Source Plugin
description: CloudQuery Mixpanel source plugin documentation
---
# Mixpanel Source Plugin

:badge{text="Premium"}

This is a premium plugin that you can buy [here](/integrations/mixpanel).

The CloudQuery Mixpanel plugin pulls data from Mixpanel and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](https://hub.cloudquery.io/plugins/destination)).

## Authentication

:authentication

### Example

This example syncs from Mixpanel to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

:configuration

:::callout{type="info"}
Note that if `backend_options` is omitted, by default no backend will be used.
This will result in all items being fetched on every sync.

For more information about managing state for incremental tables, see [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables).
:::

## Configuration Reference

This is the (nested) spec used by the Mixpanel source plugin:

- `username` (string, required):
  Your Service Account username from Mixpanel.

- `secret` (string, required):
  Service Account secret.

- `project_id` (integer, required):
  ID of the Mixpanel project to sync data from.

- `workspace_id` (integer, optional):
  ID of the Mixpanel workspace to sync data from.

- `start_date` (date in YYYY-MM-DD format, optional):
  Start date to sync data from. Defaults to 30 days ago.

- `end_date` (date in YYYY-MM-DD format, optional):
  End date to sync data until. Defaults to today.

- `region` (string, optional):
  Your Mixpanel region. Defaults to "US". Possible values are "US" and "EU".

- `timeout_secs` (integer in seconds, optional. Default: 30):
  Timeout for requests against the Mixpanel API.

- `max_retries` (integer, optional. Default: 5):
  Number of retries if a request was rate limited at the API endpoint.
