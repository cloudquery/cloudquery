---
name: Stripe
stage: GA
title: Stripe Source Plugin
description: CloudQuery Stripe source plugin documentation
---

# Stripe Source Plugin

:badge

The CloudQuery Stripe plugin pulls data from Stripe and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Authentication

:authentication

### Example

This example syncs from Stripe to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec). Incremental syncing is enabled and will be saved to a `.cq/state/` directory by default.

:configuration

:::callout{type="info"}
Note that if `backend_options` is omitted, by default no backend will be used.
This will result in all items being fetched on every sync.

For more information about managing state for incremental tables, see [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables).
:::

## Configuration Reference

This is the (nested) spec used by the Stripe source plugin:

- `api_key` (`string`) (required):

  Your secret API key from the Stripe Dashboard.

- `rate_limit` (`integer`) (optional)  (default: varies):

  Used to override number of requests allowed per second. Defaults to 90 req/sec for production environment keys, otherwise 20 req/sec.

- `max_retries` (`integer`) (optional)  (default: `2`):

  Number of retries if a request was rate limited at the API endpoint.

- `concurrency` (`integer`) (optional)  (default: `10000`):

  Number of concurrent requests to Stripe API.

- `stripe_debug` (`boolean`) (optional)  (default: `false`):

  Enables verbose logging on the Stripe client.
