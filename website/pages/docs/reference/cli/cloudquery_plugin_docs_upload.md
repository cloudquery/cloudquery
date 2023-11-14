---
title: "plugin_docs_upload"
---
## cloudquery plugin docs upload

Upload plugin docs to CloudQuery Hub.

### Synopsis

Upload plugin docs to CloudQuery Hub.

This uploads documentation for a specific plugin version from a local docs directory to CloudQuery hub.


```
cloudquery plugin docs upload [-D docs] [--sync] <team_name>/<plugin_name>@<version> [flags]
```

### Examples

```

# Upload plugin docs to CloudQuery Hub
cloudquery plugin docs upload test-team/test-plugin@v1.0.0
```

### Options

```
  -D, --docs-dir string   Path to the docs directory (default "docs")
  -h, --help              help for upload
      --sync              Syncronize docs with CloudQuery Hub, deleting any docs that are not present locally
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

