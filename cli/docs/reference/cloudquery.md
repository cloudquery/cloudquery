---
title: "cloudquery"
---
## cloudquery

CloudQuery CLI

### Synopsis

CloudQuery CLI

High performance data integration at scale.

Find more information at:
	https://www.cloudquery.io

### Options

```
      --cq-dir string            directory to store cloudquery files, such as downloaded plugins (default ".cq")
  -h, --help                     help for cloudquery
      --invocation-id uuid       useful for when using Open Telemetry integration for tracing and logging to be able to correlate logs and traces through many services (default <NEW-RANDOM-UUID>)
      --log-console              enable console logging
      --log-file-name string     Log filename (default "cloudquery.log")
      --log-file-overwrite       Overwrite log file on each run instead of appending. Use this if your filesystem does not support append mode (e.g. FUSE-mounted cloud storage).
      --log-format string        Logging format (json, text) (default "text")
      --log-level string         Logging level (trace, debug, info, warn, error) (default "info")
      --no-log-file              Disable logging to file
      --telemetry-level string   Telemetry level (none, errors, stats, all) (default "all")
```

### SEE ALSO

* [cloudquery addon](/cli/cli-reference/cloudquery_addon)	 - Addon commands
* [cloudquery init](/cli/cli-reference/cloudquery_init)	 - Generate a configuration file for a sync
* [cloudquery login](/cli/cli-reference/cloudquery_login)	 - Login to CloudQuery Hub.
* [cloudquery logout](/cli/cli-reference/cloudquery_logout)	 - Log out of CloudQuery Hub.
* [cloudquery migrate](/cli/cli-reference/cloudquery_migrate)	 - Update schema of your destinations based on the latest changes in sources from your configuration
* [cloudquery plugin](/cli/cli-reference/cloudquery_plugin)	 - Plugin commands
* [cloudquery switch](/cli/cli-reference/cloudquery_switch)	 - Switches between teams.
* [cloudquery sync](/cli/cli-reference/cloudquery_sync)	 - Sync resources from configured source plugins to destinations
* [cloudquery tables](/cli/cli-reference/cloudquery_tables)	 - Generate documentation for all supported tables of source plugins specified in the spec(s)
* [cloudquery test-connection](/cli/cli-reference/cloudquery_test-connection)	 - Test plugin connections to sources and/or destinations
* [cloudquery validate-config](/cli/cli-reference/cloudquery_validate-config)	 - Validate config


## See Also

- [Getting Started](/cli/getting-started) - Install and run your first sync
- [Configuration Guide](/cli/core-concepts/configuration) - Configure source and destination integrations
