---
title: "plugin_spec-schema"
---
# cloudquery plugin spec-schema

Export a plugin's spec JSON schema.

## Synopsis

Export a plugin's spec JSON schema to a local file.
The exported file can later be passed to `cloudquery validate-config --schemas-dir` to validate
configurations fully offline, without spawning the plugin binary or contacting the CloudQuery registry.

```
cloudquery plugin spec-schema <team_name>/<plugin_kind>/<plugin_name>@<version> [flags]
```

## Examples

```

# Print schema to stdout
cloudquery plugin spec-schema cloudquery/source/aws@v33.0.0

# Write schema to a specific file
cloudquery plugin spec-schema cloudquery/source/aws@v33.0.0 -o aws.json

# Write schema to <dir>/<name>.json (suitable for --schemas-dir consumption)
cloudquery plugin spec-schema cloudquery/source/aws@v33.0.0 -D ./schemas
```

## Options

```
  -h, --help                 help for spec-schema
  -o, --output string        Write schema to this file. Mutually exclusive with --schemas-dir.
  -D, --schemas-dir string   Write schema to <dir>/<plugin-name>.json. Mutually exclusive with --output.
```

## Options inherited from parent commands

```
      --cq-dir string            directory to store cloudquery files, such as downloaded plugins (default ".cq")
      --invocation-id uuid       useful for when using Open Telemetry integration for tracing and logging to be able to correlate logs and traces through many services (default <NEW-RANDOM-UUID>)
      --log-console              enable console logging
      --log-file-name string     Log filename (default "cloudquery.log")
      --log-file-overwrite       Overwrite log file on each run instead of appending. Use this if your filesystem does not support append mode (e.g. FUSE-mounted cloud storage).
      --log-format string        Logging format (json, text) (default "text")
      --log-level string         Logging level (trace, debug, info, warn, error) (default "info")
      --no-log-file              Disable logging to file
      --telemetry-level string   Telemetry level (none, errors, stats, all) (default "all")
```

## See Also

* [cloudquery plugin](/cli/cli-reference/cloudquery_plugin)	 - Plugin commands

