---
title: Scaling out ELT pipelines with CloudQuery and CSV
tag: tutorial
date: 2022/11/01
description: >-
  This tutorial will show you how to scale out your ELT pipelines with CloudQuery to infinity and beyond!
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

Performance is a key factor in data pipelines both from usability and cost perspective. This tutorial will show you how to scale out your ELT pipelines with CloudQuery and CSV utilizing CloudQuery's stateless and performant architecture.

## Background

Let's take a look at your typical ELT pipeline:

```bash
[Source] -> [Database]
```

This is the basic architecture where you have one instance running your ELT workload scaling vertically, and one database instance that needs to also scale vertically to handle the stream/load. This is a common architecture and it works well for small to medium workloads. However, as the workload grows you will need to scale out your ELT workload. This is where things get complicated and expensive.

```bash
[Source] --\
[Source]    -> [Database]
[Source] --/
```

Scaling out sources now creates a bottleneck where you need to scale out your database and/or have things like PgBouncer (in-case of PostgreSQL) or use scale-out databases like CockroachDB. If you use data lakes or data warehouses (such as BigQuery) directly, you can also hit streaming limits because they tend to work better with batch loads. But again you will need some kind of intermediate server that batches all requests from all sources and then loads those to the data warehouse.


## Scaling out with CloudQuery and CSV

CloudQuery workers are stateless, so you can run as many as you want and slice and dice what data they are fetching as you need without them connecting to any database and/or backend. With the new CloudQuery [CSV Destination Plugin](https://github.com/cloudquery/cloudquery/releases/tag/plugins-destination-csv-v1.0.0) each worker can write the results to a local CSV file, from where you can upload it to a cloud storage of your choice. This way you can scale out your ELT workload to infinity and beyond and only pay for the storage you use.

```bash
[Source] --\
[Source]    -> [Cloud Storage] -> [Data lake/Data warehouse]
[Source] --/
```

Loading CSV Files to data lakes or data warehouses is fast and cheap without the need to take care of scaling out a database or an intermediate proxy vertically.

## Summary

Scaling out ELT workloads with CSV and Cloud Storage will work well especially if you extract data from high number of sources and accounts, but it has the downside that data in the destination won't be as live when compared to streaming directly to a database. So as always, all depends on your use-case and constraints. Check out the [CSV plugin documentation](/docs/plugins/destinations/csv/overview) to see how to configure the CSV destination plugin.
