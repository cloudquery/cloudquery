# Destination Spec

This goes through all options available for destination plugins `spec` object:

## Spec

`name` (`string`, `required`) - Name of the plugin. If you have multiple destination plugins this should be unique.

`path` (`string`, `optional`) - By default `path` is automatically inferred from `name`. For example for `postgresql`, `path` will be `cloudquery/postgresql`, or for `community_org/dest_name` `path` will be `community_org/dest_name`.

`version` (`string`, `required`) - Should be in the form of `vX.Y.Z` (or `vX.Y.Z-sometag` as per [semver](https://semver.org/)). Latest versions can be generate by `cloudquery gen` command. Also you can find all official plugins versions under [cloudquery/cloudquery/releases](https://github.com/cloudquery/cloudquery/releases) and for community you find in the relevant community repository.

`registry` (`string`, `optional`, default: `github`, available: `github,local,grpc`) - `github`: cloudquery will look and download the plugin from github and then execute it. `local`: cloudquery will execute the plugin from a local apth. `grpc`: mostly useful in debug mode when plugin is already running in a different terminal, cloudquery will connect to the gRPC plugin server directly without spawning the process.

## Development

Spec is defined in [plugin-sdk](https://github.com/cloudquery/plugin-sdk/blob/main/specs/destination.go#L12)
