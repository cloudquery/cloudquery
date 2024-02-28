```yaml copy
kind: source
spec:
  # Source spec section
  name: "datadog"
  path: "cloudquery/datadog"
  registry: "cloudquery"
  version: "VERSION_SOURCE_DATADOG"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  spec:
    # required
    accounts:
      - # required
        name: example_account
        # required
        api_key: ${DATADOG_EXAMPLE_ACCOUNT_API_KEY}
        # required
        app_key: ${DATADOG_EXAMPLE_ACCOUNT_APP_KEY}
    # Optional parameters
    # site: datadoghq.eu
```
