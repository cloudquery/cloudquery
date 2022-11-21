---
title: Announcing Policies in CloudQuery Hub
tag: announcement
date: 2021/08/09
description: Adding CQ Policies to CQ Hub
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"
import { Callout } from 'nextra-theme-docs'

<BlogHeader/>

<Callout type="warning">
HCL policies were deprecated - see up-to-date policy documentation [here](https://www.cloudquery.io/docs/core-concepts/policies).
</Callout>

Today, we’re excited to announce the availability of [CloudQuery Policies](https://www.cloudquery.io/blog/announcing-cloudquery-policies) in [CloudQuery Hub](https://hub.cloudquery.io/).

Less than two months ago we introduced [CloudQuery Hub](https://www.cloudquery.io/blog/announcing-cloudquery-hub) as a single place to browse, install and share CloudQuery pluggable providers and integrations.

Last month we introduced [CloudQuery Policies](https://www.cloudquery.io/blog/announcing-cloudquery-policies) that brought policy-as-code to the CloudQuery ecosystem. CQ Policies enable users to codify, version and run security, governance, cost and compliance rules, using SQL as the query layer and HCL as the logical layer.

## What’s Inside CloudQuery Policy Hub?

As a natural next step we created a link between the two to enable users to browse, install and share CQ Policies on CQ Hub. Together with this release we also open-sourced three frequently asked policies: [AWS CIS 1.2.0](https://github.com/cloudquery/cq-provider-aws/tree/main/policies/cis_v1.2.0), [Azure CIS 1.3.0](https://github.com/cloudquery/cq-provider-azure/tree/main/policies/cis_v1.3.0), [GCP CIS 1.2.0](https://github.com/cloudquery/cq-provider-gcp/tree/main/policies/cis_v1.2.0).

- **Native GitHub support**: All policies are backed and versioned by GitHub.
- **Query Browser**: All policies uploaded to CQ Hub get automatic documentation, versioning and listing of queries available in each policy.

## What’s next

As with every product/feature announcement, we are eager to hear feedback and ideas - feel free to hop into our [Discord Channel](https://www.cloudquery.io/discord) or open an issue on our [github](https://github.com/cloudquery/cloudquery).

For more updates, subscribe to our newsletter below and/or follow us on [twitter](https://twitter.com/cloudqueryio).
