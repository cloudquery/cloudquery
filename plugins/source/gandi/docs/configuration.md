# Gandi Source Plugin Configuration Reference

## Example

This example syncs from Gandi to a Postgres destination, using API Key authentication. The (top level) source spec section is described in the [Source Spec Reference](https://www.cloudquery.io/docs/reference/source-spec).

```yml
kind: source
# Common source-plugin configuration
spec:
  name: gandi
  path: cloudquery/gandi
  version: "v1.0.0" # latest version of gandi plugin
  tables: ["*"]
  destinations: ["postgresql"]

  # Gandi specific configuration
  spec:
    api_key: "<YOUR_API_KEY_HERE>"
```

## Gandi Spec

This is the (nested) spec used by the Gandi source plugin.

- `api_key` (string, required):
  An API key to access your Gandi resources. Get your API key from [Gandi's Account Settings Page](https://account.gandi.net/en/).

- `sharing_id` (string, optional. Default: empty):
  This is your Organization ID. If specified, synced resources are filtered by this id.

- `gandi_debug` (boolean, optional. Default: false):
  Enables verbose debugging of HTTP calls made by the Gandi API client.

- `endpoint_url` (string, optional. Default: `https://api.gandi.net`):
  API endpoint URL to use. For Gandi's production API, this should be `https://api.gandi.net`. For the sandbox API, this should be `https://api.ote.gandi.net`.

- `timeout_secs` (integer in seconds, optional. Default: 5):
  Timeout for requests against the Gandi API.

