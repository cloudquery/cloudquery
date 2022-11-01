---
title: Scaling out ELT pipelines with CloudQuery and CSV
tag: tutorial
date: 2022/11/01
description: >-
  This turoial will show you how to scale out your ELT pipelines with CloudQuery to infinity and beyond!
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

Performance is a key factor in data pipelines both from usability and cost perspective. This tutorial will show you how to scale out you ELT pipelines with CloudQuery and CSV utilizing CloudQuery statless and performant architecture.

## Background

Let's take a look at your typical ELT pipeline:

```
[Source] -> [Database]
```

This is the basic architecture where you have one machine/container running your ELT workload scaling vertically and one database machine that needs also to scale vertically to handle the stream/load. This is a very common architecture and it works well for small to medium workloads. However, as your workload grows you will need to scale out your ELT workload. This is where things get complicated and expensive.

```
[Source] --\
[Source]    -> [Database]
[Source] --/
```

Scaling out sources now creates a problem/bottlenack where you need to scale out your database and/or have things like pgbouncer (in-case of PostgreSQL) or use scale-out databases like CockroachDB. If you use datalakes, or data warehouses directly such as BigQuery you can also hit streaming limits because they tend to work better with batch loads and then again you will need some kind of intermediate server that batches all requests from all sources and then loads those to the data warehouse.


## Scaling out with CloudQuery and CSV

CloudQuery workers are statless so you can run as many as you want and slice and dice what data they are fetching as you need witout them connecting to any database and/or backend. With the new CloudQuery [CSV Destination Plugin](https://github.com/cloudquery/cloudquery/releases/tag/plugins-destination-csv-v1.0.0) each worker can write the results to a local CSV file and then upload it to a cloud storage of your choice. This way you can scale out your ELT workload to infinity and beyond and only pay for the storage you use.

```
[Source] --\
[Source]    -> [Cloud Storage] -> [Datalake/datawarehouse]
[Source] --/
```

Loading CSV to datalakes/data warehouse is very fast and cheap without the need to take care of scaling out vertically a database or an intermediate proxy.

## Summary

Scaling out ELT workloads with CSV and Cloud Storage will work very well especially if you extract data from high number of sources and accounts, but it will have the downside that data in the destination wont be as live compared to streaming directly to a database, so as always, all depends on the use-case and constraints.



