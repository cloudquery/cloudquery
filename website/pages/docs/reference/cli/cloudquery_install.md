---
title: "install"
---
## cloudquery install

Install required plugin images from your configuration

### Synopsis

Install required plugin images from your configuration

```
cloudquery install [files or directories] [flags]
```

### Examples

```
# Install required plugins specified in directory
cloudquery install ./directory
# Install required plugins specified in directory and config files
cloudquery install ./directory ./aws.yml ./pg.yml

```

### Options

```
  -h, --help   help for install
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

* [cloudquery](/docs/reference/cli/cloudquery)	 - CloudQuery CLI

