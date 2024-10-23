---
title: Rate Limiting
description: CloudQuery provides a few options to control the rate at which resources are fetched from cloud providers.
---

# Rate Limiting

There is currently one main lever to control the rate at which CloudQuery fetches resources from cloud providers. This setting is called `concurrency` available in most source integrations, and it can be specified as part of the integration source spec (Each spec is described in the relevant page in the [hub](https://hub.cloudquery.io/)). 

## Concurrency

`concurrency` provides rough control over the number of concurrent requests that will be made while performing a sync. Setting this to a low number will reduce the number of concurrent requests, reducing the memory used and making the sync less likely to hit rate limits. The trade-off is that syncs will take longer to complete.
