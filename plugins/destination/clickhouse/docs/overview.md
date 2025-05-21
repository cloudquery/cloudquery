---
name: ClickHouse
stage: GA
title: ClickHouse Destination Plugin
description: CloudQuery ClickHouse destination plugin documentation
---
# ClickHouse destination plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to [ClickHouse](https://clickhouse.com/) database.

It supports `append` write mode only.
Write mode selection is required through [`write_mode`](/docs/reference/destination-spec#write_mode).

Supported database versions: >= `24.8.1`

## Configuration

### Example

:configuration

:::callout{type="info"}
Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.
:::

### ClickHouse spec

This is the (nested) spec used by the ClickHouse destination plugin.

- `connection_string` (`string`) (required)

  Connection string to connect to the database.
  See [SDK documentation](https://github.com/ClickHouse/clickhouse-go#dsn) for more details.

  Example connection string:

  - `"clickhouse://username:password@host1:9000,host2:9000/database?dial_timeout=200ms&max_execution_time=60"`

- `cluster` (`string`) (optional) (default: not used)

  Cluster name to be used for [distributed DDL](https://clickhouse.com/docs/en/sql-reference/distributed-ddl).
  If the value is empty, DDL operations will affect only the server the plugin is connected to.

- `ca_cert` (`string`) (optional) (default: not used)

  PEM-encoded certificate authorities.
  When set, a certificate pool will be created by appending the certificates to the system pool.

  See [file variable substitution](/docs/advanced-topics/environment-variable-substitution#file-variable-substitution-example)
  for how to read this value from a file.

- `engine` (optional, [table engine settings](#clickhouse-table-engine). Default: `MergeTree` engine)

  Engine to be used for tables.
  Only [`*MergeTree` family](https://clickhouse.com/docs/en/engines/table-engines/mergetree-family) is supported at the moment.

- `batch_size` (`integer`) (optional) (default: `10000`)

  Maximum number of items that may be grouped together to be written in a single write.

- `batch_size_bytes` (`integer`) (optional) (default: `5242880` (= 5 MiB))

  Maximum size of items that may be grouped together to be written in a single write.

- `batch_timeout` (`duration`) (optional) (default: `20s`)

  Maximum interval between batch writes.

- `partition` (optional, [partitioning](#partitioning)) (default: no partitioning)

  Partitioning strategy to be used for tables (i.e. `PARTITION BY` clause in `CREATE TABLE` statements).

- `order` (optional, [ordering](#ordering)) (default: use existing primary key)

  Ordering strategy to be used for tables (i.e. `ORDER BY` clause in `CREATE TABLE` statements).

#### ClickHouse table engine

This option allows to specify a custom table engine to be used.

- `name` (`string`) (required)

  Name of the table engine.
  Only [`*MergeTree` family](https://clickhouse.com/docs/en/engines/table-engines/mergetree-family) is supported at the moment.

- `parameters` (array of parameters) (optional) (default: empty)

  Engine parameters.
  Currently, no restrictions are imposed on the parameter types.

```yaml copy
kind: destination
spec:
  name:       "clickhouse"
  path:       "cloudquery/clickhouse"
  registry:   "cloudquery"
  version:    "VERSION_DESTINATION_CLICKHOUSE"
  write_mode: "append"

  spec:
    connection_string: "clickhouse://${CH_USER}:${CH_PASSWORD}@localhost:9000/${CH_DATABASE}"
    engine:
      name: ReplicatedMergeTree
      parameters:
      - "/clickhouse/tables/{shard}/{database}/{table}"
      - "{replica}"
```

#### Partitioning

This option allows to specify a partitioning strategy to be used for tables. It is an array of objects.

Each object has the following fields:

- `tables` (array of strings) (optional) (default: `["*"]`)

  List of glob patterns to match table names against. Follows the same rules as the top-level spec `tables` option.

  If a table matches both a pattern in `tables` and `skip_tables`, the table will be skipped.

  Partition strategy table patterns should be disjointed sets: if a table matches two partition strategies, an error will be raised at runtime.

- `skip_tables` (array of strings) (optional) (default: empty)

  List of glob patterns to skip matching table names against. Follows the same rules as the top-level spec `skip_tables` option.

  If a table matches both a pattern in `tables` and `skip_tables`, the table will be skipped.

  Partition strategy table patterns should be disjointed sets: if a table matches two partition strategies, an error will be raised at runtime.

- `skip_incremental_tables` (boolean) (optional) (default: `false`)

  If set to `true`, incremental tables will not be partitioned by this strategy.

- `partition_by` (string) (required)

  Partitioning strategy to use, e.g. `toYYYYMM(_cq_sync_time)`, the string is passed as is after "PARTITION BY" clause with no validation or quoting.

  An unset `partition_by` is not valid.

Example:

```yaml copy
partition:
- tables: ["*"]
  skip_tables: ["special_partition_table", "non_partitioned_table"]
  partition_by: "toYYYYMM(_cq_sync_time)"
- tables: ["special_partition_table"]
  partition_by: "toYYYYMMDD(_cq_sync_time)"
```

#### Ordering

This option allows to specify custom `ORDER BY` clauses for tables or groups of tables. It is an array of objects.

Each object has the following fields:

- `tables` (array of strings) (optional) (default: `["*"]`)

  List of glob patterns to match table names against. Follows the same rules as the top-level spec `tables` option.

  If a table matches both a pattern in `tables` and `skip_tables`, the table will be skipped.

  Ordering strategy table patterns should be disjointed sets: if a table matches two ordering strategies, an error will be raised at runtime.

- `skip_tables` (array of strings) (optional) (default: empty)

  List of glob patterns to skip matching table names against. Follows the same rules as the top-level spec `skip_tables` option.

  If a table matches both a pattern in `tables` and `skip_tables`, the table will be skipped.

  Ordering strategy table patterns should be disjointed sets: if a table matches two ordering strategies, an error will be raised at runtime.

- `order_by` (array of strings) (required)

  Sort key to use, the strings are passed as is after "ORDER BY" clause with no validation or quoting.

Example:

```yaml copy
order:
- tables: ["aws_ec2_instances"]
  order_by:
  - "`account_id`"
  - "`region`"
  - "toYYYYMM(`_cq_sync_time`) DESC"
  - "`_cq_id`"
```

### Connecting to ClickHouse Cloud

To connect to [ClickHouse Cloud](https://clickhouse.com/cloud), you need to set the `secure=true` parameter, username is `default`, and the port is `9440`. Use a connection string similar to:

```yaml copy
    connection_string: "clickhouse://default:${CH_PASSWORD}@<your-server-id>.<region>.<provider>.clickhouse.cloud:9440/${CH_DATABASE}?secure=true"
```

See [Quick Start: Using the ClickHouse Client](https://clickhouse.com/docs/en/cloud-quick-start#5-using-the-clickhouse-client) for more details.

#### Verbose logging for debug

The ClickHouse destination can be run in debug mode.
To achieve this pass the `debug=true` option to `connection_string`.
See [SDK documentation](https://github.com/ClickHouse/clickhouse-go#dsn) for more details.

Note: This will use [SDK](https://github.com/ClickHouse/clickhouse-go) built-in logging
and might output data and sensitive information to logs.
Make sure not to use it in production environment.

```yaml copy
kind: destination
spec:
  name:       "clickhouse"
  path:       "cloudquery/clickhouse"
  registry:   "cloudquery"
  version:    "VERSION_DESTINATION_CLICKHOUSE"
  write_mode: "append"

  spec:
    connection_string: "clickhouse://${CH_USER}:${CH_PASSWORD}@localhost:9000/${CH_DATABASE}?debug=true"
```
