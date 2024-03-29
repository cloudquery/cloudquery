# CloudQuery Datadog Source Plugin Configuration Reference

## Example

This example connects a single Datadog account to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

:configuration

## Datadog Spec

This is the (nested) spec used by the Datadog source plugin.

- `accounts` ([][account](#account), required)

  Specify which accounts to sync data from.

- `site` (`string`, optional, default: "")

  The Datadog site to connect to. This is usually one of `datadoghq.com` or `datadoghq.eu` - see [site](https://docs.datadoghq.com/getting_started/site/) documentation for more information.

- `concurrency` (`int`, optional, default: `10000`)

  A best effort maximum number of Go routines to use. Lower this number to reduce memory usage.

### account

  This is used to specify one or more accounts to extract information from.
  Note that it should be an array of objects, each with the following fields:

- `name` (`string`, required)

  Account name.

- `api_key` (`string`, required)

  Datadog API key.

- `app_key` (`string`, required)

  Datadog App key.
