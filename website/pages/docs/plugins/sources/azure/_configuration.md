```yaml copy
kind: source
spec:
  # Source spec section
  name: "azure"
  path: "cloudquery/azure"
  version: "VERSION_SOURCE_AZURE"
  destinations: ["DESTINATION_NAME"]
  tables: ["azure_compute_virtual_machines"]
  spec:
    # Optional parameters
    # subscriptions: []
    # cloud_name: ""
    # concurrency: 50000
    # discovery_concurrency: 400
    # skip_subscriptions: []
    # normalize_ids: false
    # oidc_token: ""
    # retry_options:
    #   max_retries: 3
    #   try_timeout: 0s
    #   retry_delay: 4s
    #   max_retry_delay: 60s
```
