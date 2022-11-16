# Running Plugins Locally

Plugins are normally invoked as separate processes inside CloudQuery CLI. However, for development purposes, it is possible to run plugins directly from the command line.

You can run a single plugin or multiple plugins locally. It's also possible to have some plugins (from other registries such as `github`) managed and ran by the CloudQuery CLI, and have some running locally simultaneously. 

## Required Settings

You can run a plugin locally yourself and tell the CLI to connect to it, or you can tell the CLI to run the plugin locally from a filesystem location for you.

### Getting the CLI to run your binary

Set the `registry` in the spec file to be `local`. `path` then becomes file path to the local binary:

```yaml
kind: source
spec:
  name: "cloudwidgets"
  registry: "local"
  path: "/home/user/path/to/plugin/binary"
# other settings like tables, etc.
```

In this mode, the CLI will run the binary for you and connect to it.

### Running the plugin yourself

This is useful if you want to run the plugin in a debugger, or if you want to run the plugin in a different way than the CLI would run it.
First of all, run your plugin with the `serve` argument:

```bash
/path/to/plugin serve
```

If you are running multiple plugins this way simultaneously, you will need to specify a different port for each one. You can do this with the `--address` flag:

```bash
/path/to/plugin serve --address localhost:7778
```

> `1:16PM INF Source plugin server listening address=127.0.0.1:7778`

After the plugin is running, you can tell the CLI to connect to it by setting the `registry` to `grpc` and the `path` to the listen address of the plugin:

```yaml
kind: source
spec:
  name: "cloudwidgets"
  registry: "grpc"
  path: "localhost:7778"
# other settings like tables, etc.
```

When you run CloudQuery CLI with this config it will connect to the plugin as specified.
