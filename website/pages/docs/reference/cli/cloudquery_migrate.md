---
title: "migrate"
---
## cloudquery migrate

Update schema of your destinations based on the latest changes in sources from your configuration

### Synopsis

Update schema of your destinations based on the latest changes in sources from your configuration

```
cloudquery migrate [files or directories] [flags]
```

### Examples

```
# Run migration for plugins specified in directory
cloudquery migrate ./directory
# Run migration for plugins specified in directory and config files
cloudquery migrate ./directory ./aws.yml ./pg.yml

```

### Options

```
  -h, --help   help for migrate
```

### Options inherited from parent commands

```
      --cq-dir string            directory to store cloudquery files, such as downloaded plugins (default ".cq")
      --log-console              enable console logging
      --log-file-name string     Log filename (default "cloudquery.log")
      --log-format string        Logging format (json, text) (default "text")
      --log-level string         Logging level (default "info")
      --no-log-file              Disable logging to file
      --telemetry-level string   Telemetry level (none, errors, stats, all) (default "all")
```

### SEE ALSO

* [cloudquery](/docs/reference/cli/cloudquery)	 - CloudQuery CLI

