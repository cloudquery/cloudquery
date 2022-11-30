---
title: Migrating from CloudQuery v0 to v1
tag: announcement
date: 2022/10/03
description: >-
  A guide for users migrating from CloudQuery v0 to v1
author: hermanschaaf
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

We are thrilled to announce the release of the first major version of CloudQuery--see our [v1 announcement blog post](/blog/cloudquery-v1-release) for details! With the new release comes a range of new exciting features, and this page is here to help you migrate an existing CloudQuery installation from v0 to v1.

## Changes in V1

The [announcement blog post](/blog/cloudquery-v1-release) lists many of the important improvements, and we won't re-iterate them all here. Most changes are internal and developer-facing, but some do impact existing CloudQuery teams. Those are:

- [changes to CLI commands](#changes-to-the-cli-commands)
- [changes to the configuration format](#changes-to-the-configuration-format)
- [changes to tables and schemas](#changes-to-tables-and-schemas)

## Changes to the Configuration Format

V1 introduces a new config format that is closely related to the old one, but an old config will need some massaging to work with the CloudQuery v1 CLI. Mostly because we now support multiple destinations, there are separate configs for source and destination plugins.

### Source Plugins

The new config format for source plugins are as follows:

```yaml
kind: source
spec:
  ## Required. name of the plugin to use
  name: "aws" # required
 
  # Required. Must be a specific version starting with v, e.g. v1.2.3
  version: "VERSION_SOURCE_AWS"
 
  ## Optional. Default: "github". Available: "local", "grpc"
  # registry: github
 
  ## Plugin path. For official plugins, this should be in the format "cloudquery/<name>", e.g. "cloudquery/aws"
  path: "cloudquery/aws"
 
  ## Optional. Default: ["*"] - all tables. We recommend to specify specific tables that you need to sync as this
  ## will reduce the amount of data synced and improve performance.
  # tables: ["*"]
 
  ## Required. all destinations you want to sync data to.
  destinations: ["postgresql"]
 
  spec:
    # plugin specific configuration.
```

Check the [source spec documentation](/docs/reference/source-spec) for general layout, and individual [plugin documentation](/docs/plugins/sources/overview) for details on how to configure the plugin-specific spec. Generally these will be the same as in v0, and all the same authentication functionality is still supported.

### Destination Plugins

The new config format for destination plugins (e.g. PostgreSQL) is as follows:

```yaml
kind: destination
spec:
  ## Required. name of the plugin
  name: "postgresql"
 
  path: "cloudquery/postgresql"

  # Required. Must be a specific version starting with v, e.g. v1.2.3
  version: "VERSION_DESTINATION_POSTGRESQL"
 
  ## Optional. Default: "overwrite". Available: "overwrite", "append", "overwrite-delete-stale". Not all modes are 
  ## supported by all plugins, so make sure to check the plugin documentation for more details.
  write_mode: "overwrite" # overwrite, overwrite-delete-stale, append
 
  spec:
    ## plugin-specific configuration for PostgreSQL:
 
    ## Required. Connection string to your PostgreSQL instance
    connection_string: "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"```
```

Check the [destination spec documentation](/docs/reference/destination-spec) for general layout, and individual [destination plugin documentation](/docs/plugins/destinations/overview) for details on how to configure the plugin-specific spec part. Generally these will be the same as in v0, and all the same authentication functionality is still supported.

## Changes to the CLI Commands

Users of CloudQuery v0 would be familiar with the main commands `init` and `fetch`. These have changed in v1 and `init` is longer available (you should write config files manually).

### Init

`init` was a command that generated a starter configuration template, but it is no longer a command in v1 of the CLI. Instead, please refer to our [Quickstart](/docs/quickstart) guide to see how source and destination plugins should be configured.

The previous `init` command also generated a full list of tables to fetch. In v1, you can fetch all tables by using a wildcard entry:

```
tables: ["*"]
```

in the source configuration file. This can also be combined with the `skip_tables` option to fetch all tables except some subset:

```
tables: ["*"]
skip_tables: ["aws_accessanalyzer_analyzers", "aws_acm_certificates"]
```

### Sync

`cloudquery sync` replaces the v0 `cloudquery fetch` command.

Functionally it is still the same: it loads data from a source to a destination, but `sync` now supports multiple destinations, while `fetch` only supported PostgreSQL. With this change also comes a change in expected config format, see the [next section](#changes-to-the-configuration-format) for more details on this.

`cloudquery sync` needs to be passed a path to a config file or directory containing config files. So for example, to sync using all `.yml` files in a directory named `config`:

```bash
cloudquery sync config/
```

Or to sync using a single YAML file named `config.yml`:

```bash
cloudquery sync config.yml
```

In this case `config.yml` should contain at least one source and one destination config, each separated by a line containing three dashes (`---`). More about this in [Files and Directories](#files-and-directories).

See `cloudquery sync --help` for more details, or check our [online reference](/docs/reference/cli/cloudquery_sync).

### Files and Directories

The `sync` command supports loading config from files or directories, and you may choose to combine multiple source- and destination- configs in a single file using `---` on its own line to separate different sections. For example:

```
kind: source
spec:
    name: "aws"
    version: "VERSION_SOURCE_AWS"
    # rest of source spec here
---
kind: destination
spec:
    name: "postgresql"
    version: "VERSION_DESTINATION_POSTGRESQL"
    # rest of destination spec here
```

## Changes to Tables and Schemas

Finally, during our work for v1, we endeavoured to make the table schemas more consistent, predictable and aligned with their upstream APIs. As such, some breaking changes to the schema were necessary. Each source plugin has its own schema migration guide to help you make the necessary changes to your custom queries, triggers and policies:

- [AWS](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/docs/v1-migration.md)
- [Azure](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/azure/docs/v1-migration.md)
- [CloudFlare](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/cloudflare/docs/v1-migration.md)
- [DigitalOcean](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/digitalocean/docs/v1-migration.md)
- [GCP](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/gcp/docs/v1-migration.md)
- [GitHub](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/github/docs/v1-migration.md)
- [Heroku](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/heroku/docs/v1-migration.md)
- [K8s](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/k8s/docs/v1-migration.md)
- [Okta](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/okta/docs/v1-migration.md)
- [Terraform](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/terraform/docs/v1-migration.md)

Note that these guides are (for the most part) automatically generated, so in some cases a table may be marked as removed when it was actually renamed. Please reach out to us if you find any errors.

## Start from a clean Database

V1 introduces functionality to automatically perform backwards-compatible Postgres migrations when new columns or tables are added. However, this functionality relies on a clean start being made in V1, and if you try to run it against a database with tables from v0, there is a good chance it will fail.

Therefore, it is important that you **start from a clean database**. This can either mean creating a new database and pointing the v1 configuration there, or dropping all the tables in your v0 database. 

## Get Help / Ask Questions

If you run into issues not covered here, or have any questions about migrating or CloudQuery v1, don't hesitate to reach out on [Discord](https://www.cloudquery.io/discord). We're a friendly community and would love to help however we can.
