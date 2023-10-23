```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: slack
  path: cloudquery/slack # Buy from here: https://cloudquery.io/integrations/slack
  registry: cloudquery
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  # Slack specific configuration
  spec:
    token: "<YOUR_BOT_TOKEN_HERE>"
```
