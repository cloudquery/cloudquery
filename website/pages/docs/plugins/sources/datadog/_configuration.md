```yaml copy
kind: source
spec:
  # Source spec section
  name: "datadog"
  path: "cloudquery/datadog"
  version: "VERSION_SOURCE_DATADOG"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  spec:
    accounts:
      - name: example_account # Required. Name to distinguish accounts
        api_key: <DD_CLIENT_API_KEY> # Required. API key
        app_key: <DD_CLIENT_APP_KEY> # Required. App key
    # Optional parameters
    # site: datadoghq.eu
    # concurrency: 10000
```
