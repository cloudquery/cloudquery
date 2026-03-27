---
title: "switch"
---
## cloudquery switch

Switches between teams.

### Synopsis

Switches between teams.

```
cloudquery switch [team] [flags]
```

### Examples

```

# Switch to a different team
cloudquery switch my-team

```

### Options

```
  -h, --help   help for switch
```

### Options inherited from parent commands

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

### SEE ALSO

* [cloudquery](/cli/cli-reference/cloudquery)	 - CloudQuery CLI

