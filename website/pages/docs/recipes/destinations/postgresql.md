# PostgreSQL Destination Plugin Recipes

Full spec options for PostgreSQL destination available [here](https://github.com/cloudquery/cloudquery/tree/main/plugins/destination/postgresql).

Note: Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.

## Basic

```yaml copy
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "v1.7.11" # latest version of postgresql plugin
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
  version: "v1.7.11" # latest version of postgresql plugin
  spec:
    connection_string: ${PG_CONNECTION_STRING}
    pgx_log_level: debug # Available: error, warn, info, debug, trace. Default: "error"
```
