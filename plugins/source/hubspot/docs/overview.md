---
name: HubSpot
stage: GA
title: HubSpot Source Plugin
description: CloudQuery HubSpot source plugin documentation
---

# HubSpot Source Plugin

:badge

The CloudQuery HubSpot plugin extracts HubSpot information and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)). It is based on the [HubSpot API](https://developers.hubspot.com/docs/api/overview) and the [github.com/clarkmcc/go-hubspot](https://github.com/clarkmcc/go-hubspot) library.

:::callout{type="warning"}
The HubSpot REST API is [rate limited](https://developers.hubspot.com/docs/api/usage-details#rate-limits).
By default, CloudQuery makes up to 5 API requests per second - but you may need to increase/decrease this value depending on your HubSpot subscription.
You can use the `max_requests_per_second` configuration option to change this value (see below).
:::

## Authentication

:authentication

## Configuration

The following example sets up the HubSpot plugin, and connects it to a postgresql destination:

:configuration

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
