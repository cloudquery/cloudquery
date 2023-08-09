---
title: "tables"
---
## cloudquery tables

Generate documentation for all supported tables of source plugins specified in the spec(s)

### Synopsis

Generate documentation for all supported tables of source plugins specified in the spec(s)

```
cloudquery tables [files or directories] [flags]
```

### Examples

```
# Generate documentation for all supported tables of source plugins specified in the spec(s) 
cloudquery tables ./directory
# The default format is JSON, you can override it with --format
cloudquery tables ./directory --format markdown
# You can also specify an output directory. The default is ./cq-docs
cloudquery tables ./directory --output-dir ./docs

```

### Options

```
      --format string       Output format. One of: json, markdown (default "json")
  -h, --help                help for tables
      --output-dir string   Base output directory for generated files (default "cq-docs")
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

