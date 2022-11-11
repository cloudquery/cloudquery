---
title: Introducing Glob Matching for Tables
tag: product
date: 2022/11/11
description: 
author: hermanschaaf
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

I am excited to announce a small feature that quietly rolled out today: glob matching for tables in CloudQuery config files!

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

Config files now support glob matching. This means you can add a wildcard character (`*`) anywhere, and CloudQuery will include all tables that match. We can rewrite the previous EC2 config like this now:

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

Since child tables often have many rows and require many API calls to fetch, skipping these when they're not needed can really speed up the time it takes to sync.

## Keeping up-to-date

Another nice thing about glob matching is that you can use it to automatically sync new tables that get added to CloudQuery. Using our example from above, this config:

```yaml
tables: 
- "aws_ec2_*"
```

will not only sync all the EC2 tables today, it will also automatically include any new ones when the AWS plugin version is upgraded. This is a great way to keep on the cutting edge of the data exposed by the cloud services you care about.

## Versions

Glob matching is supported in the latest version of all official source plugins. For the big cloud providers, these are:

 - AWS: 
 - GCP: 
 - Azure: 
 - Github: 

