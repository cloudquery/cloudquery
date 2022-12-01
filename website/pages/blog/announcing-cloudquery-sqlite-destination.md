---
title: Announcing the CloudQuery SQLite Destination Plugin
tag: tutorial
date: 2022/11/11
description: >-
  This tutorial will show you how to sync your cloud resources to a SQLite database.
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

## Introduction

SQLite is all the rage these days, and for a good reason. It is an embedded database with no dependencies that can be stored as a file while having a lot of the features of a full relational database, such as SQL-based query language, indexes, data types, etc…

Today I'm excited to announce the release of our new SQLite destination plugin, so you can now sync all supported CloudQuery [source plugins](/docs/plugins/sources/overview) to a [SQLite](https://www.sqlite.org/index.html) database.

Checkout [our docs](/docs/recipes/destinations/sqlite) to see how to configure the SQLite destination plugin.

## Use Cases

SQLite use-cases in data integration are infinite, but I'd like to share a few interesting ideas.

### Local data exploration

If you don't want to run a database or docker, running SQLite as destination will be super performant similar to our [CSV destination plugin](/docs/recipes/destinations/csv) but with the ability to do complex queries and store everything in one file.

### Scaling ETL workloads

Similar to previous blog posts about how it is possible to scale out ELT workloads with the [CSV plugin](./scaling-out-elt-with-cq-and-csv), you can do the same with SQLite while also enjoying the benefits of a relational database.

### Analysis/Transformation backend

[Data analysis backend](https://www.sqlite.org/whentouse.html): Given the fact that SQLite is just a file it is easy and performant to actually run various transform workloads from python, or any other language you prefer, as SQLite is supported everywhere. This can also serve as intermediate storage before you load data to a data warehouse for further analysis and visualization.

This user on [Hacker News](https://news.ycombinator.com/item?id=22153447) was lyrical about using SQLite for a similar ETL use case.

## Summary

With SQLite the possibilities are endless, and I'd love to hear your use-cases and ideas on how you leverage it. Feel free to open an issue on [GitHub](https://github.com/cloudquery/cloudquery) or join our [Discord](https://cloudquery.io/discord) to share your ideas.
