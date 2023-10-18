```yaml copy
kind: source
spec:
  name: "oracledb"
  path: "cloudquery/oracledb"
  version: "VERSION_SOURCE_ORACLEDB"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  spec:
    # To use the default port (1521) use `server:` instead of `server:port`
    connection_string: "oracle://user:password@localhost:/cloudquery"
```
