```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: mixpanel
  path: cloudquery/mixpanel # Buy from here: https://cloudquery.io/integrations/mixpanel
  registry: cloudquery
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  backend_options:
    table_name: "cq_state_mixpanel"
    connection: "@@plugins.DESTINATION_NAME.connection"
  # Mixpanel specific configuration
  spec:
    username: "<YOUR_SERVICE_ACCOUNT_USER_HERE>"
    secret: "<YOUR_SERVICE_ACCOUNT_SECRET_HERE>"
    project_id: 12345 # Your project ID
```
