```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: plausible
  path: /path/to/downloaded/plugin # Buy from here: https://cloudquery.io/integrations/plausible
  registry: local
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  # Plugin specific configuration
  spec:
    site_id: "YOUR_SITE_ID" # required
    api_key: ${API_KEY} # requried
    period: "30d" # optional
```
