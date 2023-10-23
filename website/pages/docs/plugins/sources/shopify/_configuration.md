```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: shopify
  path: cloudquery/shopify
  registry: cloudquery
  version: "VERSION_SOURCE_SHOPIFY"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  backend_options:
    table_name: "cq_state_shopify"
    connection: "@@plugins.DESTINATION_NAME.connection"
  # Shopify specific configuration
  spec:
    api_key: "<YOUR_API_KEY_HERE>"
    api_secret: "<YOUR_API_SECRET_HERE>"
    shop_url: "https://<YOUR_SHOP>.myshopify.com"
#    concurrency: 1000
```
