---
title: Managing Incremental Tables
description: Incremental tables are tables that fetch only the data that changed since the last sync. Tables that support this mode are marked as "incremental" in integration table documentation.
---

# Managing Incremental Tables

Incremental tables are tables that fetch only the data that changed since the last sync. Tables that support this mode are marked as "incremental" in integration table documentation.

When a sync runs on an incremental table, the table will first fetch the last known cursor state from the state backend, then resume syncing from that point. Incremental tables guarantee at-least-once delivery, which means that there should never be gaps in the data as a result of the cursor state being used, but there may be duplicates. If the destination uses `overwrite` or `overwrite-delete-stale` write mode, these duplicates will be handled automatically. But if the destination uses `append` mode, care will need to be taken to either exclude the duplicates at query time or to run a deduplication process on the destination. 

In order to resume from a previous position, incremental tables store some state, known as the **cursor**. When using the CloudQuery CLI, the cursor state is stored in a **backend**. Any destination integration can be used as a backend. This state destination is often the same destination data is being written to for the sync, but it doesn't need to be. In some destinations that support only `append` mode, each state update will generate a new row (with a new version) in the state table and may need to be cleaned up periodically. Destinations supporting `overwrite` mode don't have this issue.

## Example Configuration

Add a `backend_options` property to the source configuration to enable the use of a state backend for incremental tables. It accepts the following properties:

 - `table_name` (`string`, required): The name of the table to store the cursor state in.
 - `connection` (`string`, required): The connection string to the destination integration. Note that this is gRPC connection string, not the connection string to the destination database itself. See below for more details.

For example, the following configuration shows how to enable the use of a PostgreSQL destination as the state backend for the AWS integration:  

```yaml
kind: source
spec:
  name: aws
  path: cloudquery/aws
  registry: cloudquery
  version: "VERSION_SOURCE_AWS"
  tables: ["aws_cloudtrail_events"]
  destinations: ["postgresql"]
  backend_options:
    table_name: "cq_state_aws"
    connection: "@@plugins.postgresql.connection"
  spec:
    # AWS integration specific configuration
    # ...
---
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_POSTGRESQL"
  write_mode: "overwrite-delete-stale"
  spec:
    connection_string: "${CONNECTION_STRING}"
```

A special `@@plugins.<plugin-name>.<property-name>` syntax is used to reference a property from another integration. In this case, the `connection` property from the PostgreSQL integration is being referenced. Note that here the `connection` refers to the gRPC connection to the destination integration, automatically inferred after the destination integration is started, not the connection string to the destination database itself.

Sometimes it may be useful to use a different state backend than the destination. For example, if the destination is a data warehouse that only supports `append` mode, a separate database can be used as the state backend. For example, the following configuration writes AWS data to BigQuery, but stores state in a local SQLite database:

```yaml
kind: source
spec:
  name: aws
  path: cloudquery/aws
  registry: cloudquery
  version: "VERSION_SOURCE_AWS"
  tables: ["aws_cloudtrail_events"]
  destinations: ["bigquery"]
  backend_options:
    table_name: "cq_state_aws"
    connection: "@@plugins.sqlite.connection"
  spec:
    # AWS integration specific configuration
    # ...
---
kind: destination
spec:
  name: bigquery
  path: cloudquery/bigquery
  registry: cloudquery
  version: "VERSION_DESTINATION_BIGQUERY"
  write_mode: "append"
  spec:
    project_id: ${PROJECT_ID}
    dataset_id: ${DATASET_ID}
---
kind: destination
spec:
  name: sqlite
  path: cloudquery/sqlite
  registry: cloudquery
  version: "VERSION_DESTINATION_SQLITE"
  spec:
    connection_string: ./db.sqlite
``` 