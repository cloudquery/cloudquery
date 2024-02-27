```yaml copy
kind: source
spec:
  # Source spec section
  name: okta
  path: cloudquery/okta
  registry: cloudquery
  version: "VERSION_SOURCE_OKTA"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  spec:
    # Okta domain name
    domain: "https://<YOUR_OKTA_DOMAIN>.okta.com/"
    # Okta Token to access API
    token: "${OKTA_ACCESS_TOKEN}"

    # Optional. Rate limiter settings
    # rate_limit:
    #   max_backoff: 5s
    #   max_retries: 3

```
