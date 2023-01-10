---
title: Rate Limiting
---

# Rate Limiting

There is currently one main lever to control the rate at which CloudQuery fetches resources from cloud providers. This setting is called `concurrency`, and it can be specified as [part of the source spec](/docs/reference/source-spec). Note that this option was introduced in CloudQuery CLI v1.4.1.

## Concurrency

`concurrency` provides rough control over the number of concurrent requests that will be made while performing a sync. Setting this to a low number will reduce the number of concurrent requests, reducing the memory used and making the sync less likely to hit rate limits. The trade-off is that syncs will take longer to complete.
