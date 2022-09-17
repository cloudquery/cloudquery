---
title: "completion_bash"
---
## cloudquery completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(cloudquery completion bash)

To load completions for every new session, execute once:

#### Linux:

	cloudquery completion bash > /etc/bash_completion.d/cloudquery

#### macOS:

	cloudquery completion bash > $(brew --prefix)/etc/bash_completion.d/cloudquery

You will need to start a new shell for this setup to take effect.


```
cloudquery completion bash
```

### Options

```
  -h, --help              help for bash
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

