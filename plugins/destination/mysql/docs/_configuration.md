```yaml copy
kind: destination
spec:
  name: "mysql"
  path: "cloudquery/mysql"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_MYSQL"
  spec:
    connection_string: "user:password@/dbname"
    # Optional parameters:
    # batch_size: 1000 # 1K entries
    # batch_size_bytes: 4194304 # 4 MiB
```