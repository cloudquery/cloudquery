# Shopify Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `shopify`)}/>

The CloudQuery Shopify plugin pulls data from Shopify and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

In order to fetch information from Shopify, `cloudquery` needs to be authenticated. Either an API key and password (in the case of basic custom/private apps) or an access token (for OAuth apps) is required for authentication.

Refer to the Shopify Help Center article on [Custom apps](https://help.shopify.com/en/manual/apps/custom-apps) and create a custom app. Follow _Get the API credentials for a custom app_ section to get the credentials for _Admin API_ and put them in your plugin configuration as `api_key` and `api_secret`.

If you have a large or busy store, API key/secret type credentials might not be enough due to the heavy rate limiting. In this case, you can use OAuth in your custom app to get an access token which allow many more requests a second. To use that token in your plugin configuration instead, just set it in `access_token` and remove `api_key` and `api_secret` sections. For more information, refer to [Shopify.dev](https://shopify.dev/apps/distribution) on the subject.


# Configuration Reference

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

# Query Examples

## Get all your active products with a specific tag

```sql copy
SELECT * FROM shopify_products WHERE status='active' AND 'your-tag' = ANY(tags);
```

