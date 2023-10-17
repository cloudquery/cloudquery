This is a basic configuration that will save all your sync resources to `/path/to/example.db`.

```yaml copy
kind: destination
spec:
  name: duckdb
  path: cloudquery/duckdb
  version: "VERSION_DESTINATION_DUCKDB"
  write_mode: "overwrite-delete-stale"
  spec:
    connection_string: /path/to/example.db
    # Optional parameters
    # batch_size: 1000
    # batch_size_bytes: 4194304 # 4 MiB
    # debug: false

```

After running `cloudquery sync`, you can explore the data locally with the DuckDB CLI: `duckdb /path/to/example.db`.

The default `write_mode` is `overwrite-delete-stale`, but the plugin also supports `overwrite` or `append`. Note that `overwrite` and `overwrite-delete-stale` modes do not support atomic updates: to update a resource, it is first deleted and then re-inserted. This is due to a current lack of support in DuckDB for upserting list-type columns. If this is an issue for you, consider using the `append` mode instead. You may then perform a manual cleanup of stale resources after the sync completes.

import { Callout } from 'nextra-theme-docs';

<Callout type="warning">
  Note that this plugin does currently **not support Windows**. See [this issue](https://github.com/cloudquery/cloudquery/issues/8854) for more details.
</Callout>