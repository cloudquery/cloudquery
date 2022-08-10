---
id: "provider_drop"
hide_title: true
sidebar_label: "provider drop"
---
## cloudquery provider drop

Drops provider schema from database

### Synopsis

Drops provider schema from database

```
cloudquery provider drop [provider] [flags]
```

### Options

```
      --force   Really drop tables for the provider
  -h, --help    help for drop
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

* [cloudquery provider](cloudquery_provider.md)	 - Top-level command to interact with providers.

