---
title: "validate-config"
---
## cloudquery validate-config

Validate config

### Synopsis

Validate configuration without requiring any credentials or connections. This will not validate the tables specified in the tables list. This validation is stricter than the validation done during `sync`, but if it passes this validation it will pass the sync validation.

```
cloudquery validate-config [files or directories] [flags]
```

### Examples

```
# Validate configs
cloudquery validate-config ./directory
# Validate configs from directories and files
cloudquery validate-config ./directory ./aws.yml ./pg.yml

```

### Options

```
  -h, --help   help for validate-config
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

