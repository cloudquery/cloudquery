# Source Spec

This goes through all options available for sources plugins `spec` object:

## Spec

`name` (`string`, `required`) - Name of the plugin. If you have multiple source plugins this should be unique.

`path` (`string`, `optional`) - By default `path` is automatically inferred from `name`. For example for `postgresql`, `path` will be `cloudquery/postgresql`, or for `community_org/source_name` `path` will be `community_org/source_name`.

`version` (`string`, `required`) - Should be in the form of `vX.Y.Z` (or `vX.Y.Z-sometag` as per [semver](https://semver.org/)). Latest versions can be generate by `cloudquery gen` command. Also you can find all official plugins versions under [cloudquery/cloudquery/releases](https://github.com/cloudquery/cloudquery/releases) and for community you find in the relevant community repository.

`registry` (`string`, `optional`, default: `github`, available: `github,local,grpc`) - `github`: cloudquery will look and download the plugin from github and then execute it. `local`: cloudquery will execute the plugin from a local apth. `grpc`: mostly useful in debug mode when plugin is already running in a different terminal, cloudquery will connect to the gRPC plugin server directly without spawning the process.

`tables` (`[]string`, `optional`, default: `["*"]`): Tables to sync from the source plugin.

`skip_tables` (`[]string`, `optional`, default: `[]`) : Useful when using glob in `tables`, specify which tables to skip when syncing the source plugin.

`destination` (`[]string`, `required`): Specify the names of the destination to sync the data of the source plugin to.

## Development

Spec is defined in [plugin-sdk](https://github.com/cloudquery/plugin-sdk/blob/main/specs/source.go#L11)
