# Development Environment Setup

## Requirements
 * [Go](https://go.dev/doc/install) 1.19+ (to build the plugins)

## Quick Start

### Building

Clone the repository:

```bash
git clone https://github.com/cloudquery/cloudquery
```

Build the CLI and all plugins:

```
make build
```

### Running Source Plugins in Developer Mode

1. Execute `go run main.go serve` from the chosen plugin directory under [../plugins/source](../plugins/source) (e.g.  [../plugins/source/aws](../plugins/source/aws)).
2. Create a config file for the source plugin. See the plugin's README.md for details. In the global spec section, set `registry` to `grpc` and `path` to `localhost:7777`. For example:
   ```yaml
   kind: "source"
   spec:
     # global config
     name: "aws"
     version: "latest"
     registry: "grpc"
     path: "localhost:7777"
     tables: ["*"]
     destinations: ["postgresql"]
     spec:
     # plugin-specific config
   ```
3. Create a config file for the destination plugin to load data into. See the [Destination Plugin](../plugins/destination)'s README.md for examples.
4. Open another terminal and run `bin/cloudquery sync <config-dir>`, where `<config-dir>` is the directory containing the config files.

Note that plugin logs will be output to the plugin process terminal.

### Testing

To run tests all unit tests for a plugin, inside the plugin directory run:

```shell
make test-unit  # This runs go test ./...
```

Unit Tests don't require any credentials, but some may require internet access.

Unit tests for plugins include:
- Specific resource tests. You can find those next to each resource, in the `resources/services` folder under the plugin directory.
- Client tests. You can find those in the `client` folder.
