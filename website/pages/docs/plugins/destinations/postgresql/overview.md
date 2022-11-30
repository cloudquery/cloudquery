# CloudQuery PostgreSQL Destination Plugin

This destination plugin lets you sync data from a CloudQuery source to a PostgreSQL compatible database.

Supported database versions:

- PostgreSQL > v10
- CockroachDB > v20.2

## Configuration

### Example

This example configures a Postgresql destination, located at `localhost:5432`. The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "postgresql"
  registry: github
  path: "postgresql"
  version: "VERSION_DESTINATION_POSTGRESQL"

  spec:
    connection_string: "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
```

### PostgreSQL Spec

This is the (nested) spec used by the PostgreSQL destination Plugin.

- `connection_string` (string, required)

  Connection string to connect to the database. This can be a URL or a DSN, as per [`pgxpool`](https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool#ParseConfig)

  - `"postgres://jack:secret@localhost:5432/mydb?sslmode=prefer"` _connect with tcp and prefer TLS_
  - `"postgres://jack:secret@localhost:5432/mydb?sslmode=disable&application_name=pgxtest&search_path=myschema&connect_timeout=5"` _be explicit with all options_
  - `"postgres://localhost:5432/mydb?sslmode=disable"` _connect with os username cloudquery is being run as_
  - `"postgres:///mydb?host=/tmp"` _connect over unix socket_
  - `"dbname=mydb"` _unix domain socket, just specifying the db name - useful if you want to use peer authentication_
  - `"user=jack password=jack\\'ssooper\\\\secret host=localhost port=5432 dbname=mydb sslmode=disable"` _DSN with escaped backslash and single quote_

- `pgx_log_level` (string, optional. Default: "error")

  Available: "error", "warn", "info", "debug", "trace"
  define if and in which level to log [`pgx`](https://github.com/jackc/pgx) call.

- `batch_size` (int, optional. Default: 1000)

  Number of rows to insert in a single batch.
