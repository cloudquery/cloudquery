---
title: "generate"
---
## cloudquery generate

Generate initial config file for source and destination plugins

### Synopsis

Generate initial config file for source and destination plugins

```
cloudquery generate <source/destination> <path> [flags]
```

### Examples

```

# Downloads aws provider and writes config for aws provider to stdout
cloudquery generate source aws

# Downloads aws provider and generates initial config in aws.yml
cloudquery generate source --registry grpc --output aws.yml "localhost:7777"

```

### Options

```
  -h, --help              help for generate
  -O, --output string     destination file to write to (defaults to <name_of_plugin>.yml)
      --registry string   where to download the plugin (default "github")
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

