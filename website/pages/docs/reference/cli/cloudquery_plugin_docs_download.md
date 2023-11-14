---
title: "plugin_docs_download"
---
## cloudquery plugin docs download

Download plugin docs from CloudQuery Hub.

### Synopsis

Download plugin docs from CloudQuery Hub.

This downloads documentation for a specific plugin version in CloudQuery hub to a local docs directory.


```
cloudquery plugin docs download [-D docs] <team_name>/<plugin_name>@<version> [flags]
```

### Examples

```

# Download plugin docs from CloudQuery Hub
cloudquery plugin docs download test-team/test-plugin@v1.0.0
```

### Options

```
  -D, --docs-dir string   Path to the docs directory (default "docs")
  -h, --help              help for download
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

* [cloudquery plugin docs](/docs/reference/cli/cloudquery_plugin_docs)	 - Plugin docs commands

