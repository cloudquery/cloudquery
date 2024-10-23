---
title: Running Plugins Locally
description: Learn how to run custom-built or downloaded CloudQuery integrations locally.
---

# Running Plugins Locally

Plugins are normally invoked as separate processes inside CloudQuery CLI. However, for development purposes, it is possible to run integrations directly from the command line.

You can run a single integration or multiple integrations locally. It's also possible to have some integrations (from other registries such as `github`) managed and ran by the CloudQuery CLI, and have some running locally simultaneously. 

## Required Settings

You can run an integration locally yourself and tell the CLI to connect to it, or you can tell the CLI to run the integration locally from a filesystem location for you.

### Getting the CLI to run your binary

Set the `registry` in the spec file to be `local`. `path` then becomes file path to the local binary:

```yaml copy
kind: source
spec:
  name: "cloudwidgets"
  registry: "local"
  path: "/home/user/path/to/plugin/binary"
  tables: ['*']
# other settings like tables, etc.
```

In this mode, the CLI will run the binary for you and connect to it.

### Running the integration yourself

This is useful if you want to run the integration in a debugger, or if you want to run the integration in a different way than the CLI would run it.
First of all, run your integration with the `serve` argument:

```bash copy
/path/to/plugin serve
```

If you are running multiple integrations this way simultaneously, you will need to specify a different port for each one. You can do this with the `--address` flag:

```bash copy
/path/to/plugin serve --address localhost:7778
```

> `1:16PM INF Source plugin server listening address=127.0.0.1:7778`

After the integration is running, you can tell the CLI to connect to it by setting the `registry` to `grpc` and the `path` to the listen address of the integration:

```yaml copy
kind: source
spec:
  name: "cloudwidgets"
  registry: "grpc"
  path: "localhost:7778"
  tables: ['*']
# other settings like tables, etc.
```

When you run CloudQuery CLI with this configuration , and it will connect to the integration as specified.
