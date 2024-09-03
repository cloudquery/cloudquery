---
name: DuckDB
stage: GA
title: DuckDB Destination Plugin
description: CloudQuery DuckDB destination plugin documentation
---
# DuckDB Destination Plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to a [DuckDB](https://duckdb.org/) database.

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

- `appender` (`boolean`) (optional) (default: `false`)

  Enables `Appender API` (Preview)
