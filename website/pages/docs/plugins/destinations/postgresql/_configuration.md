This example configures a Postgresql destination, located at `localhost:5432`:

```yaml copy
kind: destination
spec:
  name: "postgresql"
  registry: "github"
  path: "cloudquery/postgresql"
  version: "VERSION_DESTINATION_POSTGRESQL"

  spec:
    connection_string: "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
    # can be also specified in DSN format which can hold special characters in password
    # connection_string: "user=postgres password=pass+0-[word host=localhost port=5432 dbname=postgres"
    # Optional parameters:
    # pgx_log_level: error
    # batch_size: 10000 # 10K entries
    # batch_size_bytes: 100000000 # 100 MB
    # batch_timeout: 60s
```
