# Source Spec Reference

Following are available options for the top level source plugin `spec` object.

Note: For configuring individual plugins, please refer to the configuration section from the relevant plugins from [here](/docs/plugins/sources/overview). (e.g. [AWS plugin configuration](/docs/plugins/sources/aws/configuration)).

## Example

This example configures the AWS plugin, and connects is to a `postgresql` destination:

```yaml copy
kind: source
spec:
  name: "aws"
  path: "cloudquery/aws"
  version: "VERSION_SOURCE_AWS"
  tables: ["aws_s3_buckets"]
  destinations: ["postgresql"]

  spec:
```

## Spec

### name

(`string`, required)

Name of the plugin. If you have multiple source plugins, this must be unique. 

The name field may be used to uniquely identify a particular source configuration. For example, if you have two configs for the AWS plugin for syncing different accounts, one may be named `aws-account-1` and the other `aws-account-2`. In this case, the `path` option below must be used to specify the download path for the plugin.

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

`version` must be a valid [SemVer](https://semver.org/)), e.g. `vMajor.Minor.Patch`. You can find all official plugin versions under [cloudquery/cloudquery/releases](https://github.com/cloudquery/cloudquery/releases), and for community plugins you can find it in the relevant community repository.

### tables

(`[]string`, required)

> This option was changed to required in versions >= `v3.0.0` of the CloudQuery CLI. In previous versions it was optional and defaulted to `["*"]` (sync all tables).

Tables to sync from the source plugin. It accepts wildcards. For example, to match all tables use `["*"]` and to match all EC2-related tables use `aws_ec2_*`. Matched tables will also sync all their descendant tables, unless these are skipped in `skip_tables`. Please note that syncing all tables can be slow on some plugins (e.g. AWS, GCP, Azure).

### skip_tables

(`[]string`, optional, default: `[]`)

Specify which tables to skip when syncing the source plugin. It accepts wildcards. This config is useful when using wildcards in `tables`, or when you wish to skip dependent tables. Note that if a table with dependencies is skipped, all its dependant tables will also be skipped.

<!-- vale off -->
### skip_dependent_tables
<!-- vale on -->

(`bool`, optional, default: `false`, introduced in CLI `v2.3.7`)

If set to `true`, tables that depend on the tables specified in `tables` will not be synced, unless specifically selected themselves. This allows you to choose precisely which tables to sync, and prevents automatically syncing new dependent tables that may be added to the plugin in future versions. Note that if you specify a table that depends on another table, CloudQuery will still automatically include the parent table(s).

### destinations

(`[]string`, required)

Specify the names of the destinations to sync the data of the source plugin to.

### concurrency

(`int`, optional, default: `500000`, introduced in CLI `v1.4.1`)

A best effort maximum number of Go routines to use. Lower this number to reduce memory usage.

### scheduler

(`string`, optional, default: `dfs`, introduced in CLI `v2.0.31`, **EXPERIMENTAL**)

The scheduler to use when determining the priority of resources to sync. Currently, the only supported values are `dfs` (depth-first search) and `round-robin`. This is an experimental feature, and may be removed in the future. For more information about this, see [performance tuning](/docs/advanced-topics/performance-tuning).

### backend

(`string`, optional, default: `local`, introduced in CLI `v2.1.0`)

The backend to use for storing the state of incremental tables. Currently, the only supported value is `local` (store the state in the local filesystem). For more information, see [managing incremental tables](/docs/advanced-topics/managing-incremental-tables).

### backend_spec

(`object`, optional, introduced in CLI `v2.1.0`)

The backend spec is specific to the backend used. For the `local` backend, the only option is `path`, which specifies the name of the directory to use when storing metadata files.

<!-- vale off -->
### deterministic_cq_id
<!-- vale on -->

(`bool`, optional, default: `false`, introduced in CLI `v2.4.1`)

A flag that indicates whether the value of `_cq_id` should be a UUID that is a hash of the primary keys or a random UUID. If a resource has no primary keys defined the value will always be a random UUID

Supported by source plugins released on 2023-03-08 and later

### spec

(`object`, optional)

Plugin-specific configurations. Visit [source plugins](/docs/plugins/sources/overview) documentation for more information.
