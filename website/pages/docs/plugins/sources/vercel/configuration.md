# Vercel Source Plugin Configuration Reference

## Example

This example syncs from Vercel to a Postgres destination, using API Key authentication. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec). Incremental syncing is enabled and will be saved to a `.cq/state/` directory by default.

:configuration

:::callout{type="info"}
Note that if `backend_options` is omitted, by default no backend will be used.
This will result in all items being fetched on every sync.

For more information about managing state for incremental tables, see [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables).
:::

## Vercel Spec

This is the (nested) spec used by the Vercel source plugin:

- `access_token` (`string`) (required):

  An access token for your Vercel account. Get yours from [Vercel's Account Tokens Page](https://vercel.com/account/tokens).

- `team_ids` (`[]string`) (optional):

  A list of team IDs to sync. If not specified, all teams will be synced. To find the ID of a specific team, refer to [Vercel Docs](https://vercel.com/docs/accounts/create-a-team#find-your-team-id).

- `endpoint_url` (`string`) (optional) (default: `https://api.vercel.com`):

  API endpoint URL to use. For Vercel's production API, this should be `https://api.vercel.com`.

- `timeout_secs` (`integer`) (optional) (default: `5`):

  Timeout (in seconds) for requests against the Vercel REST API.

- `max_retries` (`integer`) (optional) (default: `10`):

  Maximum number of retries for requests against the Vercel REST API.

- `max_wait_secs` (`integer`) (optional) (default: `300`):

  Maximum wait time (in seconds) between rate limited API requests. The plugin waits until the rate limit resets. If there's a need to wait longer than this time, the request will fail.

- `page_size` (`integer`) (optional) (default: `100`):

  Number of items to request in each API call. This is a tradeoff between the number of API calls and the number of items per API call, which might take too long. The maximum allowed value is 100.

- `concurrency` (`integer`) (optional) (default: `10000`):

  Number of concurrent requests to Vercel REST API.
