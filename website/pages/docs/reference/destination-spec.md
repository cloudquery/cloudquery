# Destination Spec Reference

This goes through all the available options for the destination plugin `spec` object.

## Spec

### name
(`string`, required)

Name of the plugin. If you have multiple destination plugins, this should be unique. 

The name field may be used to uniquely identify a particular destination configuration. For example, if you have two configs for the PostgreSQL plugin for syncing different databases, one may be named `db-1` and the other `db-2`. In this case, the `path` option below should be used to specify the download path for the plugin.

### path
(`string`, optional)

Configures how to retrieve the plugin. For plugins hosted on GitHub, `path` is inferred from `name` by default.
For example `name: postgresql` will resolve `path` to `clouduquery/postgresql`. If you'd like to use a plugin that is not hosted on the CloudQuery repository, provide the full path to the repository such as `community-github-org/community-github-repo`

### version
(`string`, required)

`version` must be a valid [SemVer](https://semver.org/)), e.g. `vMajor.Minor.Patch`. Latest versions can be generated using `cloudquery gen` command. You can find all official plugin versions under [cloudquery/cloudquery/releases](https://github.com/cloudquery/cloudquery/releases), and for community plugins you can find it in the relevant community repository.

### registry
(`string`, optional, default: `github`, available: `github`, `local`, `grpc`)
 
 - `github`: CloudQuery will look for and download the plugin from GitHub, and then execute it. 
 - `local`: CloudQuery will execute the plugin from a local path. 
 - `grpc`: mostly useful in debug mode when plugin is already running in a different terminal, CloudQuery will connect to the gRPC plugin server directly without spawning the process.

### write_mode
(`string`, optional, default: `overwrite`. Available: `overwrite`, `overwrite-delete-stale`, `append`)

Specifies the update method to use when inserting rows. The exact semantics depend on the destination plugin, and all destinations don't support all options, so check the destination plugin documentation for details.

- `overwrite`: Overwrites existing rows with the same primary key, but doesn't delete stale rows from previous syncs.
- `overwrite-delete-stale`: Same as `overwrite`, but stale rows are also deleted.
- `append`: Rows are never overwritten or deleted, only appended.

## Development

Spec is defined in [plugin-sdk](https://github.com/cloudquery/plugin-sdk/blob/main/specs/destination.go#L12)
