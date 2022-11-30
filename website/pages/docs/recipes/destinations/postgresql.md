# PostgreSQL Destination Plugin Recipes

Full spec options for the PostgreSQL destination plugin are available [here](/docs/plugins/destinations/postgresql/overview#postgresql-spec).

Note: Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.

## Basic

```yaml copy
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "VERSION_DESTINATION_POSTGRESQL"
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```

## Verbose logging for debug

Run PostgreSQL destination in debug mode:

Note: This will use [`pgx`](https://github.com/jackc/pgx) built-in logging and might output data/sensitive information to logs so make sure to not use it in production but only for debugging.

```yaml copy
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "VERSION_DESTINATION_POSTGRESQL"
  spec:
    connection_string: ${PG_CONNECTION_STRING}
    pgx_log_level: debug # Available: error, warn, info, debug, trace. Default: "error"
```
