---
title: "init"
---
# cloudquery init

Generate a configuration file for a sync

## Synopsis

Generate a configuration file for a sync

## Modes

The `init` command operates in one of three modes depending on your authentication state and flags:

**AI-assisted mode** (default when logged in)

Activates when you are logged in to a team (`cloudquery login`) and don't specify `--source` or `--destination`. Launches an interactive AI chat session that walks you through the setup process — selecting integrations, generating YAML configuration files, testing connections, and giving you some example queries.

Type `exit` or `quit` to end the conversation. Use `--resume-conversation` to continue a previous session instead of starting a new one.

**Basic interactive mode**

Activates when you pass `--disable-ai`, or as a fallback if the AI assistant is unavailable. Presents a searchable picker to select source and destination integrations, then generates a configuration file from their default templates.

**Non-interactive mode**

Activates when both `--source` and `--destination` are specified. Generates the configuration file directly without prompts.

Authentication via `cloudquery login` is required for AI-assisted and basic interactive modes.

```
cloudquery init [flags]
```

## Examples

```
# Display prompts to select source and destination plugins and generate a configuration file from them
cloudquery init
# Generate a configuration file for a sync from aws to bigquery
cloudquery init --source aws --destination bigquery
# Display a prompt to select a source plugin and generate a configuration file for a sync from it to bigquery
cloudquery init --destination bigquery
# Display a prompt to select a destination plugin and generate a configuration file for a sync from aws to it
cloudquery init --source aws
# Accept all defaults and generate a configuration file for a sync from the first source and destination plugins
cloudquery init --yes
```

## Options

```
      --destination string    Destination plugin name or path
      --disable-ai            Disable AI assistant
  -h, --help                  help for init
      --resume-conversation   Resume existing AI conversation instead of starting a new one
      --source string         Source plugin name or path
      --spec-path string      Output spec file path
      --yes                   Accept all defaults
```

## Options inherited from parent commands

```
      --cq-dir string            directory to store cloudquery files, such as downloaded plugins (default ".cq")
      --invocation-id uuid       useful for when using Open Telemetry integration for tracing and logging to be able to correlate logs and traces through many services (default <NEW-RANDOM-UUID>)
      --log-console              enable console logging
      --log-file-name string     Log filename (default "cloudquery.log")
      --log-file-overwrite       Overwrite log file on each run instead of appending. Use this if your filesystem does not support append mode (e.g. FUSE-mounted cloud storage).
      --log-format string        Logging format (json, text) (default "text")
      --log-level string         Logging level (trace, debug, info, warn, error) (default "info")
      --no-log-file              Disable logging to file
      --telemetry-level string   Telemetry level (none, errors, stats, all) (default "all")
```

## See Also

* [cloudquery](/cli/cli-reference/cloudquery)	 - CloudQuery CLI

- [Getting Started](/cli/getting-started) - Full quickstart guide using the init command
- [Configuration Guide](/cli/core-concepts/configuration) - Understand the generated configuration files
