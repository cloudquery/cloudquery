```yaml copy
kind: source
spec:
  name: "mysql"
  path: "cloudquery/mysql"
  registry: "cloudquery"
  version: "VERSION_SOURCE_MYSQL"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  spec:
    connection_string: "user:password@/dbname"
```
