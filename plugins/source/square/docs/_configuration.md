```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: square
  registry: docker
  path: docker.cloudquery.io/cloudquery/source-square:VERSION_SOURCE_SQUARE
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  # Square-specific configuration
  spec:
    environment: "sandbox" # sandbox or production
    access_token: "<YOUR_SECRET_ACCESS_TOKEN_HERE>"
    concurrency: 100
    queue_size: 10000
```

:::callout{type="info"}
The Square plugin is distributed as a Docker image. This requires a Docker runtime to be installed on the same machine as the CloudQuery CLI, and a CLI version that supports the `docker` registry type (`v3.12.0` and higher).
:::
