---
title: "sync"
---
# cloudquery sync

Sync resources from configured source plugins to destinations

## Synopsis

Sync resources from configured source plugins to destinations

## sync_mode: events

A source can be set to `sync_mode: events` instead of the default `sync_mode: full`. In events mode the CLI opens a change stream on the source, runs a full sync only when it has to, and then refreshes just the resources the source reports as changed, in batches. It acknowledges a batch only after every destination has written it, so an interrupted run replays the batch instead of losing it. Stale rows are deleted only after a full sync. `backend_options` is required, because the source keeps its change-stream checkpoint there.

Tune it under `events`: `follow` (keep running instead of exiting once caught up), `full_sync` (`always`, `on_first_run`, `never`), `max_checkpoint_age` (force a full sync when the checkpoint is older than this), `batch_max_changes`, `batch_max_wait` and `summary_interval`. Sources that do not implement the change stream fail with an error.

```yaml
kind: source
spec:
  name: aws
  path: cloudquery/aws
  registry: cloudquery
  version: "v28.0.0"
  tables: ["aws_ec2_instances"]
  destinations: [postgresql]
  sync_mode: events
  backend_options:
    table_name: cq_state_aws
    connection: "@@plugins.postgresql.connection"
  events:
    follow: true
    full_sync: on_first_run
    batch_max_changes: 500
    batch_max_wait: 5s
    summary_interval: 5m
```

## Flag Details

### --summary-location

When set, a JSON summary of each sync is appended to the specified file after the sync completes. The file uses JSONL format (one JSON object per line). When syncing to multiple destinations, a separate entry is written for each destination.

Example usage:

```bash
cloudquery sync config.yml --summary-location summary.jsonl
```

The summary contains the following fields:

| Field                    | Type     | Description                                                       |
| ------------------------ | -------- | ----------------------------------------------------------------- |
| `cli_version`            | string   | CloudQuery CLI version                                            |
| `sync_id`                | string   | Unique sync identifier (from `--invocation-id` or auto-generated) |
| `sync_time`              | string   | Sync start time (RFC 3339)                                        |
| `sync_duration_ms`       | number   | Sync duration in milliseconds                                     |
| `resources`              | number   | Total resources synced                                            |
| `source_name`            | string   | Source integration name                                           |
| `source_path`            | string   | Source integration path                                           |
| `source_version`         | string   | Source integration version                                        |
| `source_errors`          | number   | Errors from the source integration                                |
| `source_warnings`        | number   | Warnings from the source integration                              |
| `source_tables`          | string[] | Tables synced from source                                         |
| `destination_name`       | string   | Destination integration name                                      |
| `destination_path`       | string   | Destination integration path                                      |
| `destination_version`    | string   | Destination integration version                                   |
| `destination_errors`     | number   | Errors from the destination integration                           |
| `destination_warnings`   | number   | Warnings from the destination integration                         |
| `sync_group_id`          | string   | Rendered sync group ID (if configured)                            |
| `shard_num`              | number   | Shard number (if using `--shard`)                                 |
| `shard_total`            | number   | Total shards (if using `--shard`)                                 |
| `resources_per_table`    | object   | Resource count per table                                          |
| `errors_per_table`       | object   | Error count per table                                             |
| `durations_per_table_ms` | object   | Duration per table in milliseconds                                |

### --invocation-id

A UUID that uniquely identifies a sync invocation. If not provided, a random UUID is automatically generated.

This flag serves three purposes:

1. **OpenTelemetry correlation**: The UUID is attached to all logs and traces, allowing you to correlate CLI activity with integration activity in your [monitoring setup](/cli/managing-cloudquery/monitoring/overview).
2. **Sync summary**: The UUID is stored as the `sync_id` field in the [sync summary](#--summary-location).
3. **`sync_group_id` template**: When a destination's [`sync_group_id`](/cli/integrations/destinations#sync_group_id) uses the `{{SYNC_ID}}` placeholder, it is replaced with this UUID at runtime.

Example: using a fixed invocation ID for repeatable tracing:

```bash
cloudquery sync config.yml --invocation-id 550e8400-e29b-41d4-a716-446655440000
```

```
cloudquery sync [files or directories] [flags]
```

## Examples

```
# Sync resources from configuration in a directory
cloudquery sync ./directory
# Sync resources from directories and files
cloudquery sync ./directory ./aws.yml ./pg.yml
# Log tables metrics to a file
cloudquery sync ./directory ./aws.yml ./pg.yml --tables-metrics-location metrics.txt
# Shard the sync process into 4 shards and run the first shard
cloudquery sync spec.yml --shard 1/4

```

## Options

```
  -h, --help                             help for sync
      --license string                   Set offline license file.
      --no-migrate                       Disable auto-migration before sync. By default, sync runs a migration before syncing resources.
      --shard string                     Allows splitting the sync process into multiple shards. This feature is in Preview. Please provide feedback to help us improve it. For a list of supported plugins visit https://www.cloudquery.io/docs/cli/managing-cloudquery/running-in-parallel
      --summary-location string          Sync summary file location.
      --tables-metrics-location string   Tables metrics file location. This feature is in Preview. Please provide feedback to help us improve it. Works with plugins released on 2024-07-10 or later.
```

## Options inherited from parent commands

```
      --cq-dir string            directory to store cloudquery files, such as downloaded plugins (default ".cq")
      --invocation-id uuid       useful for when using Open Telemetry integration for tracing and logging to be able to correlate logs and traces through many services (default <NEW-RANDOM-UUID>)
      --log-console              enable console logging
      --log-file-name string     Log filename (default "cloudquery.log")
      --log-file-overwrite       Overwrite log file on each run instead of appending. Use this if your filesystem does not support append mode (e.g. FUSE-mounted cloud storage).
      --log-format string        Logging format (json, text) (default "text")
      --log-level string         Logging level (trace, debug, info, warn, error) (default "info")
      --no-log-file              Disable logging to file
      --telemetry-level string   Telemetry level (none, errors, stats, all) (default "all")
```

## See Also

* [cloudquery](/cli/cli-reference/cloudquery)	 - CloudQuery CLI

- [Syncs](/cli/core-concepts/syncs) - Understand full and incremental sync modes
- [Configuration Guide](/cli/core-concepts/configuration) - Set up sync configurations
- [Performance Tuning](/cli/advanced/performance-tuning) - Optimize sync performance
