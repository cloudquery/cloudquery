---
title: "init"
---
## cloudquery init

Generate a configuration file for a sync

### Synopsis

Generate a configuration file for a sync

```
cloudquery init [flags]
```

### Examples

```
# Display prompts to select source and destination plugins and generate a configuration file from them
cloudquery init
# Generate a configuration file for a sync from aws to bigquery
cloudquery init --source aws --destination bigquery
# Display a prompt to select a source plugin and generate a configuration file for a sync from it to bigquery
cloudquery init --destination bigquery
# Display a prompt to select a destination plugin and generate a configuration file for a sync from aws to it
cloudquery init --source aws
# Accept all defaults and generate a configuration file for a sync from the first source and destination plugins
cloudquery init --yes
```

### Options

```
      --destination string   Destination plugin name or path
  -h, --help                 help for init
      --source string        Source plugin name or path
      --spec-path string     Output spec file path
      --yes                  Accept all defaults
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

