# Vercel Source Plugin Configuration Reference

## Example

This example syncs from Vercel to a Postgres destination, using API Key authentication. The (top level) source spec section is described in the [Source Spec Reference](https://www.cloudquery.io/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: vercel
  path: cloudquery/vercel
  version: "VERSION_SOURCE_VERCEL"
  tables: ["*"]
  destinations: ["postgresql"]
  skip_tables:
    - vercel_deployment_checks

  # Vercel specific configuration
  spec:
    access_token: "<YOUR_ACCESS_TOKEN_HERE>"
    team_ids:
      - "<YOUR_OPTIONAL_TEAM_ID_HERE>"
      - "<ANOTHER_OPTIONAL_TEAM_ID_HERE>"

```

## Vercel Spec

This is the (nested) spec used by the Vercel source plugin:

- `access_token` (string, required):
  An access token for your Vercel account. Get yours from [Vercel's Account Tokens Page](https://vercel.com/account/tokens).

- `team_ids` (list of string, optional):
  A list of team IDs to sync. If not specified, all teams will be synced. To find the ID of a specific team, refer to [Vercel Docs](https://vercel.com/docs/teams-and-accounts/create-or-join-a-team#find-your-team-id).

- `endpoint_url` (string, optional. Default: `https://api.vercel.com`):
  API endpoint URL to use. For Vercel's production API, this should be `https://api.vercel.com`.

- `timeout_secs` (integer in seconds, optional. Default: 5):
  Timeout for requests against the Vercel REST API.
