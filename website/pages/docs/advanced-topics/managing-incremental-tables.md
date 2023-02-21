# Managing Incremental Tables

Incremental tables are tables that fetch only the data that changed since the last sync. Tables that support this mode are marked as "incremental" in plugin table documentation.

When a sync runs on an incremental table, the table will first fetch the last known cursor state from the state backend, then resume syncing from that point. Incremental tables guarantee at-least-once delivery, which means that there should never be gaps in the data as a result of the cursor state being used, but there may be duplicates. If the destination uses `overwrite` or `overwrite-delete-stale` write mode, these duplicates will be handled automatically. But if the destination uses `append` mode, care will need to be taken to either exclude the duplicates at query time or to run a deduplication process on the destination. 

In order to resume from a previous position, incremental tables store some state, known as the **cursor**. When using the CloudQuery CLI, the cursor state is stored in a **backend**. Currently, only one backend is supported: a local backend that writes state to the local filesystem as JSON files.

## Local Backend

The local backend stores cursor state as JSON files on the local filesystem. The default location is the `.cq/state` directory, but this can be configured by using the `path` option inside the `backend_spec` section of the [source plugin configuration](/docs/reference/source-spec#backend_spec).

When syncing incremental tables with the local backend, special care needs to be taken to ensure that the state is preserved between syncs. For example, if you are running CloudQuery in a Docker container, you will need to mount the state directory as a volume so that it is persisted between container restarts. Otherwise, incremental tables will not be able to benefit from stored cursors and will need to re-sync all data on every run. 

## Other Backends

While the CloudQuery CLI currently only supports the local backend, we intend on adding support for more backends in the near future. If you would like to see support for a Redis backend added, please 👍 [this issue](https://github.com/cloudquery/cloudquery/issues/6630) (or volunteer to implement it!). Or if there is another state backend you'd like to use, feel free to [open an issue on GitHub](https://github.com/cloudquery/cloudquery/issues).  