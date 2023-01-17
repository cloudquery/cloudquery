# Snyk Source Plugin Configuration Reference

## Example

This example syncs from Snyk to a Postgres destination, using `api_key` authentication.
The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name:    snyk
  path:    cloudquery/snyk
  version: "VERSION_SOURCE_SNYK"
  tables: [ "*" ]
  destinations: [ "postgresql" ]

  # Snyk specific configuration
  spec:
    api_key:      "<YOUR_API_KEY_HERE>"
    organizations:
    - "<YOUR_ORG_1>"
    - "<YOUR_ORG_2>"
    endpoint_url: "<YOUR_BASE_URL>"
```

## Snyk Spec

This is the (nested) spec used by the Snyk source plugin.

- `api_key` (string, required):
  An API key to access Snyk resources.
  See [Authentication for API](https://docs.snyk.io/snyk-api-info/authentication-for-api) for more information.

- `organizations`  ([]string, optional. Default: all organizations accessible via `api_key`):
  You can choose to limit what organizations to sync information from.

- `endpoint_url` (string, optional. Default: not used):
  Endpoint URL to make the API requests to.
