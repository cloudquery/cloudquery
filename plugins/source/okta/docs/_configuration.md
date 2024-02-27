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
    # Okta domain name, for example: https://example.okta.com, https://example.okta-emea.com,  https://example.oktapreview.com
    domain: "${OKTA_DOMAIN}"
    # Okta Token to access API
    token: "${OKTA_ACCESS_TOKEN}"

    # Optional. Rate limiter settings
    # rate_limit:
    #   max_backoff: 5s
    #   max_retries: 3
```
