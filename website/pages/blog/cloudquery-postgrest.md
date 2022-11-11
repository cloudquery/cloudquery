---
title: How To Run CloudQuery with PostgREST
tag: tutorial
date: 2022/06/26
description: See everything you have in the cloud with PostgREST
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

In this blog post, we will walk you through how to set up [CloudQuery](https://github.com/cloudquery/cloudquery) to build your cloud asset inventory in PostgreSQL and build a fully automated Restful API query layer with [PostgREST](https://postgrest.org/en/stable/). This can be used as a base for many custom use cases: from infrastructure search to security, cost and infrastructure automation.

## General Architecture

- **ETL (Extract-Transform-Load) ingestion layer:** [CloudQuery](https://github.com/cloudquery/cloudquery)

- **Datastore:** PostgreSQL

- **API Access Layer:** [PostgREST](https://postgrest.org/en/stable/) and any of your favorite API browsers if you wish (like [Postman](https://www.postman.com/downloads/?utm_source=postman-home) or any other)

## What You Will Get

- **Raw SQL access** to all of your cloud asset inventory to create views or explore any questions or connections between resources.

- **Multi-Cloud Asset Inventory:** Ingest configuration from all your clouds to a single datastore with a unified structure.

- **Rest API Endpoint** to access and query all your cloud configurations.

## Walkthrough

### Step 1: **Install or Deploy CloudQuery**

If it’s your first time using CloudQuery we suggest you first run it locally to get familiar with the tool. Take a look at our [quickstart guide](/docs/quickstart).

If you are already familiar with CloudQuery, take a look at how to deploy it to AWS on RDS Aurora and EKS at [github.com/cloudquery/terraform-aws-cloudquery](https://github.com/cloudquery/terraform-aws-cloudquery) , or GCP and Cloud SQL at [https://github.com/cloudquery/terraform-gcp-cloudquery](https://github.com/cloudquery/terraform-gcp-cloudquery)

### Step 2: Install PostgREST

Full full details, checkout [PostgREST](https://postgrest.org/en/stable/tutorials/tut0.html) docs. If you are on mac you can install it via `brew install postgrest`

To run it locally, all you need is the following `cq.conf` file as input for PostgREST (adjust the PG URL accordingly):

```
db-uri = "postgres://postgres:pass@localhost:5432/postgres"
db-schemas = "public"
db-anon-role = "postgres"
```

and run the following

```bash
postgrest cq.confg
```

### Step 3: Query and Profit!

That’s it! You should see something like the following in the output if all is well:

```bash
12/Jun/2022:23:36:20 +0300: Attempting to connect to the database...
12/Jun/2022:23:36:20 +0300: Connection successful
12/Jun/2022:23:36:20 +0300: Listening on port 3000
12/Jun/2022:23:36:20 +0300: Config re-loaded
12/Jun/2022:23:36:20 +0300: Listening for notifications on the pgrst channel
12/Jun/2022:23:36:20 +0300: Schema cache loaded
```

Now you can query the endpoint with `curl` or any other API browser/UI, Swagger UI.

```bash
curl "http://localhost:3000/aws_ec2_instances"
```

You can also use any filter on any of the fields which PostgREST automatically exposes.

```bash
curl "http://localhost:3000/aws_ec2_instances?arn=eq.arnsomething"
```

### Step 4: Create New Views

By default, PostgREST exposes all tables and relationships of the existing table. But let’s say you want to create a new view. All you need to do is create the new view, and PostgREST will automatically generate the model for that. For example, checkout this [blog](https://www.cloudquery.io/blog/aws-resources-view) on how to create a unified AWS resource [view](https://github.com/cloudquery/cq-provider-aws/tree/main/views) (or GCP [View](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/gcp/views)). And just like that you can now query and search all your resources by ARN, tags, or name using GraphQL!

### Step 5: Deploying in production

If you want to expose PostgREST publicly please see [PostgREST Security](https://postgrest.org/en/stable/auth.html). Or, expose it privately and use either a bastion host or something like [Tailscale Kubernetes](https://tailscale.com/kb/1185/kubernetes/) together with our [helm charts.](https://github.com/cloudquery/helm-charts)

## Summary

In this post we showed you how to build an open-source cloud asset inventory with CloudQuery as the ETL (Extract-Transform-Load) / data-ingestion layer and [PostgREST](https://postgrest.org/) as the API layer to expose the data for your internal team/users or any other downstream processing in the most convenient/preferred way.
