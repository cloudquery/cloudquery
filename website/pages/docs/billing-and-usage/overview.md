---
title: Billing and Usage | Overview
description: Learn about CloudQuery's billing model
---

# Billing and Usage Overview

## How CloudQuery Pricing Works

The basic pricing model is a usage-based model: You get billed per synced row. A row means a resource synced with a source plugin that creates or updates a single row in the destination.

If you sync from a single source to multiple destinations, you get charged only once for the rows synced.

### Plugin Categories and a Free Quota

CloudQuery bills you at the end of each calendar month based on your usage during that month: the total amount of rows that you synced that month in each of the plugin categories. The pricing tiers in the table above indicate the cost per million rows for the specified amount of rows in each tier. The more you sync, the cheaper the rows get. After the free tier, the next 9 million rows are priced as indicated in the second tier, the next 90 million as indicated in the third tier, and so on.

Plugins are split in two categories: API plugins and Database plugins. Each category has its own price per row and the price goes down the more you sync within a month. See the [Pricing](https://www.cloudquery.io/pricing) page for the actual prices.

Some source plugins are [free](https://hub.cloudquery.io/plugins/source?tiers=free) to use without charges. Similarly, your own plugins created with CloudQuery SDK are free to use.

Some plugins (such as AWS, GCP, or Azure) have some tables that are free to sync. These tables usually contain static metadata that does not change often and have large amount of rows. These tables are marked `Free` in the individual plugin documentation.

### Additional charges for cloud syncs

When running syncs with CloudQuery Cloud, you may incur additional costs in the form of Egress, vCPU, and vRAM.

### Monitoring your usage

You can see your exact consumption in CloudQuery Cloud's billing page with details on the rows synced with individual plugins broken down by day, and how much of the free quota you have used in the current month.

![Usage chart](/images/docs/billing-and-usage/usage-chart.png)

If you used cloud syncs, your egress, vCPU, and vRAM consumption will be displayed at the bottom of the page.

![Cloud syncs usage](/images/docs/billing-and-usage/cloud-syncs.png)
