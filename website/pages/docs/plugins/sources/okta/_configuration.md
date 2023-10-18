```yaml copy
kind: source
spec:
  # Source spec section
  name: okta
  path: cloudquery/okta
  version: "VERSION_SOURCE_OKTA"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  spec:
    # Required. Your Okta domain name
    domain: "https://<YOUR_OKTA_DOMAIN>.okta.com/"

    # Optional. Okta Token to access API
    # ⚠️ Warning - Your token should be kept secret and not committed to source control.
    # ⚠️ Warning - In the future versions token parameter will become required.
    # token: "<YOUR_OKTA_TOKEN>"

    # Optional. Rate limiter settings
    # rate_limit:
    #   max_backoff: 5s
    #   max_retries: 3

```

