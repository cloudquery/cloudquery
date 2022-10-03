# Destination Spec

This goes through all the options available for destination plugins' `spec` object:

## Spec

### `name`
(`string`, `required`)

Name of the plugin. If you have multiple destination plugins this should be unique.

### `path`
(`string`, `optional`)

By default, `path` is automatically inferred from `name`. For example for `postgresql`, `path` will be `cloudquery/postgresql`, or for `community_org/dest_name` `path` will be `community_org/dest_name`.

### `version`
(`string`, `required`)

`version` should be in the form of `vX.Y.Z` (or `vX.Y.Z-sometag` as per [semver](https://semver.org/)). Latest versions can be generated using `cloudquery gen` command. You can find all official plugin versions under [cloudquery/cloudquery/releases](https://github.com/cloudquery/cloudquery/releases), and for community plugins you can find it in the relevant community repository.

### `registry` 
(`string`, `optional`, default: `github`, available: `github`, `local`, `grpc`)
 
 - `github`: CloudQuery will look for and download the plugin from GitHub, and then execute it. 
 - `local`: CloudQuery will execute the plugin from a local path. 
 - `grpc`: mostly useful in debug mode when plugin is already running in a different terminal, CloudQuery will connect to the gRPC plugin server directly without spawning the process.

## Development

Spec is defined in [plugin-sdk](https://github.com/cloudquery/plugin-sdk/blob/main/specs/destination.go#L12)
