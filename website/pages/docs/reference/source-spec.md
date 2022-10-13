# Source Spec Reference

Following are available options for the top level source plugin `spec` object. 

Note: For configuring individual plugins, please refer to the configuration section from the relevant plugins from [here](https://www.cloudquery.io/docs/plugins/sources). (eg. [AWS plugin configuration](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/docs/configuration.md)). Also check out our [recipes](https://www.cloudquery.io/docs/recipes/aws-postgresql) for common configuration examples.

## Spec

### name

(`string`, required)

Name of the plugin. If you have multiple source plugins, this should be unique. 

The name field may be used to uniquely identify a particular source configuration. For example, if you have two configs for the AWS plugin for syncing different accounts, one may be named `aws-account-1` and the other `aws-account-2`. In this case, the `path` option below should be used to specify the download path for the plugin.

### path

(`string`, optional)

Configures how to retrieve the plugin. For plugins hosted on GitHub, `path` is inferred from `name` by default.
For example `name: aws` will resolve `path` to `clouduquery/aws`. If you'd like to use a plugin that is not hosted on the CloudQuery repository, provide the full path to the repository such as `community-github-org/community-github-repo`

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

### table_concurrency

(`int`, optional, default: `500000`, introduced in CLI `v1.0.8`)

Sets a global limit on the number of tables to sync concurrently.

### resource_concurrency

(`int`, optional, default: `500000`, introduced in CLI `v1.0.8`)

Sets an (approximate) global limit on the number of concurrent requests performed to fetch further details about resources. Note that this limit only applies to top-level tables and not their child relations.

