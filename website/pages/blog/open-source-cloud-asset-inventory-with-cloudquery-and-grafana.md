---
title: Building an Open-Source Cloud Asset Inventory with CloudQuery and Grafana
tag: security
date: 2021/11/15
description: Visualise the cloud asset inventory from CloudQuery with Grafana
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

In this blog post you will learn how to build an open-source cloud asset inventory with CloudQuery and Grafana.

General architecture:

- [CloudQuery](https://github.com/cloudquery/cloudquery) will take care of extracting, transforming and loading all your assets, across cloud and SaaS apps to PostgreSQL.
- [Grafana](https://github.com/grafana/grafana) will be used to query, visualize, monitor, and alert.

This is what you will get:

- All your assets configuration across cloud providers and SaaS apps in **one** database
- Vanilla PostgreSQL
- Reuse your current (assuming you use Grafana) visualization, monitoring and alerting workflows - send reports and alerts via email, slack.
- 3 out-of-the-box filterable asset inventory dashboards for [AWS](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/dashboards) and [GCP](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/dashboards) including security & compliance dashboards.

## Running

To try out the below dashboards (or build a custom one) you will need to connect the same PostgreSQL used by CloudQuery to Grafana as a [datasource](https://grafana.com/docs/grafana/latest/datasources/postgres/).

You can run out CloudQuery locally or in your cloud environment.
For production deployment see [terraform-aws-cloudquery](https://github.com/cloudquery/terraform-aws-cloudquery) and/or helm-charts](https://github.com/cloudquery/helm-charts).

## Importing Dashboards

You can try out some of our pre-made dashboards by [importing](https://grafana.com/docs/grafana/latest/dashboards/export-import/#import-dashboard) them straight from our GitHub: [AWS](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/dashboards), [GCP](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/dashboards)

## Dashboard Examples and Use Cases

### Asset (Resource) Search

One of the most common but a lot of times non-trivial tasks is to find a specific when the only identifier you have is one piece of information across accounts, regions and resource types.
Just a few examples (as number of real-world scenarios is really indefinite ):

- **Finding an asset across accounts/regions by name/ARN:** In AWS specifically it might involve either clicking through 30+ regions, if you know in which account it is located or even more if not.
- **Finding an ec2 instance by its public/private IP:** This will also either involve click-ops, or bash sorcery.

Some of those can be also solved by AWS Config but has the following limitations:

- **AWS Only** - Works only on AWS resources (can’t ingest data from other services/cloud-providers).
- As this is using a proprietary subset of SQL, it can’t be integrated to your current Visualization, Monitoring Alerts workflows, such as Grafana.

This is why we created open-source **“basic inventory”** Grafana dashboards that you are free to use, customize or build completely new ones (feel free to share back or suggest other):

#### Filterable AWS dashboards by accounts and regions

Here is snippet from our AWS Asset Inventory Dashboard:
![A Grafana dashboard of AWS EC2 data](/images/blog/open-source-cloud-asset-inventory-with-cloudquery-and-grafana/image1.png "A Grafana dashboard of AWS EC2 data")

#### Filterable GCP dashboards across projects

Similar challenge exists in GCP though the situation is a bit better in some sense as you have a single view for each resource/asset type per project. Though, If you want to have a single view of all types of assets across multiple projects (which is common), this would be ClickOps or bash magic again.

Here is an example of GCP Compute Asset inventory dashboard:
![A Grafana dashboard of GCP Compute Asset data](/images/blog/open-source-cloud-asset-inventory-with-cloudquery-and-grafana/image3.png "A Grafana dashboard of GCP Compute Asset data")

Some of that can be solved with the in-house GCP Cloud Asset Inventory but has similar limitations:

- **GCP Only** - Works only on GCP resources (can’t ingest data from other services/cloud-providers).
- **Custom Query Language (not SQL)** - You will need to learn a new query engine that might be also more limited then SQL.
- Can’t integrate with best-in-class visualization, monitoring alerting systems such as Grafana.
- Cannot integrate across different organization accounts

### Security

You can create your own security views and dashboards that you can then monitor and alert. Each company has its own security and compliance policies but we will share a basic one in this blog (more is coming…) we found useful.

#### AWS Public/Private EC2 Instances

Filterable dashboards by **VPN**, **subnet**, **region** including public, private ec2 instances.
![A Grafana dashboard of Public/Private IPs for AWS EC2 instances](/images/blog/open-source-cloud-asset-inventory-with-cloudquery-and-grafana/image2.png "A Grafana dashboard of Public/Private IPs for AWS EC2 instances")

## Summary

We are excited for the future of open-source cloud asset inventory and are looking for your feedback, either on [GitHub](https://github.com/cloudquery/cloudquery) or [Discord](https://www.cloudquery.io/discord). Also, feel free to contribute back or request additional Grafana dashboards.
