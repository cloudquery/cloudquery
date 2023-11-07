---
title: "addon_publish"
---
## cloudquery addon publish

Publish to CloudQuery Hub.

### Synopsis

Publish to CloudQuery Hub.

This publishes an addon version to CloudQuery Hub from a manifest file and directory.


```
cloudquery addon publish manifest.json /path/to/directory v1.0.0 [--finalize] [flags]
```

### Examples

```

# Publish an addon version from a manifest file and directory
cloudquery addon publish /path/to/manifest.json /path/to/addon-dir v1.0.0
```

### Options

```
  -f, --finalize   Finalize the addon version after publishing. If false, the addon version will be marked as draft=true.
  -h, --help       help for publish
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

* [cloudquery addon](/docs/reference/cli/cloudquery_addon)	 - Addon commands

