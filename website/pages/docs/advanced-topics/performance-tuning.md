---
title: Performance Tuning
---

# Performance Tuning

This page contains a number of tips and tricks for improving the performance of `cloudquery sync` for large cloud estates.

## Wildcard Matching

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

## Improving Performance by Skipping Relations

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
