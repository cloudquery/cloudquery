```yaml copy
kind: source
spec:
  # Source spec section
  name: digitalocean
  path: cloudquery/digitalocean
  registry: cloudquery
  version: "VERSION_SOURCE_DIGITALOCEAN"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  spec:
    # required, unless env variable DIGITALOCEAN_TOKEN or DIGITALOCEAN_ACCESS_TOKEN is set
    token: "${DIGITALOCEAN_ACCESS_TOKEN}"
    # Optional parameters
    # spaces_regions: ["nyc3", "sfo3", "ams3", "sgp1", "fra1", "syd1"]
    # spaces_access_key: ""
    # spaces_access_key_id: ""
    # spaces_debug_logging: false
    # concurrency: 10000
```
