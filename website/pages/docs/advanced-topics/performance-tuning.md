---
title: Performance Tuning
---

# Performance Tuning

This page contains a number of tips and tricks for improving the performance of `cloudquery sync` for large cloud estates.

## Tune Concurrency

The `concurrency` setting, available for all source plugins as part of the [source spec](/docs/reference/source-spec#concurrency), controls the approximate number of concurrent requests that will be made while performing a sync. Setting this to a low number will reduce the number of concurrent requests, reducing the memory used and making the sync less likely to hit rate limits. The trade-off is that syncs will take longer to complete.

## Use a Different Scheduler

By default, CloudQuery syncs will fetch all tables in parallel, writing data to the destination(s) as they come in. However, the `concurrency` setting, mentioned above, places a limit on how many **table-clients** can be synced at a time. What "table-client" means depends on the source plugin and the table. In AWS, for example, a client is usually a combination of account and region. Get all the combinations of accounts and regions for all tables, and you have all the table-clients for a sync. For the GCP source plugin, clients generally map to projects.

The default CloudQuery scheduler, known as `dfs`, will sync up to `concurrency / 100` table-clients at a time (we are ignoring child relations for the purposes of this discussion). Let's take an example GCP cloud estate with 5000 projects, syncing 100 tables. This makes for approximately 500,000 table-client pairs, and a concurrency of 10,000 will allow 100 table-client pairs to be synced at a time. The `dfs` scheduler will start with the first table and its first 100 projects, and then move on to finish all projects for that table before moving on to the next table. This means, in practice, only one table is really being synced at a time!

Usually this works out fine, as long as the cloud platform's rate limits are aligned with the clients. But if rate limits are applied per-table, rather than per-project, `dfs` can be suboptimal. A better strategy in this case would be to choose the first client for every table before moving on to the next client. This is what the `round-robin` scheduler does. Note that this scheduler is currently **experimental**. 

The following example config enables `round-robin` scheduling for the GCP source plugin, with a concurrency of 100000:

```yaml
kind: source
spec:
  name: "gcp"
  path: "cloudquery/gcp"
  version: "VERSION_SOURCE_GCP"
  destinations: ["postgresql"]
  concurrency: 100000
  scheduler: "round-robin"  # experimental
```

## Use Wildcard Matching

import { Callout } from 'nextra-theme-docs'

Sometimes the easiest way to improve the performance of the `sync` command is to limit the number of tables that get synced. The `tables` and `skip_tables` source config options both support wildcard matching. This means that you can use `*` anywhere in a name to match multiple tables.

For example, when using the `aws` source plugin, it is possible to use a wildcard pattern to match all tables related to AWS EC2:

```yaml copy
tables:
 - aws_ec2_*
```

This can also be combined with `skip_tables`. For example, let's say we want to include all EC2 tables, but not EBS-related ones:

```yaml copy
tables: 
- "aws_ec2_*"
skip_tables:
- "aws_ec2_ebs_*"
```

<Callout> 

The CloudQuery CLI will warn if a wildcard pattern does not match any known tables.

</Callout>

## Improve Performance by Skipping Relations

Some tables require many API calls to sync. This is especially true of tables that depend on other tables, because often multiple API calls need to be made for every row in the parent table. This can lead to thousands of API calls, increasing the time it takes to sync. If you know that some child tables are not strictly necessary, you can improve sync performance by skipping them with the `skip_tables` setting.

Let's say we have three tables: `A`, `B` and `C`. `A` is the top-level table. `B` depends on it, and `C` depends on `B`:

```text
A 
↳ B
  ↳ C
```

We might want table `A`, but not need the information in table `B`. We can then write our source config as:

```yaml copy
tables:
 - A
skip_tables:
 - B
```

By skipping table `B`, we are automatically skipping its dependant table `C` as well. Likewise, by including table `A`, we are automatically including its dependant tables `B` and `C` as well, unless they are explicitly skipped in the `skip_tables` section (like in the example above).
