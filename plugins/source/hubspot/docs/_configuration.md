```yaml copy
kind: source
spec:
  name: "hubspot"
  path: cloudquery/hubspot
  registry: cloudquery
  version: "VERSION_SOURCE_HUBSPOT"
  destinations: ["DESTINATION_NAME"]
  tables: ["*"]
  spec:
    # required, unless the HUBSPOT_APP_TOKEN environment variable is set
    app_token: "${HUBSPOT_APP_TOKEN}"
    # optional, default is 5.
    # See https://developers.hubspot.com/docs/api/usage-details#rate-limits
    # max_requests_per_second: 5
```
