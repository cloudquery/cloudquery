# Destination Spec Reference

This goes through all the available options for the destination plugin `spec` object.

## Example

This example configures the `postgresql` plugin to connect to a PostgreSQL database located at `localhost:5432`.

```yaml copy
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  version: "VERSION_DESTINATION_POSTGRESQL"

  spec:
    connection_string: "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
```

## Spec

### name

(`string`, required)

Name of the plugin. If you have multiple destination plugins, this must be unique.

The name field may be used to uniquely identify a particular destination configuration. For example, if you have two configs for the PostgreSQL plugin for syncing different databases, one may be named `db-1` and the other `db-2`. In this case, the `path` option below must be used to specify the download path for the plugin.

### registry

(`string`, optional, default: `github`, available: `github`, `local`, `grpc`)

- `github`: CloudQuery will look for and download the plugin from GitHub, and then execute it.
- `local`: CloudQuery will execute the plugin from a local path. 
- `grpc`: mostly useful in debug mode when plugin is already running in a different terminal, CloudQuery will connect to the gRPC plugin server directly without spawning the process.

### path

(`string`, required)

Configures how to retrieve the plugin. The contents depend on the value of `registry` (`github` by default).
- For plugins hosted on GitHub, `path` should be of the form `"<org>/<repository>"`. For official plugins, should be `cloudquery/<plugin-name>`.
- For plugins that are located in the local filesystem, `path` should a filesystem path to the plugin binary.
- To connect to a running plugin via `grpc` (mostly useful for debugging), `path` should be the host-port of the plugin (e.g. `localhost:7777`).

### version

(`string`, required)

`version` must be a valid [SemVer](https://semver.org/)), e.g. `vMajor.Minor.Patch`. You can find all official plugin versions under [our GitHub releases page](https://github.com/cloudquery/cloudquery/releases), and for community plugins you can find it in the relevant community repository.

### write_mode

(`string`, optional, default: `overwrite-delete-stale`. Available: `overwrite-delete-stale`, `overwrite`, `append`)

Specifies the update method to use when inserting rows. The exact semantics depend on the destination plugin, and all destinations don't support all options, so check the destination plugin documentation for details.

- `overwrite-delete-stale`: `sync`s overwrite existing rows with the same primary key, and delete rows that 
                             are no longer present in the cloud.
- `overwrite`: Same as `overwrite-delete-stale`, but doesn't delete stale rows from previous `sync`s.
- `append`: Rows are never overwritten or deleted, only appended.

### spec

(`object`, optional)

Plugin specific configurations. Visit [destination plugins](/docs/plugins/destinations/overview) documentation for more information.
