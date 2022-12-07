---
title: >-
  How to Build Open Source Cloud Asset Inventory with CloudQuery and Apache
  Superset (Preset)
tag: tutorial
date: 2022/06/01
description: >-
  How to setup CloudQuery to build your cloud asset inventory in PostgreSQL and
  connect it to Apache Superset (or a hosted version such as preset.io) for
  visualization, monitoring and reporting.
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

In this blog post, we will walk you through how to setup CloudQuery to build your cloud asset inventory in PostgreSQL and connect it to [Apache Superset](https://superset.apache.org/) (or a hosted version such as [preset.io](https://preset.io)) for visualization, monitoring and reporting.

## General Architecture

- **ETL (Extract-Transform-Load) ingestion layer:** [CloudQuery](https://github.com/cloudquery/cloudquery)
- **Datastore:** PostgreSQL
- **Data Visualization and Exploration Platform:** [Apache Superset](https://github.com/apache/superset)

## What you will get

- **Raw SQL access** to all your cloud asset inventory to create views or explore any questions or connection between resources.
- **Multi-Cloud Asset Inventory:** Ingest configuration from all your clouds to a single datastore with a unified structure.
- **Avoid yet-another-dashboard fatigue:** Use you existing [Apache Superset](https://superset.apache.org/) setup as a best in class visualization platform to build cloud asset inventory.

## Walkthrough

### Step 1: Install or Deploy CloudQuery

If it’s your first time using CloudQuery we suggest you first run it locally to get familiar with the tool, take a look at our [quickstart guide](/docs/quickstart) and [AWS source plugin](/docs/plugins/sources/aws/overview).

If you are already familiar with CloudQuery, take a look at how to deploy it to AWS on Amazon Aurora and EKS at [here](https://github.com/cloudquery/terraform-aws-cloudquery).

### Step 2: Install or sign up to Superset

[Apache Superset](https://github.com/apache/superset) is an open source Data Visualization and Data Exploration Platform so there are a number of ways to deploy it:

- **Self-hosted (local, docker, k8s):** Official guide.
- **SaaS/managed:** [Preset.io](https://preset.io/)
- **Cloud Marketplaces:**
  - [https://aws.amazon.com/quickstart/architecture/apache-superset/](https://aws.amazon.com/quickstart/architecture/apache-superset/)

### Step 3: Connecting Apache Superset (Preset) to PostgreSQL

By default RDS Aurora instances are not accessible from the public internet. In order to enable access by Preset you are going to have to update your security groups to include the IP ranges that Preset publishes (or alternatively look at the `publicly_accessible` variable in our terraform modules - [aws](https://github.com/cloudquery/terraform-aws-cloudquery#inputs) , [GCP](https://github.com/cloudquery/terraform-gcp-cloudquery#inputs)). If you deploy it in your own VPC you might be able to connect it in your private network.

Now you can connect preset to your PostgreSQL database by clicking **“New Database”**, choosing **PostgreSQL** and filling-in the following form:

![](/images/blog/cloud-asset-inventory-cloudquery-apache-superset/1.png)

If all your credentials are correct, click **connect** and then **finish**

### Step 4: Create a Dataset

To be able to create **charts** that you will later add to a **dashboard** you first need to create a **dataset** (for full details see [Creating Charts and Dashboards in Superset](https://superset.apache.org/docs/creating-charts-dashboards/creating-your-first-dashboard)).

In this guide we will create a dataset from a [view](https://github.com/cloudquery/cq-provider-aws/blob/main/views/resources.sql) (also, see [blog](https://www.cloudquery.io/blog/aws-resources-view)) we already created on our database but you can also create a dataset from any complex query you can think off in **Superset SQL Lab**.

![](/images/blog/cloud-asset-inventory-cloudquery-apache-superset/2.png)

### Step 5: Create your first chart!

Now you are ready to create your first chart on top of our [AWS Resource View](https://github.com/cloudquery/cq-provider-aws/blob/main/views/resources.sql) (Also, see [blog](https://www.cloudquery.io/blog/aws-resources-view))! Choose the dataset you created in the last step (in this case `aws_resources`) and choose the chart type you want to create (we will take the **Bar Chart** but you can also change this in the chart screen).

![](/images/blog/cloud-asset-inventory-cloudquery-apache-superset/3.png)

If you are familiar with the Superset UI this step is pretty easy where you just choose or drag and drop the columns you want to visualize (and you can always revert to plain SQL if needed).

![](/images/blog/cloud-asset-inventory-cloudquery-apache-superset/4.png)

### Step 6: Save & Add to dashboard

Click **Save** and choose the **Dashboard** you want to add the chart (or create a new one if you don’t have one) and then click **Save** again or **Save & Go to Dashboard**

![](/images/blog/cloud-asset-inventory-cloudquery-apache-superset/5.png)

You can now repeat steps **4-5** or **3-5** to create all charts you need and add them to your dashboard/s.

### Step 7: Edit Dashboard, Add Filters and Share and Alert!

Now we are at the final stage where you can edit, resize and tweak the charts/widgets so it looks something like this:

![](/images/blog/cloud-asset-inventory-cloudquery-apache-superset/6.png)

You can add **Filters** to make the dashboard more interactive and you can now share this with other team members, send links, image of dashboards as well as periodic alerts and reports via Superset.

## Summary

In this post we showed you how to build an open-source cloud asset inventory with CloudQuery as the ETL (Extract-Transform-Load) / data-ingestion layer and Apache Superset as the visualization and monitoring platforms. This approach eliminates the yet-another-dashboard fatigue and gives you the ability to pick the best-in-class visualization tools or reuse your current stack.
