---
title: "provider_purge"
---
## cloudquery provider purge

Remove stale resources from one or more providers in database

### Synopsis

Remove stale resources from one or more providers in database

```
cloudquery provider purge [provider] [flags]
```

### Options

```
      --dry-run                 (default true)
  -h, --help                   help for purge
      --last-update duration   last-update is the duration from current time we want to remove resources from the database. For example 24h will remove all resources that were not update in last 24 hours. Duration is a string with optional unit suffix such as "2h45m" or "7d" (default 1h0m0s)
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

* [cloudquery provider](/docs/cli/commands/cloudquery_provider)	 - Top-level command to interact with providers.

