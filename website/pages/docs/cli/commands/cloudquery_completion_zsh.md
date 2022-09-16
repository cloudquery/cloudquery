---
title: "completion_zsh"
---
## cloudquery completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(cloudquery completion zsh); compdef _cloudquery cloudquery

To load completions for every new session, execute once:

#### Linux:

	cloudquery completion zsh > "${fpath[1]}/_cloudquery"

#### macOS:

	cloudquery completion zsh > $(brew --prefix)/share/zsh/site-functions/_cloudquery

You will need to start a new shell for this setup to take effect.


```
cloudquery completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
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

