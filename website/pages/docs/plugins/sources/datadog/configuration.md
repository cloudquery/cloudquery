# CloudQuery Datadog Source Plugin Configuration Reference

## Example

This example connects a single Datadog account to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml copy
kind: source
spec:
  # Source spec section
  name: "datadog"
  path: "cloudquery/datadog"
  version: "VERSION_SOURCE_DATADOG"
  destinations: ["postgresql"]

  spec:
    accounts:
      - name: example_account # Required. Name to distinct accounts
        api_key: <DD_CLIENT_API_KEY> # Required. API key
        app_key: <DD_CLIENT_APP_KEY> # Required. app key
```

## Datadog Spec

This is the (nested) spec used by the Datadog source plugin.

- `accounts` ([]struct) 

  Specify which accounts to sync data from.
