---
title: Billing and Usage | Overview
description: Learn about CloudQuery's billing model
---

# Billing and Usage Overview

## How CloudQuery Pricing Works

The basic pricing model is usage-based: You get billed per synced row. We charge when a record is pulled from a source, whether a record has changed or not.

If you sync from a single source to multiple destinations, you get charged only once for the rows synced.

### Plugin Categories and a Free Quota

CloudQuery bills you at the end of each calendar month based on your usage during that month: the total number of rows you synced that month in each integration category. The pricing tiers in the table above indicate the cost per million rows for the specified amount of rows in each tier. The more you sync, the cheaper the rows get. After the free tier, the next 9 million rows are priced as indicated in the second tier, the next 90 million as indicated in the third tier, and so on.

Plugins are divided into two categories: API integrations and Database integrations. Each category has a price per row, and the price decreases the more you sync within a month. For the actual prices, see the [Pricing](https://www.cloudquery.io/pricing) page.

Some source integrations are [free](https://hub.cloudquery.io/ integrations/source?tiers=free) and can be used without charges. Similarly, your integrations created with CloudQuery SDK are free to use.

Some integrations (such as AWS, GCP, or Azure) have tables that are free to sync. These tables usually contain static metadata that does not change often and have a large number of rows. They are marked `Free` in the individual integration documentation.

### Additional charges for cloud syncs

When running syncs with CloudQuery Cloud, you may incur additional costs in Egress, vCPU, and vRAM.

### Monitoring your usage

You can see your exact consumption on CloudQuery Cloud's billing page, which includes details on the rows synced with individual integrations broken down by day and how much of the free quota you used in the current month.

![Usage chart](/images/docs/billing-and-usage/usage-chart.png)

If you use cloud syncs, your egress, vCPU, and vRAM consumption will be displayed at the bottom of the page.

![Cloud syncs usage](/images/docs/billing-and-usage/cloud-syncs.png)

You can purchase CloudQuery directly through the [AWS Marketplace](https://aws.amazon.com/marketplace/pp/prodview-lowyuyay5a37s).
