# Cloudflare Source Plugin Configuration Reference

## Example

This example syncs from Cloudflare to a Postgres destination, using in-line `api_token` authentication (instead of the `CLOUDFLARE_API_TOKEN` environment variable). The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: cloudflare
  path: cloudquery/cloudflare
  version: "VERSION_SOURCE_CLOUDFLARE"
  tables: ["*"]
  destinations: ["postgresql"]

  # Cloudflare specific configuration
  spec:
    api_token: "<YOUR_TOKEN_HERE>"
```

## Cloudflare Spec

This is the (nested) spec used by the Cloudflare source plugin.

- `api_token` (string, optional. Default: empty):
  An API token to access Cloudflare resources. This can also be set with the `CLOUDFLARE_API_TOKEN` environment variable. An API token authentication is preferred over API email and key authentication. If `api_token` is specified, `api_email` and `api_key` shouldn't be specified.

- `api_email` (string, optional. Default: empty):
  API email to access Cloudflare resources. If `api_email` is specified, `api_key` should also be specified. This can also be set with the `CLOUDFLARE_EMAIL` environment variable. If `api_email` is used, `api_token` shouldn't be specified.

- `api_key` (string, optional. Default: empty):
  API key to access Cloudflare resources. If `api_key` is used, `api_email` should also be specified. This can also be set with the `CLOUDFLARE_API_KEY` environment variable. If `api_key` is used, `api_token` shouldn't be specified.

- `accounts` ([]string, optional. Default: empty):
  List of accounts to target. If empty, all available accounts will be targeted.

- `zones` ([]string, optional. Default: empty):
  List of zones to target. If empty, all available zones will be targeted.

