# Tailscale Source Plugin Configuration Reference

## Example

This example syncs from Tailscale to a Postgres destination, using `api_key` authentication.
The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name:    tailscale
  path:    cloudquery/tailscale
  version: "VERSION_SOURCE_TAILSCALE"
  tables: [ "*" ]
  destinations: [ "postgresql" ]

  # Tailscale specific configuration
  spec:
    api_key:      "<YOUR_API_KEY_HERE>"
    tailnet:      "<YOUR_TAILNET>"
    endpoint_url: "<YOUR_BASE_URL>"
```

## Tailscale Spec

This is the (nested) spec used by the Tailscale source plugin.

- `api_key` (string, required):
  An API key to access Tailscale resources.
  This can be obtained from [Tailscale Keys Settings Page](https://login.tailscale.com/admin/settings/keys).

- `tailnet`  (string, required):
  This is your Tailscale tailnet name (also known as organization name).

- `endpoint_url` (string, optional. Default: not used):
  Endpoint URL to make the API requests to.
