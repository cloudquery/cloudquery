---
name: DuckDB
stage: GA
title: DuckDB Destination Plugin
description: CloudQuery DuckDB destination plugin documentation
---

# DuckDB Destination Plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to a [DuckDB](https://duckdb.org/) database, either self-hosted, or in a compatible cloud data warehouse such as [MotherDuck](https://motherduck.com/).

## Example Config

:configuration

## DuckDB Spec

This is the top level spec used by the DuckDB destination Plugin.

- `connection_string` (`string`) (required)

  Absolute or relative path to a file, such as `./example.duckdb`.

- `batch_size` (`integer`) (optional) (default: `1000`)

  Maximum number of items that may be grouped together to be written in a single write.

- `batch_size_bytes` (`integer`) (optional) (default: `4194304` (4 MiB))

  Maximum size of items that may be grouped together to be written in a single write.

- `debug` (`boolean`) (optional) (default: `false`)

  Enables debug logging.

## Connecting to MotherDuck

To authenticate with MotherDuck, generate a service token and export it as an environment variable for CloudQuery to use. See [MotherDuck Documentation](https://motherduck.com/docs/key-tasks/service-accounts-guide/#managing-service-accounts-and-tokens) to learn how to generate a service token.

When setting the connection string, use the `md:` prefix to indicate syncing to the MotherDuck data warehouse and add the service token using the `motherduck_token` query string. We recommend passing the service token using variable expansion like this:

```yaml
spec:
  connection_string: "md:my_db?motherduck_token=${MOTHERDUCK_TOKEN}"
```

For a full walk-through on setting up a sync to a MotherDuck database, see [Moving Data from Postgres to MotherDuck](https://www.cloudquery.io/blog/moving-data-from-postgres-to-motherduck).
