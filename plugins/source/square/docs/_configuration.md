```yaml copy
kind: source
# Common source-integration configuration
spec:
  name: square
  registry: docker
  path: docker.cloudquery.io/cloudquery/source-square:VERSION_SOURCE_SQUARE
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  # Square-specific configuration
  spec:
    # required
    access_token: "${SQUARE_ACCESS_TOKEN}"
    # required
    environment: "sandbox" # sandbox or production
    # optional, default: 100
    # concurrency: 100
    # optional, default: 10000
    # queue_size: 10000
```

:::callout{type="info"}
The Square plugin is distributed as a Docker image. This requires a Docker runtime to be installed on the same machine as the CloudQuery CLI, and a CLI version that supports the `docker` registry type (`v3.12.0` and higher).
:::
