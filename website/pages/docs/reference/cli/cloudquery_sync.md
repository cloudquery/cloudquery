---
title: "sync"
---
## cloudquery sync

Sync resources from configured source plugins to destination

### Synopsis

Sync resources from configured source plugins to destination

```
cloudquery sync [directory] [flags]
```

### Examples

```
# Sync configured providers to PostgreSQL as configured in cloudquery.yml
	cloudquery sync ./directory
```

### Options

```
  -h, --help   help for sync
```

### Options inherited from parent commands

```
      --color string           Enable colorized output (on, off, auto) (default "auto")
      --data-dir string        set persistent data directory (env: CQ_DATA_DIR) (default "./.cq")
      --log-console            enable console logging
      --log-file-name string   Log filename (default "cloudquery.log")
      --log-format string      Logging format (json, text) (default "text")
      --log-level string       Logging level (default "info")
      --no-log-file            Disable logging to file
      --no-telemetry           disable telemetry collection
```

### SEE ALSO

* [cloudquery](/docs/cli/commands/cloudquery)	 - CloudQuery CLI

