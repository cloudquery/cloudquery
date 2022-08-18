---
title: "provider"
---
## cloudquery provider

Top-level command to interact with providers.

### Synopsis

Top-level command to interact with providers.

### Examples

```

  # Sync (Upgrade or Downgrade) all providers specified in cloudquery.yml This will also create the schema.
  cloudquery provider sync 
  # Sync one or more providers
  cloudquery provider sync aws, gcp
  # Drop provider schema, running fetch again will recreate all tables unless --skip-build-tables is specified
  cloudquery provider drop aws

```

### Options

```
  -h, --help   help for provider
```

### Options inherited from parent commands

```
      --config string               path to configuration file. can be generated with 'init {provider}' command (env: CQ_CONFIG_PATH) (default "./cloudquery.yml")
      --data-dir string             set persistent data directory (env: CQ_DATA_DIR) (default "./.cq")
      --disable-log-color           disable log colors
      --dsn string                  database connection string (env: CQ_DSN) (example: 'postgres://postgres:pass@localhost:5432/postgres')
      --enable-console-log          enable console logging
      --enable-file-logging         enable file logging (default true)
      --encode-json                 enable JSON log format, instead of key/value
      --force-drop                  when upgrading schema, force dropping of any dependent views
      --log-directory string        set output directory for logs (default ".")
      --log-file string             set output filename for logs (default "cloudquery.log")
      --max-age int                 set max age in days to keep a logfile (default 3)
      --max-backups int             set max number of rolled files to keep (default 3)
      --max-size int                set max size in MB of the logfile before it's rolled (default 30)
      --no-provider-update          disable checking for new provider versions
      --no-telemetry                disable telemetry collection
      --no-verify                   disable plugins verification
      --reattach-providers string   path to reattach unmanaged plugins, mostly used for testing purposes (env: CQ_REATTACH_PROVIDERS)
      --skip-build-tables           enable skipping building tables. Should only be set if tables already exist
  -v, --verbose                     enable verbose logging
```

### SEE ALSO

* [cloudquery](/docs/cli/commands/cloudquery)	 - CloudQuery CLI
* [cloudquery provider drop](/docs/cli/commands/cloudquery_provider_drop)	 - Drops provider schema from database
* [cloudquery provider purge](/docs/cli/commands/cloudquery_provider_purge)	 - Remove stale resources from one or more providers in database
* [cloudquery provider sync](/docs/cli/commands/cloudquery_provider_sync)	 - Download the providers specified in config and re-create their database schema

