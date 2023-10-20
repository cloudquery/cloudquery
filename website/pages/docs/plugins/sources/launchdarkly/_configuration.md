```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: launchdarkly
  path: /path/to/downloaded/plugin # Buy from here: https://cloudquery.io/integrations/launchdakrly
  registry: local
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  backend_options:
    table_name: "cq_state_launchdarkly"
    connection: "@@plugins.DESTINATION_NAME.connection"
  # LaunchDarkly specific configuration
  spec:
    access_token: "<YOUR_ACCESS_TOKEN_HERE>"
```
