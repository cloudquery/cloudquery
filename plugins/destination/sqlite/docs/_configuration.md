This is a basic configuration that will save all your sync resources to `db.sql`.

```yaml copy
kind: destination
spec:
  name: sqlite
  path: cloudquery/sqlite
  registry: cloudquery
  version: "VERSION_DESTINATION_SQLITE"
  # Learn more about the configuration options at https://cql.ink/sqlite_destination
  spec:
    connection_string: ./db.sql
```

After running `cloudquery sync`, you can explore the data locally with the SQLite CLI: `sqlite ./db.sql`.
