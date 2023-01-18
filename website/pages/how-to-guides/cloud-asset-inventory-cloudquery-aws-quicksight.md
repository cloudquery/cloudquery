---
title: >-
  How to Visualize CloudQuery Data with AWS QuickSight
tag: tutorial
description: >-
  How to set up CloudQuery to build your cloud asset inventory in PostgreSQL and
  connect it to AWS QuickSight for visualization, monitoring and reporting.
author: yevgenypats
---

import { HowToGuideHeader } from "../../components/HowToGuideHeader"

<HowToGuideHeader/>

In this guide, we will walk you through how to set up CloudQuery to build your cloud asset inventory in PostgreSQL and connect it to [AWS QuickSight](https://aws.amazon.com/quicksight/) for visualization, monitoring and reporting.

## General Architecture

- **ETL (Extract-Transform-Load) ingestion layer:** [CloudQuery](https://github.com/cloudquery/cloudquery)
- **Datastore:** PostgreSQL
- **Data Visualization and Exploration Platform:** [AWS QuickSight](https://aws.amazon.com/quicksight/)

## What you will get

- **Raw SQL access** to all your cloud asset inventory to create views or explore any questions or connection between resources.
- **Multi-Cloud Asset Inventory:** Ingest configuration from all your clouds to a single datastore with a unified structure.
- **Avoid yet-another-dashboard fatigue:** Reuse your existing AWS QuickSight setup to build a cloud asset inventory. AWS QuickSight is not cheap but if you already use it won’t have additional cost (as the cost is per user) and you will re-use your current permissions and workflows.

## Walkthrough

### Step 1: Install or Deploy CloudQuery

If it’s your first time using CloudQuery we suggest you first run it locally to get familiar with the tool, take a look at our [quickstart guide](/docs/quickstart) and [AWS source plugin](/docs/plugins/sources/aws/overview).

If you are already familiar with CloudQuery, take a look at how to deploy it to AWS on Amazon Aurora and EKS at [here](https://github.com/cloudquery/terraform-aws-cloudquery).

### Step 2: Connecting AWS QuickSight to PostgreSQL

You can connect QuickSight to a private PostgreSQL (RDS Aurora) if you are on the enterprise edition. See [full walkthrough](https://aws.amazon.com/premiumsupport/knowledge-center/quicksight-redshift-private-connection/). If not, you can also make the Aurora publicly accessible but it is highly advised to only allow connections from [QuickSight IP Addresses](https://docs.aws.amazon.com/quicksight/latest/user/regions.html) .

Click Create New Dataset and choose PostgresSQL (In this tutorial we will connect to publicly accessible RDS with authorized QuickSight IP Address) and fill-in the connection details:

![](/images/blog/cloud-asset-inventory-cloudquery-aws-quicksight/1.png)

### Step 3: Visualize the Data!

Choose the table you want to visualize, in this case we will choose the`aws_resources` view which you need to create. To create or update the view, run [this query](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/views/resources.sql) (part of the CloudQuery AWS Source plugin source code) on your database. This will allow you to visualize all resources across all accounts and regions.

![](/images/blog/cloud-asset-inventory-cloudquery-aws-quicksight/2.png)

1. Now visualize away with QuickSight! Here is a small example of global filterable resources we created pretty quickly

![](/images/blog/cloud-asset-inventory-cloudquery-aws-quicksight/3.png)

You can reuse QuickSight to export/share those dashboards as well as create alerts!

## Summary

In this post we showed you how to build an open-source cloud asset inventory with CloudQuery as the ETL (Extract-Transform-Load) / data-ingestion layer and QuickSight as the visualization and monitoring platforms. This approach eliminates the yet-another-dashboard fatigue and gives you the ability to pick the best-in-class visualization tools and/or reuse your current stack.
