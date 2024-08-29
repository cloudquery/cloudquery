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
      --log-format string        Logging format (json, text) (default "text")
      --log-level string         Logging level (trace, debug, info, warn, error) (default "info")
      --no-log-file              Disable logging to file
      --telemetry-level string   Telemetry level (none, errors, stats, all) (default "all")
```

### SEE ALSO

* [cloudquery addon](/docs/reference/cli/cloudquery_addon)	 - Addon commands
* [cloudquery init](/docs/reference/cli/cloudquery_init)	 - Generate a configuration file for a sync
* [cloudquery login](/docs/reference/cli/cloudquery_login)	 - Login to CloudQuery Hub.
* [cloudquery logout](/docs/reference/cli/cloudquery_logout)	 - Log out of CloudQuery Hub.
* [cloudquery migrate](/docs/reference/cli/cloudquery_migrate)	 - Update schema of your destinations based on the latest changes in sources from your configuration
* [cloudquery plugin](/docs/reference/cli/cloudquery_plugin)	 - Plugin commands
* [cloudquery switch](/docs/reference/cli/cloudquery_switch)	 - Switches between teams.
* [cloudquery sync](/docs/reference/cli/cloudquery_sync)	 - Sync resources from configured source plugins to destinations
* [cloudquery tables](/docs/reference/cli/cloudquery_tables)	 - Generate documentation for all supported tables of source plugins specified in the spec(s)
* [cloudquery test-connection](/docs/reference/cli/cloudquery_test-connection)	 - Test plugin connections to sources and/or destinations
* [cloudquery validate-config](/docs/reference/cli/cloudquery_validate-config)	 - Validate config

