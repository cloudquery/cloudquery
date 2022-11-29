---
title: Introducing Wildcard Matching for Tables
tag: product
date: 2022/11/11
description: 
author: hermanschaaf
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

I am excited to introduce a feature that quietly rolled out today: wildcard matching for tables in CloudQuery config files!

## Before

Before, the config passed to `cloudquery sync` had two options: either specify a wildcard (`*`) to match all tables, or specify a list of tables explicitly. So for example, if we wanted to include all EC2 resources, we would have to write a config like this:

```yaml
tables:
- aws_ec2_byoip_cidrs
- aws_ec2_customer_gateways
- aws_ec2_ebs_snapshots
- aws_ec2_ebs_volumes
- aws_ec2_egress_only_internet_gateways
- aws_ec2_eips
- aws_ec2_flow_logs
- aws_ec2_hosts
# ... plus ~21 more, omitted here to keep this post short
```

As you can see, this was quite verbose!

## Now

Config files now support wildcard matching. This means you can add a wildcard character (`*`) anywhere, and CloudQuery will include all tables that match. We can rewrite the previous EC2 config like this now:

```yaml
tables: 
- "aws_ec2_*"
```

Much easier! This also works in the `skip_tables` field. Let's say we want to include all EC2 tables, but not EBS-related tables. The config for this would now be:

```yaml
tables: 
- "aws_ec2_*"
skip_tables:
- "aws_ec2_ebs_*"
```

## Skipping relations

With the same change also came the ability to skip relations. By default, if a table is matched, all of its descendants will also be synced. So for example, this config:

```yaml
tables: 
- "aws_ec2_transit_gateways"
``` 

will populate `aws_ec2_transit_gateways`, as well as all its descendent tables (`aws_ec2_transit_gateway_attachments`, `aws_ec2_transit_gateway_route_tables`, etc). This was always the case. But now, using skip tables, some (or all) of these relations can be individually skipped:

```yaml
tables: 
- "aws_ec2_transit_gateways"
skip_tables:
- "aws_ec2_transit_gateway_*"  # will skip all child relations of aws_ec2_transit_gateways
```

Since child tables often have many rows and can require thousands API calls to fetch, skipping these when they're not needed can really speed up the time it takes to sync.

## Keeping up-to-date

Another nice thing about wildcard matching is that you can use it to automatically sync new tables that get added to CloudQuery. Using our example from above, this config:

```yaml
tables: 
- "aws_ec2_*"
```

will not only sync all the EC2 tables today, it will also automatically include any new ones when the AWS plugin version is upgraded. This is a good way to always have data from the latest APIs exposed by the cloud services you care about (along with reviewing CloudQuery plugin changelogs once in a while, of course).

## A note about versions

Wildcard matching was introduced as a backwards-compatible change, so no config updates are necessary when you upgrade to a version with support for wildcard matching. Wildcard matching is supported in the latest version of all official source plugins. For the big three cloud providers, it was first released in the following versions, and will be supported in these versions and higher:

 - AWS [v4.15.0](https://github.com/cloudquery/cloudquery/releases/tag/plugins-source-aws-v4.15.0)
 - GCP [v2.4.15](https://github.com/cloudquery/cloudquery/releases/tag/plugins-source-gcp-v2.4.15)
 - Azure: [v1.4.4](https://github.com/cloudquery/cloudquery/releases/tag/plugins-source-azure-v1.4.4)
 
See [this link](/docs/plugins/sources/overview) for a list of the latest versions of all official source plugins.

## Conclusion

With wildcard matching support, we're hoping to reduce a lot of the boilerplate required to configure tables to sync with CloudQuery. If you have any thoughts or feedback about the feature, please share them as a [GitHub issue](https://github.com/cloudquery/cloudquery) or chat to us on [Discord](https://cloudquery.io/discord)! ❤️
