---
title: >-
  How to Visualize CloudQuery Data with Google Data Studio
tag: tutorial
description: >-
  How to set up CloudQuery to build your cloud asset inventory in PostgreSQL and
  connect it to Google Data Studio for visualization, monitoring and reporting.
author: itay
---

import { HowToGuideHeader } from "../../components/HowToGuideHeader"

<HowToGuideHeader/>

In this guide, we will walk you through how to set up CloudQuery to build your cloud asset inventory in PostgreSQL and connect it to [Google Data Studio](https://marketingplatform.google.com/about/data-studio/) for visualization, monitoring and reporting.

## General Architecture

- **ETL (Extract-Transform-Load) ingestion layer:** [CloudQuery](https://github.com/cloudquery/cloudquery)
- **Datastore:** PostgreSQL
- **Data Visualization and Exploration Platform:** [Google Data Studio](https://marketingplatform.google.com/about/data-studio/)

## What you will get

- **Raw SQL access** to all your cloud asset inventory to create views or explore any questions or connection between resources.
- **Multi-Cloud Asset Inventory:** Ingest configuration from all your clouds to a single datastore with a unified structure.
- **Avoid yet-another-dashboard fatigue:** Reuse your existing Google Data Studio setup to build a cloud asset inventory.

## Walkthrough

### Step 1: **Install or Deploy CloudQuery**

If it’s your first time using CloudQuery we suggest you first run it locally to get familiar with the tool, take a look at our [quickstart guide](/docs/quickstart) and [GCP source plugin](/docs/plugins/sources/gcp/overview).

If you are already familiar with CloudQuery, take a look at how to deploy it to GCP on Cloud SQL and GKE at [https://github.com/cloudquery/terraform-gcp-cloudquery](https://github.com/cloudquery/terraform-gcp-cloudquery).

### Step 2: Connecting Google Data Studio to PostgreSQL

You can only connect Data Studio to a public PostgreSQL (GCP Cloud SQL).

For security purpose you should allow connection only from [Google Data Studio IP Addresses](https://support.google.com/datastudio/answer/7288010?hl=en#zippy=%2Cin-this-article%2Cshow-the-list-of-ip-addresses).

See connection [full walkthrough](https://support.google.com/datastudio/answer/7288010#zippy=%2Cin-this-article).

Click Create New datasource and choose PostgresSQL (In this tutorial we will connect to publicly accessible RDS with authorized Data Studio IP Address) and fill-in the connection details:

![](/images/blog/cloud-asset-inventory-cloudquery-google-data-studio/step2.png)

### Step 3: Visualize the Data!

Choose the table you want to visualize, in this case we will choose the `gcp_resources` view.

💡 To create the `gcp_resources` view, run the following [view](https://github.com/cloudquery/cq-provider-gcp/blob/main/views/resource.sql) before importing to the data studio.

**Choose the table to visualize**
![](/images/blog/cloud-asset-inventory-cloudquery-google-data-studio/step3-1.png)

**Design your report**
![](/images/blog/cloud-asset-inventory-cloudquery-google-data-studio/step3-2.png)

You can reuse Data Studio to export/share those reports as well!

## Summary

In this post we showed you how to build an open-source cloud asset inventory with CloudQuery as the ETL (Extract-Transform-Load) / data-ingestion layer and Google Data Studio as the visualization platforms. This approach eliminates the yet-another-dashboard fatigue and gives you the ability to pick the best-in-class visualization tools and/or reuse your current stack.
