---
title: >-
  How to Build Open Source Cloud Asset Inventory with CloudQuery and Microsoft
  Power BI
tag: tutorial
date: 2022/07/10
description: >-
  How to setup CloudQuery to build your cloud asset inventory in PostgreSQL and
  connect it to Microsoft Power BI for visualization, monitoring and reporting.
author: itay
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

In this blog post, we will walk you through how to setup CloudQuery to build your cloud asset inventory in PostgreSQL and connect it to [Microsoft Power BI](https://powerbi.microsoft.com/) for visualization, monitoring and reporting.

## General Architecture

- **ETL (Extract-Transform-Load) ingestion layer:** [CloudQuery](https://github.com/cloudquery/cloudquery)

- **Datastore:** PostgreSQL

- **Data Visualization and Exploration Platform:** [Microsoft Power BI](https://powerbi.microsoft.com/)

## What you will get

- **Raw SQL access** to all your cloud asset inventory to create views or explore any questions or connection between resources.

- **Multi-Cloud Asset Inventory:** Ingest configuration from all your clouds to a single datastore with a unified structure.

- **Avoid yet-another-dashboard fatigue:** Reuse your existing Power BI setup to build a cloud asset inventory.

## Walkthrough

### Step 1: **Install or Deploy CloudQuery**

If it’s your first time using CloudQuery we suggest you first run it locally to get familiar with the tool, take a look at our [quickstart guide](/docs/quickstart) and [Azure recipe](/docs/recipes/sources/azure).

If you are already familiar with CloudQuery, take a look at how to deploy it to Azure on AKS at [https://github.com/cloudquery/terraform-azure-cloudquery](https://github.com/cloudquery/terraform-azure-cloudquery).

### Step 2: Downloading Microsoft Power BI

Unfortunately, Power BI does not support the PostgresSQL connector from the web application.

We can still use the PostgresSQL connector by downloading the Power BI Desktop from [here](https://www.microsoft.com/en-us/download/details.aspx?id=58494).

### Step 3: Connecting Microsoft Power BI to PostgreSQL

Connection to Local / Remote server requires you to authorize the IP of the machine you’re running Power BI Desktop on.

See connection [full walkthrough](https://docs.microsoft.com/en-us/power-query/connectors/postgresql).

Click `Get Data` and choose PostgresSQL database (In this tutorial we will connect to publicly accessible PostgresSQL server with authorized IP) and fill-in the connection details:

![](/images/blog/cloud-asset-inventory-cloudquery-microsoft-power-bi/image0.png)

### Step 4: Visualize the Data!

Choose the table you want to visualize, in this case we will choose the `azure_resources` view.

> 💡 To create the `azure_resources` view, run the following [view](https://github.com/cloudquery/cq-provider-azure/blob/main/views/resource.sql) before importing to the data studio.

**Choose the table to visualize**

![](/images/blog/cloud-asset-inventory-cloudquery-microsoft-power-bi/image1.png)

**Design your report**

![](/images/blog/cloud-asset-inventory-cloudquery-microsoft-power-bi/image2.png)

You can reuse Power BI to export/share those reports as well!

## Summary

In this post we showed you how to build an open-source cloud asset inventory with CloudQuery as the ETL (Extract-Transform-Load) / data-ingestion layer and Microsoft Power BI as the visualization platforms. This approach eliminates the yet-another-dashboard fatigue and gives you the ability to pick the best-in-class visualization tools and/or reuse your current stack.
