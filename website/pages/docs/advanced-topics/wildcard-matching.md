---
title: Wildcard Matching
---

# Wildcard Matching

The `tables` and `skip_tables` config options both support wildcard matching. This means that you can use `*` in the name to match multiple tables at once.

For example, it is possible to use a wildcard pattern to match all tables related to AWS EC2:

```yaml
tables:
 - aws_ec2_*
```

This can also be combined with `skip_tables`. For example, let's say we want to include all EC2 tables, but not EBS-related ones:

```yaml
tables: 
- "aws_ec2_*"
skip_tables:
- "aws_ec2_ebs_*"
```

## Skipping Relations

Some tables require a lot of API calls to sync. This can be especially true of tables that depend on other tables, because often multiple API calls need to be made for each individual resource in the parent table. This can easily lead to thousands of API calls, increasing the time it takes to sync. If you know that some child tables are not strictly necessary, you can boost the sync performance by skipping them with the `skip_tables` setting.

Let's say we have three tables: `A`, `B` and `C`. `A` is the top-level table. `B` depends on `A`, and `C` depends on `B`:

```text
A <- B <- C
```

We might want table `A`, but not need the information in table `B`. We can then write this:

```yaml
tables:
 - A
skip_tables:
 - B
```

By skipping table `B`, we are automatically skipping its dependant table `C` as well. Likewise, by including table `A`, we are automatically including its dependant tables `B` and `C` as well, unless they are explicitly skipped in the `skip_tables` section (like in the example above).