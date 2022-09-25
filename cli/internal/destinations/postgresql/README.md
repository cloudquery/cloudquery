# CloudQuery PostgreSQL Destination Plugin

This destination plugin let's you sync data from a CloudQuery source to a PostgreSQL compatible database.

Known supported databases:

- PostgreSQL > v10

## PostgreSQL Spec

This is the top level spec used by the PostgreSQL destination Plugin.

- `connection_string` (string) (required)

  Connection string to connect to the database

- `pgx_log_level` (string) (optional, defaults to "error")

  Available: "error", "warn", "info", "debug", "trace"
  define if and in which level to log [pgx](https://github.com/jackc/pgx) call.
