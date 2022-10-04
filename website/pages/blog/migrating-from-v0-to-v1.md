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

## Changes to the CLI Commands

Users of CloudQuery v0 would be familiar with the main commands `init` and `fetch`, as well as some others. These are no longer in V1, but similar commands have replaced them to reflect that CloudQuery now supports both source and destination plugins.

### Generate commands

- Use `cloudquery gen source <plugin-name>` to generate the config for a source plugin (e.g. `cloudquery gen source aws`)
- Use `cloudquery gen destination <plugin-name>` to generate the config for a destination plugin (e.g. `cloudquery gen destination postgresql`)

`cloudquery init` has been replaced by `cloudquery gen` and is no longer part of the CLI. See `cloudquery gen --help` for more information, or check our [online reference](/docs/reference/cli/cloudquery_generate).

### Sync commands

`cloudquery sync` replaces the v0 `cloudquery fetch` command.

Functionally it is still the same: it loads data from a source to a destination, but `sync` now supports multiple destinations, while `fetch` only supported PostgreSQL. With this change also comes a change in expected config format, see the [next section](#changes-to-the-configuration-format) for more details on this.

`cloudquery sync` needs to be passed a path to a config file or directory containing config files. So for example, to sync using a single yaml file:

```
cloudquery sync config.yml
```

or to sync using a directory of files (all `.yml` files in the directory will be used):

```
cloudquery sync config/
```

See `cloudquery sync --help` for more details, or check our [online reference](/docs/reference/cli/cloudquery_sync).

## Changes to the Configuration Format

V1 introduces a new config format that is closely related to the old one, but an old config will need some massaging to work with the CloudQuery v1 CLI. 

### Source Plugins

The new config format for source plugins are as follows:

```yaml
kind: source
spec:
  ## Required. name of the plugin to use
  name: "aws" # required
 
  # Required. Must be a specific version starting with v, e.g. v1.2.3
  version: "vX.Y.Z"
 
  ## Optional. Default: "github". Available: "local", "grpc"
  # registry: github
 
  ## Optional. Default: cloudquery/name
  # path: cloudquery/aws
 
  ## Optional. Default: ["*"] - all tables. We recommend to specify specific tables that you need to sync as this
  ## will reduce the amount of data synced and improve performance.
  # tables: ["*"]
 
  ## Required. all destinations you want to sync data to.
  destinations: ["postgresql"]
 
  spec:
    # plugin specific configuration.
```

Check the documentation for each plugin for details on how to configure the plugin-specific spec part. However, generally these will be exactly the same as in v0, and all the same authentication functionality is still supported.

### Destination Plugins

The new config format for destination plugins (e.g. PostgreSQL) is as follows:

```yaml
kind: destination
spec:
  ## Required. name of the plugin
  name: "postgresql"
 
  # Required. Must be a specific version starting with v, e.g. v1.2.3
  version: "vX.Y.Z"
 
  ## Optional. Default: "overwrite". Available: "overwrite", "append", "overwrite-delete-stale". Not all modes are 
  ## supported by all plugins, so make sure to check the plugin documentation for more details.
  write_mode: "overwrite" # overwrite, overwrite-delete-stale, append
 
  spec:
    ## plugin-specific configuration for PostgreSQL:
 
    ## Required. Connection string to your PostgreSQL instance
    connection_string: "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"```
```

### Files and Directories

The `sync` command supports loading config from files or directories, and you may choose to combine multiple source- and destination- configs in a single file using `---` to separate different sections.

## Changes to Tables and Schemas

Finally, during our work for v1, we endeavoured to make the table schemas more consistent, predictable and aligned with their upstream APIs. As such, some breaking changes to the schema were necessary. Each source plugin has its own schema migration guide to help you make the necessary changes to your custom queries, triggers and policies:

- [AWS](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/docs/v1-migration.md)
- [Azure](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/azure/docs/v1-migration.md)
- [CloudFlare](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/cloudflare/docs/v1-migration.md)
- [DigitalOcean](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/digitalocean/docs/v1-migration.md)
- [GCP](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/gcp/docs/v1-migration.md)
- [Github](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/github/docs/v1-migration.md)
- [Heroku](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/heroku/docs/v1-migration.md)
- [K8s](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/k8s/docs/v1-migration.md)
- [Okta](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/okta/docs/v1-migration.md)
- [Terraform](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/terraform/docs/v1-migration.md)

Note that these guides are (for the most part) automatically generated, so in some cases a table may be marked as removed when it was actually renamed. Please reach out to us if you find any errors.

### Get Help / Ask Questions

If you have any questions about migrating or CloudQuery v1, don't hesitate to reach out on [Discord](https://www.cloudquery.io/discord). We're a friendly community and would love to help however we can.  