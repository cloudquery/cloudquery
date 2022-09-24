# CloudQuery PostgreSQL Destination Plugin

This destination plugin let's you sync data from any CloudQuery source to any PostgreSQL compatible database.

Known supported databases:

- PostgreSQL > v10

## PostgreSQL Spec

This is the top level spec used by PostgreSQL Destination Plugin.

- `connection_string` (string) (required)

  Connection string to connect to the database

- `pgx_log_level` (string) (default: "error")

  Available: "error", "warn", "info", "debug", "trace"
  define if and in which level to log [pgx](https://github.com/jackc/pgx) call.
