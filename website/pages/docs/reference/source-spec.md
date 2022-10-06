# Source Spec Reference

This goes through all the available options for the source plugin `spec` object:

## Spec

### `name`

(`string`, required)

Name of the plugin. If you have multiple source plugins, this should be unique. 

The name field may be used to uniquely identify a particular source configuration. For example, if you have two configs for the AWS plugin for syncing different accounts, one may be named `aws-account-1` and the other `aws-account-2`. In this case, the `path` option below should be used to specify the download path for the plugin.

### path

(`string`, optional)

Configures how to retrieve the plugin. For plugins hosted on GitHub, `path` is inferred from `name` by default.
For example `name: aws` will resolve `path` to `clouduquery/aws`. If you'd like to use a plugin that is not hosted on the CloudQuery repository, provide the full path to the repository such as `community-github-org/community-github-repo`.  
If plugin registry is set to `grpc`, path should be an address that plugin is listening on. For example if you started a plugin locally in debug mode it will be listening on `localhost:50051` and you can use that as the path.
If plugin registry is `local`, path should be a path to the plugin binary. For example if you have a plugin binary in `~/.cloudquery/plugins` you can use `~/.cloudquery/plugins/<bin_name>` as the path.

### version

(`string`, required)

`version` must be a valid [SemVer](https://semver.org/)), e.g. `vMajor.Minor.Patch`. Latest versions can be generated using `cloudquery gen` command. You can find all official plugin versions under [cloudquery/cloudquery/releases](https://github.com/cloudquery/cloudquery/releases), and for community plugins you can find it in the relevant community repository.

### registry

(`string`, optional, default: `github`, available: `github`, `local`, `grpc`)

- `github`: CloudQuery will look for and download the plugin from GitHub, and then execute it.
- `local`: CloudQuery will execute the plugin from a local path. 
- `grpc`: mostly useful in debug mode when plugin is already running in a different terminal, CloudQuery will connect to the gRPC plugin server directly without spawning the process.

### tables

(`[]string`, optional, default: `["*"]`)

Tables to sync from the source plugin.

### skip_tables

(`[]string`, optional, default: `[]`)

Useful when using glob in `tables`, specify which tables to skip when syncing the source plugin.

### destinations

(`[]string`, required)

Specify the names of the destinations to sync the data of the source plugin to.

## Development

Spec is defined in [plugin-sdk](https://github.com/cloudquery/plugin-sdk/blob/main/specs/source.go#L11)
