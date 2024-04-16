In order to fetch information from Cloudflare, `cloudquery` needs to be authenticated. There are a few options for authentication:

- Via an [API token](https://developers.cloudflare.com/fundamentals/api/get-started/create-token/) (preferred)
- Via an API email and key

The plugin requires only _read_ permissions.  To start, Cloudflare has a `read all resources` [API token template](https://developers.cloudflare.com/fundamentals/api/reference/template/) that will grant CloudQuery the necessary permissions to fetch information from Cloudflare.  As necessary, those permissions can be refined and modified further to meet your needs.