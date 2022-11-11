---
title: Announcing CloudQuery SQLite Destination Plugin
tag: tutorial
date: 2022/11/11
description: >-
  This tutorial will show you how to sync your cloud resources to SQLite database.
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

# Introduction

SQLite is all the rage those days and for a good reason, it is embedded database with no dependency that can be stored as a file while having a lot of the features of a full blown relational database such as query language, indexes, data types, etc…

Today I'm excited to announce the release of our new SQLite destination plugin so you can now sync all supported cloudquery [source plugins](../docs/plugins/sources) to a [SQLite](https://www.sqlite.org/index.html) database.

Checkout [our docs](../docs/recipes/destinations/sqlite) to see how to configure the SQLite destination plugin.

# Use Cases

SQLite use-cases in data integration are infinite but I'd like to share a few interesting but popular ideas.

## Local data exploration

If you don't want to run a database and/or docker, running sqlite as destination will be super performant similar to our [CSV destination plugin](../docs/recipes/destinations/csv) but with the ability to do complex queries and store everything in one file.

## Scaling ETL workloads

Similar to previous blog posts about how it is possible to scale out ELT workloads with [CSV plugin](./scaling-out-elt-with-cq-and-csv), you can do the same with sqlite and enjoy all the benefits of a relational database.

## Analysis/Transformation backend

[Data analysis backend](https://www.sqlite.org/whentouse.html): Given the fact that sqlite is just a file it is easy and performant to actually run various transform workloads from python, or any other language you perfer as sqlite is supported everywhere. This can also serve as intermediate storage before you load them to a data warehouse for further analysis and visualization.

Found this discussion on [HN](https://news.ycombinator.com/item?id=22153447) about similar use case.

# Summary

With SQLite the possibilities are endless and I'd love to hear your use-cases and ideas on how you leverage it, feel free to open an issue on [github](https://github.com/cloudquery/cloudquery) and/or join our [discord](https://cloudquery.io/discord) to share your ideas.