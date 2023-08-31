---
title: Source Spec Reference
description: Reference for the source spec CloudQuery configuration object.
---

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

(`string`, optional, default: `github`, available: `github`, `local`, `grpc`, `docker`)

- `github`: CloudQuery will look for and download the plugin from GitHub, and then execute it.
- `local`: CloudQuery will execute the plugin from a local path.
- `grpc`: mostly useful in debug mode when plugin is already running in a different terminal, CloudQuery will connect to the gRPC plugin server directly without spawning the process.
- `docker`: CloudQuery will run the plugin in a Docker container. This is most useful for plugins written in Python, as they do not support the `local` and `github` registries.

### path

(`string`, required)

Configures how to retrieve the plugin. The contents depend on the value of `registry` (`github` by default).

- For plugins hosted on GitHub, `path` should be of the form `"<org>/<repository>"`. For official plugins, should be `cloudquery/<plugin-name>`.
- For plugins that are located in the local filesystem, `path` should a filesystem path to the plugin binary.
- To connect to a running plugin via `grpc` (mostly useful for debugging), `path` should be the host-port of the plugin (e.g. `localhost:7777`).
- For plugins distributed via Docker, `path` should be the name of the Docker image (optionally including a tag, the same as you would use for `docker run`, e.g. `ghcr.io/cloudquery/cq-source-typeform:v1.0.0`).

### version

(`string`, required)

`version` must be a valid [SemVer](https://semver.org/)), e.g. `vMajor.Minor.Patch`. You can find all official plugin versions under [cloudquery/cloudquery/releases](https://github.com/cloudquery/cloudquery/releases), and for community plugins you can find it in the relevant community repository. This is only relevant for plugins hosted on GitHub.

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

<!-- vale off -->

### deterministic_cq_id

<!-- vale on -->

(`bool`, optional, default: `false`, introduced in CLI `v2.4.1`)

A flag that indicates whether the value of `_cq_id` should be a UUID that is a hash of the primary keys or a random UUID. If a resource has no primary keys defined the value will always be a random UUID

Supported by source plugins released on 2023-03-08 and later

<!-- vale off -->

### otel_endpoint (preview)

<!-- vale on -->

(`string`, optional, introduced in CLI `v3.10.0`)

Open Telemetry [OTLP/HTTP](https://opentelemetry.io/docs/specs/otel/protocol/exporter/) exporter. Also, supports Jaeger endpoint. This will send traces of syncs to that endpoint.

<!-- vale off -->

### otel_endpoint_insecure (preview)

<!-- vale on -->

(`bool`, optional, default: `false`, introduced in CLI `v3.10.0`)

If set to `true`, the exporter will not verify the server will connect via `http` instead of `https`.

### spec

(`object`, optional)

Plugin-specific configurations. Visit [source plugins](/docs/plugins/sources/overview) documentation for more information.

## Top level deprecated options

### concurrency

This option was deprecated in CLI `v3.6.0` in favor of plugin level concurrency, as each plugin as its own concurrency requirements. See more in each plugin documentation.

### scheduler

This option was deprecated in CLI `v3.6.0` in favor of plugin level scheduler, as each plugin as its own scheduler requirements. See more in each plugin documentation.

### backend

This option was deprecated in CLI `v3.6.0` in favor of `backend_options`. See [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables) for more information.

### backend_spec

This option was deprecated in CLI `v3.6.0` in favor of `backend_options`. See [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables) for more information.
