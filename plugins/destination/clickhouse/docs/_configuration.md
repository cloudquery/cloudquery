```yaml copy
kind: destination
spec:
  name: "clickhouse"
  path: "cloudquery/clickhouse"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_CLICKHOUSE"
  write_mode: "append"
  # Learn more about the configuration options at https://cql.ink/clickhouse_destination
  spec:
    connection_string: "clickhouse://${CH_USER}:${CH_PASSWORD}@localhost:9000/${CH_DATABASE}"
    # Optional parameters
    # cluster: ""
    # ca_cert: ""
    # engine:
    #   name: MergeTree
    #   parameters: []
    #
    # batch_size: 10000
    # batch_size_bytes: 5242880 # 5 MiB
    # batch_timeout: 20s
```

This example configures a ClickHouse instance, located at `localhost:9000`.
It expects `CH_USER`, `CH_PASSWORD` and `CH_DATABASE` environment variables to be set.
The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).
