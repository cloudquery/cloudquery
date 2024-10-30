---
title: Troubleshooting
description: Troubleshooting common issues with CloudQuery.
---

# Troubleshooting

## Help Channels

### CloudQuery Community

First things first - feel free to join our [Community](https://community.cloudquery.io)!

### GitHub Issues

There are a couple of ways to get help for any CloudQuery-related issues or questions.

1. Check out previous issues at [https://github.com/cloudquery/cloudquery](https://github.com/cloudquery/cloudquery) and open a new one if no previous one has been opened or resolved.
2. Check out previous threads on our [Community](https://community.cloudquery.io) and open a new one if no previous one matches your issue. 

## Debugging

### Verbose Logging

Usually the first step that will be needed to debug/resolve an issue is to run `cloudquery` with `--log-level debug` to enable verbose logging.

:::callout{type="warning"}
The CLI may print out environment variables, including sensitive information, when verbose logging is enabled.
:::

### Error: "failed to migrate source"…

If you see an error such as `failed to migrate source`, it means that, while upgrading an integrations, the migration of the SQL schema failed.
CloudQuery makes a best-effort attempt to automatically and transparently manage the schemas of integrations, but this can sometimes fail during version upgrades.

The easiest solution is to drop and recreate the database or schema (or less destructively, all the integration's tables, such as `aws_*`).

### I am running `cloudquery sync` with multiple source integrations (or multiple `cloudquery sync`s), but some tables are empty / some rows are duplicated

When running `cloudquery sync` with multiple source integrations (or multiple `cloudquery sync`s in parallel),
it is important that every integrations-configuration has a unique `name`. If the names are not unique,
the different integrations may overwrite/delete each others data.

It is also important that every integrations-configuration is fetching different data (i.e. no two integrations are fetching the same account/table/region combination).

You can read more about this [here](/docs/advanced-topics/running-cloudquery-in-parallel).

### My AWS sync is taking a long time. What can I do to speed it up?

A few specific tables in AWS are quite slow to sync. You can try skipping them if you don't need this data.
Take a look at the [skip_tables list](https://hub.cloudquery.io/plugins/source/cloudquery/aws).
If syncs are still taking a long time, you can also take a look at our [performance tuning guide](/docs/advanced-topics/performance-tuning).

### I am running `cloudquery sync` locally, but it is taking a long time / doesn't seem to finish

(If running the AWS integrations, try the `skip_tables` solution from the previous paragraph first).

Large CloudQuery integrations, such as AWS, make a lot of DNS queries (for example, the AWS integrations makes a DNS query per `num_regions * num_services`).
Some less performant DNS servers may not be able to handle this load (e.g. home routers).
You can try pointing your machine to a different DNS server (such as the reliable [Google DNS Server](https://developers.google.com/speed/public-dns)).