---
title: Introducing AWS Resources View
tag: security
date: 2022/05/24
description: >-
  Here at CloudQuery we built a simple aws_resources view, to demonstrate the
  power of using a SQL database to create a single pane for all your fetched AWS
  resources.
author: roneliahu
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

As of the writing of this blog, CloudQuery supports over 155 resources across 61 services in AWS! and many more in GCP, Azure, etc. Although this gives users the capabilities to answer many of their questions regarding your security, visibility and infrastructure, it would be great to have a single view of all your AWS resources, right?

So here at CloudQuery we built a simple `aws_resources` view, to demonstrate the power of using a SQL database to create a single pane for all your fetched AWS resources. The `aws_resources` view allows us to ask questions on all our resources allowing us to filter by service, region, account and more!

## Getting Started

As always, before we create this view, you should check out our [quickstart guide](/docs/quickstart), and make sure you execute a sync.

After our AWS provider is set up and we executed our sync, run the following SQL in your database. This statement creates a view that extracts and transform each row in our `aws` tables with an `arn` column into a `aws_resource` form, and unites all of them rows into our singular view.

```sql
DROP VIEW IF EXISTS aws_resources;

DO $$
DECLARE
    tbl TEXT;
    strSQL TEXT = '';
BEGIN
-- iterate over every table in our information_schema that has an `arn` column available
FOR tbl IN
    SELECT table_name
    FROM information_schema.columns
    WHERE table_name LIKE 'aws_%s' and COLUMN_NAME = 'account_id'
    INTERSECT
    SELECT table_name
    FROM information_schema.columns
    WHERE table_name LIKE 'aws_%s' and COLUMN_NAME = 'arn'
LOOP 
    -- UNION each table query to create one view
 	IF NOT (strSQL = ''::TEXT) THEN
		strSQL = strSQL || ' UNION ALL ';
	END IF;
	-- create an SQL query to select from table and transform it into our resources view schema
	strSQL = strSQL || FORMAT('
        select _cq_id, _cq_source_name, _cq_sync_time, %L as _cq_table, account_id, %s as region, arn, %s as tags
        FROM %s',
        tbl,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='region' AND table_name=tbl) THEN 'region' ELSE E'\'unavailable\'' END,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='tags' AND table_name=tbl) THEN 'tags' ELSE '''{}''::jsonb' END,
        tbl);
END LOOP;

IF strSQL = ''::TEXT THEN
    RAISE EXCEPTION 'No tables found with ARN and ACCOUNT_ID columns. Run a sync first and try again.';
ELSE
	EXECUTE FORMAT('CREATE VIEW aws_resources AS (%s)', strSQL);
END IF;

END $$;

```

Up to date version of this query can be [found here](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/views).

### Run the following query to view all your AWS resources:

```sql
select * from aws_resources limit 100;
```

![List of All AWS Resources](/images/blog/aws-resources-view/all-resources.png)

## Example Queries

### What resources don’t have tags?

```sql
select * from aws_resources where tags = '{}';
```

![Table of AWS Resources that don't have tags](/images/blog/aws-resources-view/resources-without-tags.png)

### What resources don’t have any of these tags?

```sql
select * from aws_resources where not tags ?| array['name', 'version'];
```

We can easily invert this or make sure that `all` of these tags exist with the `?&` operator instead.

![Table of AWS Resources that don't have particular tags](/images/blog/aws-resources-view/resources-without-particular-tags.png)

### What resources of Type Z in service X exist in Region Y?

```sql
SELECT * FROM aws_resources WHERE region LIKE 'us-east%'
AND service = 'ec2' AND (type = 'instance' OR type = 'network-interface');
```

Here we can easily query all resources in the `ec2` service from the `us-east` regions, that they are of type `instance` and `network-interface`

![AWS Resources of type Z in service X in region Y](/images/blog/aws-resources-view/resources-of-type-z-in-service-x-in-region-y.png)

### Join To existing tables

```sql
SELECT instance_type, aws_resources.id, aws_resources.arn, launch_time,
	public_ip_address, private_ip_address, state_name, vpc_id FROM aws_resources
INNER JOIN aws_ec2_instances ON aws_resources.cq_id = aws_ec2_instances.cq_id
WHERE aws_resources.region LIKE 'us-east%' AND aws_resources.service = 'ec2' AND aws_resources.type = 'instance' AND aws_resources.tags = '{}'
```

We can easily create `join` to our existing tables to get more information, we can join either on the `cq_id` or event the `id` column. This allows to get more specific information on the resources,

in this case, `launch_time`, `public_ip_address`, `vpc_id` etc’.

### Count total distinct resources by ARN

```sql
select count(distinct arn) as distinct_resources, count(*) as total from aws_resources
```

## What's next?

There are many views we can create on top of CloudQuery that make it easier to query our data, some examples can be found in our [policies](/docs/core-concepts/policies). We are working on more awesome views that will make your life even easier, such as `aws_policies`

We are always excited to hear use cases or questions around CloudQuery so feel free to hop into our [discord](https://www.cloudquery.io/discord) and message us.
