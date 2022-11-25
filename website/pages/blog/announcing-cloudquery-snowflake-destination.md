---
title: Announcing the CloudQuery Snowflake Destination Plugin
tag: tutorial
date: 2022/11/24
description: >-
  This tutorial will show you how to sync your cloud resources to a Snowflake database.
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

## Introduction

Cloud Infrastructure data exploded over the last decade which created the need to store those in data-warehouses and datalakes for analysis and reporting enabled by cheap storage. Today I'm excited to announce the release of our new [Snowflake destination plugin](../docs/recipes/destinations/snowflake), which enables you to take advantage of [CloudQuery open source](https://github.com/cloudquery/cloudquery) high-performance data-integration platform and sync all supported CloudQuery [source plugins](../docs/plugins/sources) to a [Snowflake](https://www.snowflake.com/) database.

## Use Cases

### Cloud Infrastructure data-lake and data-warehouse

Snowflake is already used as a ["Security data-lake"](https://www.snowflake.com/en/data-cloud/workloads/cybersecurity/). CloudQuery brings all your cloud [infrastructure configuration data](../docs/plugins/sources) to the same place consolidating all your security and infrastructure data in one place and enabling new insights.

### Historical data

Maintaining historical data is a common use case for data-lakes and data-warehouses. By using Snowflake as a destination you can now store all your cloud infrastructure data for long term storage, analysis and investigation use-cases.

## Syncing data

Syncing data to snowflake can be done as with any other destination plugin, so check out our [Quickstart](../docs/quickstart) and [Snowflake recipe](../docs/recipes/destinations/snowflake).

There are two ways to sync data to Snowflake:

1. Direct (easy but not recommended for production or large data sets): This is the default mode of operation where CQ plugin will stream the results directly to the Snowflake database. There is no additional setup needed apart from authentication to Snowflake.

2. Loading via CSV/JSON from a remote storage: This is the standard way of loading data into Snowflake, it is recommended for production and large data sets. This mode requires a remote storage (e.g. S3, GCS, Azure Blob Storage) and a Snowflake stage to be created. The CQ plugin will stream the results to the remote storage. You can then load those files via a cronjob or via SnowPipe. This method is still in the works and will be updated soon with a guide.

## Summary

This is the first Snowflake destination plugin release with more improvments coming, and I'd love to hear your use-cases and ideas on how you leverage it. Feel free to open an issue on [GitHub](https://github.com/cloudquery/cloudquery) or join our [Discord](https://cloudquery.io/discord) to share your ideas.