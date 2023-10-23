```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: render
  path: cloudquery/render # Buy from here: https://cloudquery.io/integrations/render
  registry: cloudquery
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  # Plugin specific configuration
  spec:
    token: ${API_KEY} # requried
```
