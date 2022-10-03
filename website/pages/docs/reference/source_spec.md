# Source Spec

This goes through all the options available for source plugins' `spec` object:

## Spec

### `name` 
(`string`, `required`)

Name of the plugin. If you have multiple source plugins, this should be unique.

### `path` 
(`string`, `optional`)
 
By default, `path` is automatically inferred from `name`. For example, for `aws`, `path` will be `cloudquery/aws`, or for `community_org/source_name` `path` will be `community_org/source_name`.


###  `version` 
(`string`, `required`) 

`version` should be in the form of `vX.Y.Z` (or `vX.Y.Z-sometag` as per [semver](https://semver.org/)). Latest versions can be generated using `cloudquery gen` command. You can find all official plugin versions under [cloudquery/cloudquery/releases](https://github.com/cloudquery/cloudquery/releases), and for community plugins you can find it in the relevant community repository.

### `registry` 
(`string`, `optional`, default: `github`, available: `github`, `local`, `grpc`)
 
 - `github`: CloudQuery will look for and download the plugin from GitHub, and then execute it. 
 - `local`: CloudQuery will execute the plugin from a local path. 
 - `grpc`: mostly useful in debug mode when plugin is already running in a different terminal, CloudQuery will connect to the gRPC plugin server directly without spawning the process.

### `tables`
(`[]string`, `optional`, default: `["*"]`)

Tables to sync from the source plugin.

### `skip_tables` 
(`[]string`, `optional`, default: `[]`)

Useful when using glob in `tables`, specify which tables to skip when syncing the source plugin.

### `destinations`
(`[]string`, `required`)

Specify the names of the destinations to sync the data of the source plugin to.

## Development

Spec is defined in [plugin-sdk](https://github.com/cloudquery/plugin-sdk/blob/main/specs/source.go#L11)
