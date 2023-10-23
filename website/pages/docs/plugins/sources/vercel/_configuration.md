```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: vercel
  path: /path/to/downloaded/plugin # Buy from here: https://cloudquery.io/integrations/vercel
  registry: local
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  skip_tables:
    - vercel_deployment_checks
  backend_options:
    table_name: "cq_state_vercel"
    connection: "@@plugins.DESTINATION_NAME.connection"

  # Vercel specific configuration
  spec:
    access_token: "<YOUR_ACCESS_TOKEN_HERE>"
```
