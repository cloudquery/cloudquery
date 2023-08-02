```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: vercel
  path: /path/to/downloaded/plugin # Buy from here: https://cloudquery.io/integrations/vercle
  registry: local
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  skip_tables:
    - vercel_deployment_checks

  # Vercel specific configuration
  spec:
    backend_options:
      table_name: "test_state_table"
      connection: "@@plugins.DESTINATION_NAME.connection"
    access_token: "<YOUR_ACCESS_TOKEN_HERE>"
```

The Vercel plugin supports incremental syncing. This means that only new data will be fetched from Vercel and loaded into your destination for supported tables (support depending on API endpoint). This is done by keeping track of the last `paginator` for each table and only fetching data that has been updated since then.
To enable this, `backend` option must be set in the spec (as shown in the [example](/docs/plugins/sources/vercel/configuration#example)). This is documented in the [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables) section.
