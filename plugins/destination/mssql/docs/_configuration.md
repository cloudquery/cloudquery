```yaml copy
kind: destination
spec:
  name: "mssql"
  path: "cloudquery/mssql"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_MSSQL"
  spec:
    # Connection string in the format `server=localhost;user id=SA;password=yourStrongP@ssword;port=1433;database=cloudquery;`
    connection_string: "${MSSQL_CONNECTION_STRING}"
    # Optional parameters:
    # auth_mode: ms
    # schema: dbo
    # batch_size: 1000 # 1K entries
    # batch_size_bytes: 5242880 # 5 MiB
    # batch_timeout: 20s
```
