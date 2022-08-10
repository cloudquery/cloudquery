---
id: "init"
hide_title: true
sidebar_label: "init"
---
## cloudquery init

Generate initial cloudquery.yml for fetch command

### Synopsis

Generate initial cloudquery.yml for fetch command

```
cloudquery init [choose one or more providers (aws gcp azure okta ...)] [flags]
```

### Examples

```

  # Downloads aws provider and generates cloudquery.yml for aws provider
  cloudquery init aws

  # Downloads aws,gcp providers and generates one cloudquery.yml with both providers
  cloudquery init aws gcp
```

### Options

```
  -h, --help   help for init
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

* [cloudquery](cloudquery.md)	 - CloudQuery CLI

