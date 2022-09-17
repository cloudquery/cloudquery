---
title: Rate Limiting
---

import { Callout } from 'nextra-theme-docs'

# Rate Limiting

You can configure CloudQuery to limit the amount of resources fetched in parallel, to prevent the remote provider API from being overwhelmed by too many requests.

CloudQuery currently allows setting `max_parallel_resource_fetch_limit` to limit how many resources are fetched simultaneously. This flag can be added to any
provider block as following:

```yaml
providers:
  - name: aws
    aws_debug: false
    // list of resources to fetch
    resources:
      - "*"
    // Limit provider to fetch only 5 resources at a given time
    max_parallel_resource_fetch_limit: 5
```

Another global setting is `max_goroutines` by default cloudquery sets this limit based on the current machine's specs based on the number of available cores. This flag can be added to any provider block to override the default cloudquery set limit.

```yaml
providers:
  - name: aws
    aws_debug: false
    // list of resources to fetch
    resources:
      - "*"
    // Limit cloudquery to limit maximum spawned goroutines for fetching to 5000.
    max_goroutines: 5000

```

<Callout type="info">

Some providers allow for more precise rate limiting and retry and backoff mechanisms. the AWS provider for example allows such [controls](/plugins/aws). Make sure to check
their configuration options [here](/plugins).

</Callout>
