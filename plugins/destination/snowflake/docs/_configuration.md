This example sets the connection string to a value read from the `SNOWFLAKE_CONNECTION_STRING` environment variable:

```yaml copy
kind: destination
spec:
  name: snowflake
  path: cloudquery/snowflake
  registry: cloudquery
  version: "VERSION_DESTINATION_SNOWFLAKE"
  write_mode: "append"
  # Learn more about the configuration options at https://cql.ink/snowflake_destination
  spec:
    connection_string: "${SNOWFLAKE_CONNECTION_STRING}"
    # Optional parameters
    # migrate_concurrency: 1
    # batch_size: 1000 # 1K entries
    # batch_size_bytes: 4194304 # 4 MiB
```
