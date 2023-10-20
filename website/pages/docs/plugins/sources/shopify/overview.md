---
name: Shopify
stage: GA
title: Shopify Source Plugin
description: CloudQuery Shopify source plugin documentation
---

# Shopify Source Plugin

:badge

The CloudQuery Shopify plugin pulls data from Shopify and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Authentication

:authentication

## Incremental Syncing

The Shopify plugin supports incremental syncing. This means that only new data will be fetched from Shopify and loaded into your destination for supported tables (support depending on API endpoint). This is done by keeping track of the last item fetched and only fetching data that has been created since then.
To enable this, `backend` option must be set in the spec (as shown below). This is documented in the [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables) section.

## Example

This example syncs from Shopify to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec). Incremental syncing is enabled and will be saved to a `.cq/state/` directory by default.

:configuration

:::callout{type="info"}
Note that if `backend_options` is omitted, by default no backend will be used.
This will result in all items being fetched on every sync.

For more information about managing state for incremental tables, see [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables).
:::

# Configuration Reference

This is the (nested) spec used by the Shopify source plugin:

- `api_key` (`string`) (required if `access_token` isn't used):

  The API Key for your custom app in your store.

- `api_secret` (`string`) (required if `access_token` isn't used):

  The API Secret for your custom app in your store.

- `access_token` (`string`) (required if `api_key` & `api_secret` aren't used):

  An access token for your Shopify custom app. This is an alternative way of authenticating, use either this or the ones above.

- `shop_url` (`string`) (required):

  The URL of your Shopify store. Must start with `https://` and end with `.myshopify.com`.

- `api_version` (string) (optional) (default: `2023-01`):

  The Shopify Admin API version to use. See [here](https://shopify.dev/docs/api/usage/versioning) for more information.

- `timeout_secs` (`integer`) (optional) (default: `10`):

  Timeout (in seconds) for requests against the Shopify Admin API.

- `max_retries` (`integer`) (optional) (default: `30`):

  Number of retries if a request was rate limited.

- `page_size` (`integer`) (optional) (default: `50`):

  Maximum number of items queried each request. Find an optimum value to balance amount of data fetched and requests timing out. Maximum value 250.

- `concurrency` (`integer`) (optional) (default: `1000`):

  Maximum number of concurrent requests to the Shopify Admin API.

# Query Examples

## Get all your active products with a specific tag

```sql copy
SELECT * FROM shopify_products WHERE status='active' AND 'your-tag' = ANY(tags);
```
