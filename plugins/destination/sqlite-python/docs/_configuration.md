This is a basic configuration that will save all your sync resources to `db.sql`.

```yaml copy
kind: destination
spec:
  name: sqlite-python
  path: cloudquery/sqlite-python
  registry: cloudquery
  version: "VERSION_DESTINATION_SQLITE_PYTHON"
  # Learn more about the configuration options at https://cql.ink/sqlite-python_destination
  spec:
    connection_string: ./db.sql
```

After running `cloudquery sync`, you can explore the data locally with the SQLite CLI: `sqlite ./db.sql`.
