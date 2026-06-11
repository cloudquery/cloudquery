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

The (top level) spec section is described in the [Destination Spec Reference](/docs/cli/integrations/destinations#complete-destination-spec-reference).

:::callout{type="info"}
Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.
:::

The PostgreSQL destination utilizes batching, and supports [`batch_size`](/docs/cli/integrations/destinations#batch_size) and [`batch_size_bytes`](/docs/cli/integrations/destinations#batch_size_bytes).

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

- `lakebase` (`object`) (optional)

  Configuration to connect to [Databricks Lakebase](https://docs.databricks.com/aws/en/oltp), a PostgreSQL-compatible managed database. When set, the plugin uses the Databricks SDK to generate a short-lived OAuth database credential before each new connection and uses it as the connection password. The `connection_string` still supplies the host, port, database name and user (the service principal client ID), and must use `sslmode=require` (or `verify-ca`/`verify-full`); TLS is required and enforced. See the [Databricks Lakebase example](#databricks-lakebase) below.

  - `endpoint` (`string`) (required)

    The Lakebase database endpoint resource name, in the format `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`.

  - `host` (`string`) (optional)

    Databricks workspace host, for example `https://your-workspace.cloud.databricks.com`. If empty, the Databricks SDK resolves it from the `DATABRICKS_HOST` environment variable (or other default Databricks configuration sources).

  - `client_id` (`string`) (optional)

    Databricks service principal OAuth client ID. If empty, the Databricks SDK resolves it from the `DATABRICKS_CLIENT_ID` environment variable (or other default Databricks configuration sources).

  - `client_secret` (`string`) (optional)

    Databricks service principal OAuth client secret. If empty, the Databricks SDK resolves it from the `DATABRICKS_CLIENT_SECRET` environment variable (or other default Databricks configuration sources).

- `pgvector_config` (`object`) (optional)

  Optional configuration to enable PgVector embedding support.

  Note: source plugin must sync the `_cq_id` column on target tables if this is enabled.

  - `tables` (`array`) (required)
    
    Tables to create embeddings for. For each entry, embeddings are created from a source table and stored in a configured target table.

    - `source_table_name` (`string`) (required)

      Name of the source table from which text columns are read to generate embeddings.

    - `target_table_name` (`string`) (required)

      Name of the embeddings table to create/populate. This table will contain the `embedding` vector column, a `chunk` text column, and the configured metadata columns. The `_cq_id` column is always included and indexed.

    - `embed_columns` (`array`) (required)

      Columns on the source table to concatenate and create embeddings for.

    - `metadata_columns` (`array`) (optional)

      These columns will be added as-is from the source table for context. The `_cq_id` column will be added automatically and an index will be created on it.

  - `text_splitter` (`object`) (optional)

    Optional text splitting configuration for the embeddings. If unset, defaults are used.

    - `recursive_text` (`object`) (required)

      - `chunk_size` (`integer`) (required)

      Chunk size for the text splitting.

      - `chunk_overlap` (`integer`) (required)

      Chunk overlap for the text splitting.

  - `openai_embedding` (`object`) (required)

    OpenAI Embedding API configuration. Currently only OpenAI is supported.

    - `dimensions` (`integer`) (required)
    
    The number of dimensions to use for the embeddings. For `text-embedding-3-small`, this is 1536. For `text-embedding-3-large`, this is 3072.

    - `api_key` (`string`) (required)

    The OpenAI API key to use for the embedding API.

    - `model_name` (`string`) (required)

    The model name to use for the embedding API. Currently, `text-embedding-3-small` and `text-embedding-3-large` are supported.
  
  - `retry_on_deadlock` (`integer`) (optional) (default: `0`)
  
    Number of times to retry a transaction if a deadlock is detected by Postgres (Postgres error code `40P01`).

### Databricks Lakebase

[Databricks Lakebase](https://docs.databricks.com/aws/en/oltp) is a PostgreSQL-compatible managed database. To sync into it, set the `lakebase` block. The plugin authenticates with a Databricks service principal (OAuth M2M) and mints a fresh short-lived database credential for every new connection, so you do not put a password in the `connection_string`. The connection string supplies the Lakebase host, port, database, user (the service principal client ID) and must use `sslmode=require`.

```yaml copy
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  registry: cloudquery
  version: "VERSION_DESTINATION_POSTGRESQL"
  write_mode: "overwrite-delete-stale"
  send_sync_summary: true
  spec:
    # No password in the connection string - it is generated per-connection from the Databricks credentials below.
    connection_string: "host=${PGHOST} port=5432 dbname=databricks_postgres user=${DATABRICKS_CLIENT_ID} sslmode=require"
    lakebase:
      # projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
      endpoint: "${LAKEBASE_ENDPOINT_NAME}"
      host: "${DATABRICKS_HOST}"
      client_id: "${DATABRICKS_CLIENT_ID}"
      client_secret: "${DATABRICKS_CLIENT_SECRET}"
```

`host`, `client_id` and `client_secret` are optional. If omitted, the Databricks SDK resolves them from the standard `DATABRICKS_HOST`, `DATABRICKS_CLIENT_ID` and `DATABRICKS_CLIENT_SECRET` environment variables (or other [default Databricks configuration sources](https://docs.databricks.com/aws/en/dev-tools/auth)).

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
  send_sync_summary: true
  spec:
    connection_string: ${PG_CONNECTION_STRING}
    pgx_log_level: debug # Available: error, warn, info, debug, trace. Default: "error"
```
