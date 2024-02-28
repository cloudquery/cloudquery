```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: cloudflare
  path: cloudquery/cloudflare
  registry: cloudquery
  version: "VERSION_SOURCE_CLOUDFLARE"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  # Cloudflare specific configuration
  spec:
    # required, if api_email and api_key are not set
    api_token: "${CLOUDFLARE_API_TOKEN}"
    # required, if api_token is not set
    # api_email: ""
    # required, if api_token is not set
    # api_key: ""
    # Optional parameters
    # accounts: []
    # zones: []
```
