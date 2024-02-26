```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: typeform
  registry: docker
  path: docker.cloudquery.io/cloudquery/source-typeform:VERSION_SOURCE_TYPEFORM
  tables: ["typeform_forms"]
  destinations: ["DESTINATION_NAME"]
  # Typeform-specific configuration
  spec:
    # required
    access_token: "${TYPEFORM_ACCESS_TOKEN}"
    # optional, default: https://api.typeform.com, use https://api.eu.typeform.com for EU accounts
    # base_url: "https://api.typeform.com"
    # Optional, default: 100
    # concurrency: 100
    # Optional, default: 10000
    # queue_size: 10000
```

:::callout{type="info"}
The Typeform plugin is distributed as a Docker image. This requires a Docker runtime to be installed on the same machine as the CloudQuery CLI, and a CLI version that supports the `docker` registry type (`v3.12.0` and higher).
:::