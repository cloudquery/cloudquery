---
title: "login"
---
## cloudquery login

Login to CloudQuery Hub.

### Synopsis

Login to CloudQuery Hub.

This is required to download plugins from CloudQuery Hub.

Local plugins and different registries don't need login.


```
cloudquery login [flags]
```

### Examples

```

# Log in to CloudQuery Hub
cloudquery login

# Log in to a specific team
cloudquery login --team my-team

```

### Options

```
  -h, --help          help for login
  -t, --team string   Team to login to. Specify the team name, e.g. 'my-team' (not the display name)
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

