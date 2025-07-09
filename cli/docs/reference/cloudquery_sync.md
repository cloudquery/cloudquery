---
title: "sync"
---
## cloudquery sync

Sync resources from configured source plugins to destinations

### Synopsis

Sync resources from configured source plugins to destinations

```
cloudquery sync [files or directories] [flags]
```

### Examples

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

### Options

```
  -h, --help                             help for sync
      --license string                   Set offline license file.
      --no-migrate                       Disable auto-migration before sync. By default, sync runs a migration before syncing resources.
      --shard string                     Allows splitting the sync process into multiple shards. This feature is in Preview. Please provide feedback to help us improve it. For a list of supported plugins visit https://docs.cloudquery.io/docs/advanced-topics/running-cloudquery-in-parallel
      --summary-location string          Sync summary file location.
      --tables-metrics-location string   Tables metrics file location. This feature is in Preview. Please provide feedback to help us improve it. Works with plugins released on 2024-07-10 or later.
```

### Options inherited from parent commands

```
      --cq-dir string            directory to store cloudquery files, such as downloaded plugins (default ".cq")
      --invocation-id uuid       useful for when using Open Telemetry integration for tracing and logging to be able to correlate logs and traces through many services (default <NEW-RANDOM-UUID>)
      --log-console              enable console logging
      --log-file-name string     Log filename (default "cloudquery.log")
      --log-format string        Logging format (json, text) (default "text")
      --log-level string         Logging level (trace, debug, info, warn, error) (default "info")
      --no-log-file              Disable logging to file
      --telemetry-level string   Telemetry level (none, errors, stats, all) (default "all")
```

### SEE ALSO

* [cloudquery](/docs/reference/cli/cloudquery)	 - CloudQuery CLI

