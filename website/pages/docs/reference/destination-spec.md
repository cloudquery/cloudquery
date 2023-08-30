---
title: Destination Spec Reference
description: Reference for the destination spec CloudQuery configuration object.
---

# Destination Spec Reference

This goes through all the available options for the destination plugin `spec` object.

## Example

This example configures the `postgresql` plugin to connect to a PostgreSQL database located at `localhost:5432`.

```yaml copy
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  version: "VERSION_DESTINATION_POSTGRESQL"

  spec:
    connection_string: "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
```

## Spec

### name

(`string`, required)

Name of the plugin. If you have multiple destination plugins, this must be unique.

The name field may be used to uniquely identify a particular destination configuration. For example, if you have two configs for the PostgreSQL plugin for syncing different databases, one may be named `db-1` and the other `db-2`. In this case, the `path` option below must be used to specify the download path for the plugin.

### registry

(`string`, optional, default: `github`, available: `github`, `local`, `grpc`)

- `github`: CloudQuery will look for and download the plugin from GitHub, and then execute it.
- `local`: CloudQuery will execute the plugin from a local path.
- `grpc`: mostly useful in debug mode when plugin is already running in a different terminal, CloudQuery will connect to the gRPC plugin server directly without spawning the process.

### path

(`string`, required)

Configures how to retrieve the plugin. The contents depend on the value of `registry` (`github` by default).

- For plugins hosted on GitHub, `path` should be of the form `"<org>/<repository>"`. For official plugins, should be `cloudquery/<plugin-name>`.
- For plugins that are located in the local filesystem, `path` should a filesystem path to the plugin binary.
- To connect to a running plugin via `grpc` (mostly useful for debugging), `path` should be the host-port of the plugin (e.g. `localhost:7777`).

### version

(`string`, required)

`version` must be a valid [SemVer](https://semver.org/)), e.g. `vMajor.Minor.Patch`. You can find all official plugin versions under [our GitHub releases page](https://github.com/cloudquery/cloudquery/releases), and for community plugins you can find it in the relevant community repository.

### write_mode

(`string`, optional, default: `overwrite-delete-stale`. Available: `overwrite-delete-stale`, `overwrite`, `append`)

Specifies the update method to use when inserting rows. The exact semantics depend on the destination plugin, and all destinations don't support all options, so check the destination plugin documentation for details.

- `overwrite-delete-stale`: `sync`s overwrite existing rows with the same primary key, and delete rows that
  are no longer present in the cloud.
- `overwrite`: Same as `overwrite-delete-stale`, but doesn't delete stale rows from previous `sync`s.
- `append`: Rows are never overwritten or deleted, only appended.

> Switching from `overwrite-delete-stale` or `overwrite` to `append`, or from `append` to `overwrite-delete-stale` or `overwrite` is not supported without dropping all tables specified in the configuration.
> To drop tables automatically, use the `migrate_mode: forced` option.

<!-- vale off -->

### migrate_mode

<!-- vale on -->

(`string`, optional, default: `safe`. Available: `safe`, `forced`)

Specifies the migration mode to use when source tables are changed. In `safe` mode (the default), CloudQuery will not run migrations that would result in data loss, and will print an error instead. In `forced` mode, CloudQuery will run migrations that may result in data loss and the migration should succeed without errors, unless a table has user created dependant objects (e.g. views).
`migrate_mode: forced` is only supported for the ClickHouse, MySQL, PostgreSQL, MSSQL and SQLite destination plugins at the moment.

Read more about how CloudQuery handles migrations [here](/docs/advanced-topics/migrations).

<!-- vale off -->

### pk_mode

<!-- vale on -->

(`string`, optional, default: `default`, Available: `default`, `cq-id-only` introduced in CLI `v2.5.2`)

Specifies the Primary Keys that the destination will configure when using the `overwrite` or `overwrite-delete-stale` mode.

- `default`: The default primary keys are used.
- `cq-id-only`: The `_cq_id` field is used as the only primary key for each table. This is useful when you don't want breaking changes to primary keys to impact your schema. It is highly recommended that if you are using this feature you should also use the [`deterministic_cq_id` feature in the source](/docs/reference/source-spec#deterministic_cq_id). If you are using `overwrite` mode and a source updates a primary key, this will result in a new row being inserted. If you are using `overwrite-delete-stale` mode, a new row will be inserted and the old row will be deleted as a stale resource. Note: using this parameter might result in changes to query performance as CloudQuery will not be creating indexes for the default primary key columns.

Supported by destination plugins released on 2023-03-21 and later

### spec

(`object`, optional)

Plugin specific configurations. Visit [destination plugins](/docs/plugins/destinations/overview) documentation for more information.

The following options are available for most destination plugins **under the nested plugin spec**:

<!-- vale off -->

#### batch_size

<!-- vale on -->

(`int`, optional)

The number of resources to insert in a single batch. Only applies to plugins that utilize batching. This setting works in conjunction with `batch_size_bytes`, and batches are written whenever either `batch_size` or `batch_size_bytes` is reached. Every plugin has its own default value for `batch_size`.

<!-- vale off -->

#### batch_size_bytes

<!-- vale on -->

(`int`, optional)

The max number of bytes to use for a single batch. Only applies to plugins that utilize batching. This setting works in conjunction with `batch_size`, and batches are written whenever either `batch_size` or `batch_size_bytes` is reached. Every plugin has its own default value for `batch_size_bytes`. Note that the size in bytes is calculated based on the size of data in memory, not the serialized data, and it is best to choose a `batch_size_bytes` significantly lower than any hard limits.
