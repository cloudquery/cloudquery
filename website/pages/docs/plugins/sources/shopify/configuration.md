# Shopify Source Plugin Configuration Reference

## Example

This example syncs from Shopify to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](https://www.cloudquery.io/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: shopify
  path: cloudquery/shopify
  version: "VERSION_SOURCE_SHOPIFY"
  tables: ["*"]
  destinations: ["postgresql"]

  # Shopify specific configuration
  spec:
    api_key: "<YOUR_API_KEY_HERE>"
    api_secret: "<YOUR_API_SECRET_HERE>"
    shop_url: "https://<YOUR_SHOP>.myshopify.com"
```

## Shopify Spec

This is the (nested) spec used by the Shopify source plugin:

- `api_key` (string, required*):
  The API Key for your custom app in your store.

- `api_secret` (string, required*):
  The API Secret for your custom app in your store.

- `access_token` (string, required if api_key/secret is not used):
  An access token for your Shopify custom app. This is an alternative way of authenticating, use either this or the ones above.

- `shop_url` (string, required): The URL of your Shopify store. Must start with `https://` and end with `.myshopify.com`.

- `timeout_secs` (integer in seconds, optional. Default: 10):
  Timeout for requests against the Shopify Admin API.

- `max_retries` (integer, optional. Default: 30):
  Number of retries if a request was rate limited.
