---
title: Rate Limiting
---

# Rate Limiting

There are two main levers to control the rate at which CloudQuery fetches resources from cloud providers. These are the `table_concurrency` and `resource_concurrency` options that can be specified as [part of the source spec](/docs/reference/source-spec). Note that these options were introduced in CloudQuery CLI v1.0.8.

## Table Concurrency

`table_concurrency` controls the number of concurrent tables that will be processed while performing a sync. Setting this to a low number will reduce the number of concurrent requests, making it less likely to hit rate limits. The trade-off is that syncs will take longer to complete.

## Resource Concurrency

`resource_concurrency` is an approximate global limit on how many concurrent requests will be made to fetch details about the initial rows returned by a table's resolver. This limit applies only to top-level tables, and child relations will not be limited. Setting this to a lower number will also reduce the number of concurrent requests made, regardless of how many tables are being synced at any one time. As with `table_concurrency`, the trade-off is that syncs will take longer to complete.