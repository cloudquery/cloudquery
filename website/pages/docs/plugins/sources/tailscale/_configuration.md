```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: tailscale
  path: /path/to/downloaded/plugin # Buy from here: https://cloudquery.io/integrations/tailscale
  registry: local
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  # plugin specific configuration
  spec:
    client_id: "<YOUR_CLIENT_ID_HERE>"
    client_secret: ${CLIENT_SECRET_ENV_VARIABLE}
    tailnet: "<YOUR_TAILNET>"
    endpoint_url: "<YOUR_BASE_URL>"
```
