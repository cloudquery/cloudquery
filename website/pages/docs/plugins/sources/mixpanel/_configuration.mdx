```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: mixpanel
  path: /path/to/downloaded/plugin # Buy from here: https://cloudquery.io/integrations/mixpanel
  registry: local
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  # Mixpanel specific configuration
  spec:
    backend_options:
      table_name: "test_state_table"
      connection: "@@plugins.DESTINATION_NAME.connection"
    username: "<YOUR_SERVICE_ACCOUNT_USER_HERE>"
    secret: "<YOUR_SERVICE_ACCOUNT_SECRET_HERE>"
    project_id: 12345 # Your project ID
```

The Mixpanel plugin supports incremental syncing for event data. This means that only new events will be fetched from Mixpanel and loaded into your destination. This is done by keeping track of the last event fetched and only fetching events that has been created since then.
To enable this, `backend` option must be set in the spec (as shown in the example). This is documented in the [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables) section.
