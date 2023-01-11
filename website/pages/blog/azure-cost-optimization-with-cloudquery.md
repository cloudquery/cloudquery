---
title: Correlating Costs to Resources for Azure Cost Optimization with CloudQuery
tag: tutorial
date: 2023/01/04
description: >-
  This tutorial will show how to correlate between Azure Cost Analysis data and CloudQuery to optimize cost
author: kemal
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

## Introduction

Optimizing cloud costs is a never-ending task. There are many ways to do it, and many tools to help you do it. One of the most important things to do is to correlate your cloud costs with your cloud infrastructure. This tutorial will show how to correlate between Azure Cost Analysis data and CloudQuery to optimize cost.

## Prerequisites

- [Azure Source Plugin](/docs/plugins/sources/azure/overview)
- [PostgreSQL Destination Plugin](/docs/plugins/destinations/postgresql/overview)
- A PostgreSQL instance (can be local or remote)

## Create a Cost Analysis Report in Azure

First, we need to create a _Cost Analysis Report_ in Azure. To do that, go to the Cost Management section in the Azure Portal and click on the [Cost Analysis](https://portal.azure.com/#view/Microsoft_Azure_CostManagement/Menu/~/costanalysis/openedBy/AzurePortal) tab.

![Azure Cost Analysis Start Page](/images/blog/azure-cost-optimization-with-cloudquery/cost-analysis-cost-by-resource-report.png)

Then, click on the _Cost by Resource_ button.

![Azure Cost By Resource Menu](/images/blog/azure-cost-optimization-with-cloudquery/cost-analysis-cost-by-resource-menu.png)

Here you can make adjustments to your report. Make sure to select the time range you want to analyze. For this tutorial, we'll leave this at the default of current billing period. Make sure the _Group By_ option is set to **Resource**.

![Azure Cost By Resource Report](/images/blog/azure-cost-optimization-with-cloudquery/cost-analysis-cost-by-resource-report-annotated.png)

Use the _Save as_ button to save this as a new report view. Give it a name and click _Save_.

![Azure Cost By Resource Report Save Button](/images/blog/azure-cost-optimization-with-cloudquery/cost-analysis-cost-by-resource-report-save.png)

Now we can configure CloudQuery so each time we run `cloudquery sync` the data from this report will be synced to a PostgreSQL database.


## Syncing data

Before going off and writing queries, we need to sync the data from Azure to our PostgreSQL database. To do that, we'll use the `cloudquery sync` command. Let's start with creating a configuration file (For the full configuration reference, check out the [Azure Source Plugin](/docs/plugins/sources/azure/configuration) and [PostgreSQL Destination Plugin](/docs/plugins/destinations/postgresql/overview)):

```yaml copy
kind: source
spec:
  # Source spec section
  name: "azure"
  path: "cloudquery/azure"
  version: "VERSION_SOURCE_AZURE"
  destinations: ["postgresql"]
  concurrency: 100 # The Azure source plugin supports many resources and can fail on small machines. Increase this if you run it on a beefy machine
  tables: ["*"] # "azure_costmanagement_views" and "azure_costmanagement_view_queries" are required for this tutorial, but let's get them all
---
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "VERSION_DESTINATION_POSTGRESQL"
  spec:
    # You will need to edit this to specify your own DSN
    connection_string: "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
```

Save this file as `config.yml` and run `cloudquery sync ./config.yml`. This might take a hot minute or two, depending on how many subscriptions you have and how many resources you have in them.

Once the data is synced we can move on to the next section.

## Understanding the Cost Report

Now that we have the data synced, we can start writing queries. Let's start with a simple query to see what the data looks like. We'll start with the `azure_costmanagement_views` table. This table contains the data from the Cost Analysis report we created earlier. Let's see what the data looks like:

```sql copy
SELECT * FROM azure_costmanagement_views LIMIT 1;
```

Focus on the `properties` column, which looks like this if formatted nicely:

![azure cost management properties column](/images/blog/azure-cost-optimization-with-cloudquery/azure_costmanagement_views_props_small.png)

As you can see this only includes metadata about our report. Report contents are nowhere to be seen here. For that, we need to look at the `azure_costmanagement_view_queries` table. Let's see what the data looks like:

```sql copy
SELECT * FROM azure_costmanagement_view_queries LIMIT 1;
```

Again, if we focus on the `properties` column and format it in JSON Viewer:

![azure cost management view queries properties column](/images/blog/azure-cost-optimization-with-cloudquery/azure_costmanagement_view_queries_props.png)

Looks like the data we're interested in is sitting nicely in properties.

We need to run one more thing before we can start correlating away though.

## The Azure Resources View

The `azure_resources` view is useful for seeing all your Azure resources in one place or finding a specific resource, by id or name. It's a view, so it's not synced to the database, but we can create it by running the queries provided [here](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/azure/views/resource.sql). First query will drop the view if it exists, and the second query will create it. Run these queries in your PostgreSQL database.

## Correlating the data

### First Steps

As the first query we should correlate our specific named report with the data in the `azure_costmanagement_view_queries` table. We can do this by using the `azure_costmanagement_views` table to get the `_cq_id` of our report, and then using that `cq_id` (matching it with `_cq_parent_id`) to find the data in the `azure_costmanagement_view_queries` table. Something like:

```sql copy
SELECT v.*, q.* FROM azure_costmanagement_views v, azure_costmanagement_view_queries q WHERE v.name = 'cost-matching' AND v._cq_id = q._cq_parent_id;
```

Keep in mind that the `cost-matching` name is the name we gave our report when we created it. If you named it something else, you'll need to change it here.

### Deeper Dive

After seeing actual data it will become apparent that we'll need to figure out the column indexes for the data we're interested in. We can do this by looking at the `properties.columns` data in the `azure_costmanagement_view_queries` table. We can use `JSONB_ARRAY_ELEMENTS WITH ORDINALITY` to match column names to indexes. Something like `JSONB_ARRAY_ELEMENTS(q.properties->'columns') WITH ORDINALITY cid(col, pos) ON cid.col->>'name'='ResourceId'` will give us the index of the ResourceId column as `cid.pos`. A full query to return _Resource Id_ vs. _Cost (USD)_ might look like:

```sql copy 
SELECT (r->>(cid.pos::int - 1))::text AS res_id, (r->>(ccost.pos::int - 1))::numeric AS cost_usd
FROM azure_costmanagement_view_queries q
JOIN JSONB_ARRAY_ELEMENTS(q.properties->'rows') AS r ON TRUE
JOIN JSONB_ARRAY_ELEMENTS(q.properties->'columns') WITH ORDINALITY cid(col, pos) ON cid.col->>'name'='ResourceId'
JOIN JSONB_ARRAY_ELEMENTS(q.properties->'columns') WITH ORDINALITY ccost(col, pos) ON ccost.col->>'name'='CostUSD'
JOIN azure_costmanagement_views v ON v._cq_id=q._cq_parent_id
WHERE v.name='cost-matching'
```

Let's break it down:

  * `SELECT (r->>(cid.pos::int - 1))::text AS res_id, (r->>(ccost.pos::int - 1))::numeric AS cost_usd` - This is the meat of the query. We're selecting the ResourceId and CostUSD columns from the `azure_costmanagement_view_queries` table. We're using the `cid.pos` and `ccost.pos` values we got from the previous query to get the correct column indexes. We're also casting the values to the correct types. `- 1` because we're using 1-based indexing, and the `rows` array is 0-based.
  * `FROM azure_costmanagement_view_queries q` - We're selecting from the `azure_costmanagement_view_queries` table, and we're calling it `q` for the rest of the query.
  * `JOIN JSONB_ARRAY_ELEMENTS(q.properties->'rows') AS r ON TRUE` - We're joining the `azure_costmanagement_view_queries` table to itself, and we're calling the joined table `r`. We're using `JSONB_ARRAY_ELEMENTS` to get the rows from the `properties` column. We're using `ON TRUE` because we don't need to filter the rows, we just need to join them.
  * `JOIN JSONB_ARRAY_ELEMENTS(q.properties->'columns') WITH ORDINALITY cid(col, pos) ON cid.col->>'name'='ResourceId'` - We're joining the `azure_costmanagement_view_queries` table to itself again, and we're calling the joined table `cid`. We're using `JSONB_ARRAY_ELEMENTS` to get the columns from the `properties` column. We're using `WITH ORDINALITY` to get the column index as well. We're using `ON cid.col->>'name'='ResourceId'` to filter the columns to only the ResourceId column.
  * `JOIN JSONB_ARRAY_ELEMENTS(q.properties->'columns') WITH ORDINALITY ccost(col, pos) ON ccost.col->>'name'='CostUSD'` - Same as above, but we're filtering to only the CostUSD column.
  * `JOIN azure_costmanagement_views v ON v._cq_id=q._cq_parent_id` - We're joining the `azure_costmanagement_views` table to the `azure_costmanagement_view_queries` table. We're using the `_cq_id` column to match the two tables.
  * `WHERE v.name='cost-matching'` - We're filtering the `azure_costmanagement_views` table to only the report we're interested in.

### Final Query

Now that we have a query that returns the `res_id` and `cost_usd` columns, we can use that to correlate the data with the `azure_resources` view. We can do this by using the above query to get the _Resource ID_, and then using that to find the data in the `azure_resources` view (match using `LOWER()`, as the ids don't always match case-wise). Something like:

```sql copy
WITH cost_by_res AS (
    SELECT (r->>(cid.pos::int - 1))::text AS res_id, (r->>(ccost.pos::int - 1))::numeric AS cost_usd
    FROM azure_costmanagement_view_queries q
             JOIN JSONB_ARRAY_ELEMENTS(q.properties->'rows') AS r ON TRUE
             JOIN JSONB_ARRAY_ELEMENTS(q.properties->'columns') WITH ORDINALITY cid(col, pos) ON cid.col->>'name'='ResourceId'
    JOIN JSONB_ARRAY_ELEMENTS(q.properties->'columns') WITH ORDINALITY ccost(col, pos) ON ccost.col->>'name'='CostUSD'
    JOIN azure_costmanagement_views v ON v._cq_id=q._cq_parent_id
WHERE v.name='cost-matching' -- <-- this is the name of your saved view with "resource" field
    )
SELECT
    c.*,
    r.name,
    r.kind,
    r.location
FROM cost_by_res c JOIN azure_resources r ON LOWER(r.full_id)=LOWER(c.res_id);
```

## Summary

In this blog post we just shared an example of what you can do by combining cost data with your infrastructure state/metadata synced by CloudQuery. The number of use cases around cost (aka "FinOps" :) ) is really infinite and it all depends on what you are trying to achieve and optimize for.

We hope you enjoyed this tutorial and found it useful. If you have any questions or feedback, please reach out to us on [Discord](https://www.cloudquery.io/discord) or [Twitter](https://twitter.com/cloudqueryio).
