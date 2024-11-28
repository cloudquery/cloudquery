---
title: "plugin_install"
---
## cloudquery plugin install

Install required plugin images from your configuration

### Synopsis

Install required plugin images from your configuration

```
cloudquery plugin install [files or directories] [flags]
```

### Examples

```
# Install required plugins specified in directory
cloudquery plugin install ./directory
# Install required plugins specified in directory and config files
cloudquery plugin install ./directory ./aws.yml ./pg.yml

```

### Options

```
  -h, --help   help for install
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

* [cloudquery plugin](/docs/reference/cli/cloudquery_plugin)	 - Plugin commands

