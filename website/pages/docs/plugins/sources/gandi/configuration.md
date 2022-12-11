# Gandi Source Plugin Configuration Reference

## Example

This example syncs from Gandi to a Postgres destination, using `api_key` authentication. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: gandi
  path: cloudquery/gandi
  version: "VERSION_SOURCE_GANDI"
  tables: ["*"]
  destinations: ["postgresql"]

  # Gandi specific configuration
  spec:
    api_key: "<YOUR_API_KEY_HERE>"
```

## Gandi Spec

This is the (nested) spec used by the Gandi source plugin.

- `api_key` (string, required):
  An API key to access Gandi resources. This can be generated from [Gandi's Account Settings Page](https://account.gandi.net/en/).

- `sharing_id` (string, optional. Default: not used):
  This is your Gandi Organization ID if you wish to limit the queries to a specific Gandi org.

- `gandi_debug` (bool, optional. Default: false):
  Enables verbose logging of HTTP calls made by the [underlying library](https://github.com/go-gandi/go-gandi).

- `endpoint_url` (string, optional. Default: `https://api.gandi.net`):
  Endpoint URL to make the API requests to. To use the [Gandi Sandbox](https://api.sandbox.gandi.net/docs/sandbox/) API endpoint, set this to `https://api.sandbox.gandi.net/`.

- `timeout_secs` (integer as seconds, optional. Default: `300`):
  Default timeout for each API request. Defaults to 5 minutes.

