---
title: "completion_fish"
---
## cloudquery completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	cloudquery completion fish | source

To load completions for every new session, execute once:

	cloudquery completion fish > ~/.config/fish/completions/cloudquery.fish

You will need to start a new shell for this setup to take effect.


```
cloudquery completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --color string           Enable colorized output (on, off, auto) (default "auto")
      --data-dir string        set persistent data directory (env: CQ_DATA_DIR) (default "./.cq")
      --log-console            enable console logging
      --log-file-name string   Log filename (default "cloudquery.log")
      --log-format string      Logging format (json, text) (default "text")
      --log-level string       Logging level (default "info")
      --no-log-file            Disable logging to file
      --no-telemetry           disable telemetry collection
```

### SEE ALSO

* [cloudquery completion](/docs/cli/commands/cloudquery_completion)	 - Generate the autocompletion script for the specified shell

