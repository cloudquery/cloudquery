---
title: "upgrade_check"
---
# cloudquery upgrade check

Preview how upgrading a source plugin affects your destinations, without migrating or writing anything

## Synopsis

Preview how upgrading a source plugin affects your destinations, without migrating or writing anything

Exit codes:
  0  no action needed: no schema change, automatically migratable changes, or no output difference
  3  action needed: manual migration or rebuild, selected tables removed, or file schema or output changed
  4  unknown: a destination or source could not be assessed, or coverage is incomplete
  1  the check failed with an error

With several destinations, the highest of 0, 3 and 4 is used. The report is always printed first.

```
cloudquery upgrade check [files or directories] [flags]
```

## Examples

```
# Check how upgrading the datadog source to v6.0.0 affects the destinations in config.yml
cloudquery upgrade check ./config.yml --source datadog --to v6.0.0

# Compare a source that runs from a local binary, gRPC server or Docker image with cloudquery/semgrep v3.1.0 from the CloudQuery Hub
cloudquery upgrade check ./config.yml --source semgrep --to v3.1.0 --to-path cloudquery/semgrep

```

## Options

```
  -h, --help             help for check
      --license string   set offline license file
      --output string    Output format. One of: text, json (default "text")
      --source string    Name of the source to upgrade, as set in the configuration
      --to string        Source plugin version to upgrade to
      --to-path string   CloudQuery Hub path of the source plugin to upgrade to, as <team>/<name>. Only for sources from the local, grpc or docker registry. Default: cloudquery/<name reported by the configured source>
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

* [cloudquery upgrade](/cli/cli-reference/cloudquery_upgrade)	 - Plugin upgrade commands

