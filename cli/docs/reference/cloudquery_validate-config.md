---
title: "validate-config"
---
# cloudquery validate-config

Validate config

## Synopsis

Validate configuration without running a sync.

For `registry: cloudquery` plugins, the spec JSON schema is fetched from
the CloudQuery Hub API (https://api.cloudquery.io). This avoids downloading
the plugin binary and works for public plugins without authentication; if a
CloudQuery API token is available (via login or CLOUDQUERY_API_KEY) it is
propagated so private plugins resolve too.

For other registries (`local`, `grpc`, `docker`) the plugin is still spawned
locally to obtain its schema, identical to the previous behaviour. The tables
list is not validated against the source — this validation is stricter than
the validation done during `sync`, so a config passing here will also pass
sync's validation.

```
cloudquery validate-config [files or directories] [flags]
```

## Examples

```
# Validate configs
cloudquery validate-config ./directory
# Validate configs from directories and files
cloudquery validate-config ./directory ./aws.yml ./pg.yml

```

## Options

```
  -h, --help                                help for validate-config
      --license cloudquery sync --license   Set offline license file. When provided, the Hub API is bypassed and plugins are spawned locally (mirrors cloudquery sync --license).
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

* [cloudquery](/cli/cli-reference/cloudquery)	 - CloudQuery CLI

- [Configuration Guide](/cli/core-concepts/configuration) - Configuration format and options
- [Environment Variables](/cli/managing-cloudquery/environment-variables) - Variable substitution in configuration files
