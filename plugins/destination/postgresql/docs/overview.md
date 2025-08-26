---
name: PostgreSQL
stage: GA
title: PostgreSQL Destination Plugin
description: CloudQuery PostgreSQL destination plugin documentation
---

# PostgreSQL Destination Plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to a PostgreSQL compatible database.

Supported database versions:

- PostgreSQL >= v10
- CockroachDB >= v20.2

## Configuration

### Example

:configuration

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

:::callout{type="info"}
Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.
:::

The PostgreSQL destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes).

### PostgreSQL Spec

This is the (nested) spec used by the PostgreSQL destination Plugin.

- `connection_string` (`string`) (required)

  Connection string to connect to the database. This can be a DSN or a URI, as per [`pgxpool`](https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool#ParseConfig)

  - `"user=user password=pass+0-[word host=localhost port=5432 dbname=mydb sslmode=disable"` _DSN format_
  - `"postgres://user:pass@localhost:5432/mydb?sslmode=prefer"` _connect with tcp and prefer TLS_
  - `"postgres://user:pass@localhost:5432/mydb?sslmode=disable&application_name=pgxtest&search_path=myschema&connect_timeout=5"` _be explicit with all options_
  - `"postgres://localhost:5432/mydb?sslmode=disable"` _connect with os username cloudquery is being run as_
  - `"postgres:///mydb?host=/tmp"` _connect over unix socket_
  - `"dbname=mydb"` _unix domain socket, just specifying the db name - useful if you want to use peer authentication_

- `pgx_log_level` (`string`) (optional) (default: `error`)

  Available: `error`, `warn`, `info`, `debug`, `trace`.
  Defines what [`pgx`](https://github.com/jackc/pgx) call events should be logged.

- `batch_size` (`integer`) (optional) (default: `10000`)

  Maximum number of items that may be grouped together to be written in a single write.

- `batch_size_bytes` (`integer`) (optional) (default: `100000000` (= 100 MB))

  Maximum size of items that may be grouped together to be written in a single write.

- `batch_timeout` (`duration`) (optional) (default: `60s` (= 60 seconds))

  Maximum interval between batch writes.

- `create_performance_indexes` (`boolean`) (optional) (default: `false`)

  Creates indexes on tables that help with performance when using `write_mode: overwrite-delete-stale`.

- `pgvector_config` (`object`) (optional)

  Optional configuration to enable PgVector embedding support.

  Note: source plugin must sync the `_cq_id` column on target tables if this is enabled.

  - `tables` (`array`) (required)
    
    Tables to create embeddings for. A table with the `_embeddings` suffix will be created for each table in this list.

    - `table_name` (`string`) (required)

    - `embed_columns` (`array`) (required)

    Columns to create embeddings for.

    - `metadata_columns` (`array`) (required)

    These columns will be added as-is from the source table for context. The `_cq_id` column will be added automatically and an index will be created on it.

  - `text_splitter` (`object`) (optional)

    Optional text splitting configuration for the embeddings. If unset, defaults are used.

    - `recursive_text` (`object`) (required)

      - `chunk_size` (`integer`) (required)

      Chunk size for the text splitting.

      - `chunk_overlap` (`integer`) (required)

      Chunk overlap for the text splitting.

  - `embedding` (`object`) (required)

    Embedding API provider configuration. Currently only OpenAI is supported.

    - `dimensions` (`integer`) (required)
    
    The number of dimensions to use for the embeddings. For `text-embedding-3-small`, this is 1536. For `text-embedding-3-large`, this is 3072.

    - `api_key` (`string`) (required)

    The OpenAI API key to use for the embedding API.

    - `model_name` (`string`) (required)

    The model name to use for the embedding API. Currently, `text-embedding-3-small` and `text-embedding-3-large` are supported.

### Verbose logging for debug

The PostgreSQL destination can be run in debug mode.

Note: This will use [`pgx`](https://github.com/jackc/pgx) built-in logging and might output data/sensitive information to logs so make sure to not use it in production but only for debugging.

```yaml copy
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  registry: cloudquery
  version: "VERSION_DESTINATION_POSTGRESQL"
  spec:
    connection_string: ${PG_CONNECTION_STRING}
    pgx_log_level: debug # Available: error, warn, info, debug, trace. Default: "error"
```
