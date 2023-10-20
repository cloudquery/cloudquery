```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: cloudflare
  path: cloudquery/cloudflare
  version: "VERSION_SOURCE_CLOUDFLARE"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  # Cloudflare specific configuration
  spec:
    # Optional parameters
    # api_token: ""
    # api_email: ""
    # api_key: ""
    # accounts: []
    # zones: []
    # concurrency: 10000
```
