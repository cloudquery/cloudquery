This is a basic configuration that will save all your sync resources to `./database.db` in the local directory.

```yaml copy
kind: destination
spec:
  name: duckdb
  path: cloudquery/duckdb
  registry: cloudquery
  version: "VERSION_DESTINATION_DUCKDB"
  write_mode: "overwrite-delete-stale"
  # Learn more about the configuration options at https://cql.ink/duckdb_destination
  spec:
    connection_string: ./database.db
    # Optional parameters
    # batch_size: 1000
    # batch_size_bytes: 4194304 # 4 MiB
    # debug: false

```

After running `cloudquery sync`, you can explore the data locally with the DuckDB CLI: `duckdb ./database.db`.

The default `write_mode` is `overwrite-delete-stale`, but the plugin also supports `overwrite` or `append`. Note that `overwrite` and `overwrite-delete-stale` modes do not support atomic updates: to update a resource, it is first deleted and then re-inserted. This is due to a current lack of support in DuckDB for upserting list-type columns. If this is an issue for you, consider using the `append` mode instead. You may then perform a manual cleanup of stale resources after the sync completes.

:::callout{type="info"}
Note that this plugin does currently **not support Windows**. See [this issue](https://github.com/cloudquery/cloudquery/issues/8854) for more details.
:::