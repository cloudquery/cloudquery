---
title: "addon_download"
---
## cloudquery addon download

Download addon from CloudQuery Hub.

### Synopsis

Download addon from CloudQuery Hub.

This downloads an addon from CloudQuery Hub to local disk.


```
cloudquery addon download addon-team/addon-type/addon-name@v1.0.0 [-t directory] [flags]
```

### Examples

```

# Download an addon to local disk
cloudquery addon download <publisher>/<addon-type>/<addon-name>@v1.0.0

# Further example 
cloudquery addon download cloudquery/transformation/aws-compliance-premium@v1.9.0
```

### Options

```
  -h, --help            help for download
  -t, --target string   Download to specified directory. Use - for stdout (default ".")
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

* [cloudquery addon](/docs/reference/cli/cloudquery_addon)	 - Addon commands

