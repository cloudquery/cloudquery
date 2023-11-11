```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: googleads
  path: /path/to/downloaded/plugin # Buy from here: https://cloudquery.io/integrations/googleads
  registry: local
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  # Google Ads specific configuration
  spec:
    developer_token: "<GOOGLEADS_DEVELOPER_TOKEN>"
    login_customer_id: "<GOOGLEADS_MANAGEMENT_ACCOUNT_ID>"
    oauth:
      access_token: "<YOUR_OAUTH_ACCESS_TOKEN>"
```
