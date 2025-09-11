This example configures a Gremlin destination, located at `ws://localhost:8182`. The username and password are stored in environment variables.

```yaml
kind: destination
spec:
  name: "gremlin"
  path: "cloudquery/gremlin"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_GREMLIN"
  send_sync_summary: true
  spec:
    endpoint: "ws://localhost:8182"
    # Optional parameters
    # auth_mode: none
    # username: ""
    # password: ""
    # aws_region: ""
    # aws_neptune_host: ""
    # max_retries: 5
    # max_concurrent_connections: 5 # default: number of CPUs
    # batch_size: 200
    # batch_size_bytes: 4194304 # 4 MiB
```
