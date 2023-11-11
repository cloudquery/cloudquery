---
name: LaunchDarkly
stage: GA (Premium)
title: LaunchDarkly Source Plugin
description: CloudQuery LaunchDarkly source plugin documentation
---
# LaunchDarkly Source Plugin

:badge{text="Premium"}

This is a premium plugin that you can buy [here](/integrations/launchdarkly).

The CloudQuery LaunchDarkly plugin pulls data from LaunchDarkly and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Authentication

:authentication

### Example

This example syncs from LaunchDarkly to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

:configuration

:::callout{type="info"}
Note that if `backend_options` is omitted, by default no backend will be used.
This will result in all items being fetched on every sync.

For more information about managing state for incremental tables, see [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables).
:::

## Configuration Reference

This is the (nested) spec used by the LaunchDarkly source plugin:

- `access_token` (`string`) (required):

  Your access token from LaunchDarkly.
- `timeout_secs` (`integer`) (optional) (default: `30`):

  Timeout (in seconds) for requests against the LaunchDarkly API.
