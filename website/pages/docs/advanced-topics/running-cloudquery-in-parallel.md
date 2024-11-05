---
title: Running CloudQuery in Parallel
description: Running multiple instances of `cloudquery sync` in parallel can be useful when a single sync is too slow, for example when syncing a large number of accounts, or when fetching from large accounts.
---

# Running CloudQuery in Parallel

Running multiple instances of `cloudquery sync` in parallel can be useful when a single sync is too slow, for example when syncing a large number of accounts, or when fetching from large accounts.

## Splitting Syncs Automatically

Starting from version [v6.8.0](https://github.com/cloudquery/cloudquery/releases/tag/cli-v6.8.0) of the CloudQuery CLI, you can use the `--shard` flag to automatically split a sync into smaller parts that can be run in parallel.

For example, to split a sync into 4 parts, you can run:

```bash
cloudquery sync config.yml --shard 1/4
cloudquery sync config.yml --shard 2/4
cloudquery sync config.yml --shard 3/4
cloudquery sync config.yml --shard 4/4
```

The `shard` flag will automatically split the sync into parts, ensure each part gets a unique source name, and that the parts don't overlap.
It's recommended to run the parts in parallel, as the sync will be faster than running a single sync.

You can find an example of how to run the syncs in parallel in the [GitHub Actions Deployment Guide](/docs/deployment/github-actions#running-cloudquery-in-parallel-to-speed-up-sync-time) section.

### Supported Source Integrations for Sharding

| Source Integration | Minimal Version                                                                   |
| ------------- | --------------------------------------------------------------------------------- |
| AWS           | [v27.20.0](https://hub.cloudquery.io/plugins/source/cloudquery/aws/latest/docs) |
| Azure         | [v14.8.0](https://hub.cloudquery.io/plugins/source/cloudquery/azure/latest/docs)  |
| GCP           | [v16.3.0](https://hub.cloudquery.io/plugins/source/cloudquery/gcp/latest/docs)  |

## Splitting Syncs Manually

If you are using an older version of the CloudQuery CLI, or if you want to manually split a sync, you can do so by creating different configurations for each part of the sync, using the guidelines below.

### Unique Names

Every source and destination integration configuration must have a unique `name`. This is required because the `name` is
written into the database (`_cq_source_name`), and is used to later delete stale resources.

For instance, a configuration with multiple source integrations could look like:

```yaml copy
kind: source
spec:
  name: aws1
  path: cloudquery/aws
  registry: cloudquery
  ...
---
kind: source
spec:
  name: aws2
  path: cloudquery/aws
  registry: cloudquery
  ...
---
kind: destination
spec:
  name: "postgresql"
  path: cloudquery/postgresql
  registry: cloudquery
  ...
```

If the names are not unique, then the different integrations may delete/overwrite each other's resources.

### No Overlapping Syncs

When splitting a sync into multiple source-integration configurations to be run in parallel, it is important
that these syncs don't overlap - the set of Account/Table/Region that every source-integration grabs must not intersect.

For instance, in GCP, if the first source-integration fetches resource `A` from project `1`, the second source-integration
can fetch resource `B` from project `1`, or resource `A` from project `2`, but can never fetch resource `A` from project `1`.

For another example, if the first source-integration fetches from region `europe-west1` in project `1`, the second source-integration
can fetch from region `europe-west1` in project `2`, or from region `europe-west2` in project `1`, but can never fetch from
region `europe-west1` in project `1`.

If the configurations overlap, the behavior is undefined, and the database may contain duplicate rows.
