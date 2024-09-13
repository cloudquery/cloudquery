This example configures a Postgresql destination using an environment variable called `POSTGRESQL_CONNECTION_STRING`:

```yaml copy
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  registry: "cloudquery"
  version: "v8.0.0"
  write_mode: "overwrite-delete-stale"
  # Learn more about the configuration options at https://cql.ink/postgresql_destination
  spec:
    connection_string: "${POSTGRESQL_CONNECTION_STRING}" # set the environment variable in a format like postgres://postgres:pass@localhost:5432/postgres?sslmode=disable
    # you can also specify it in DSN format, which can hold special characters in the password field:
    # connection_string: "user=postgres password=pass+0-[word host=localhost port=5432 dbname=postgres"
    # Optional parameters:
    # pgx_log_level: error
    # batch_size: 10000 # 10K entries
    # batch_size_bytes: 100000000 # 100 MB
    # batch_timeout: 60s

    # create_performance_indexes: false #create indexes that help with performance when using `write_mode: overwrite-delete-stale`
```
