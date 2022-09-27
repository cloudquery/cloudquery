---
title: CloudQuery V1 Release
tag: announcement
date: 2022/09/24
description: >-
  New exciting features, architecture changes and roadmap!
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

Today is an exciting day for us at CloudQuery. For the last three months, we have been working on a completely new version of CloudQuery: with all the learnings and feedback from our amazing users over the past year. And today we are announcing its release! (We are zero-based, so it's v1 even though it's the second version 🙂) In this post, I will walk through the major features, architecture changes and upcoming work.

Quick Recap: If you are unfamiliar with CloudQuery and why we are bringing the best of data engineering to infrastructure and security teams - check out our post from the recent series A [announcement](https://www.cloudquery.io/blog/cloudquery-raises-15m-series-a#story-time).

And now to the new features!

## New Destinations

Data integration is a scale problem of developing and maintaining quality plugins for huge number of APIs. This problem grows exponentially if you want also to support multiple destinations such as databases, data-warehouses, datalakes, message queues and storage systems.

To be able to support a growing number of destinations, we decoupled the source plugins from the destination layer.

Now, any CloudQuery source plugins will get support of new destinations plugins out-of-the-box without any need to update source plugins.

Destinations are also designed in the same pluggable (gRPC) way, so it will be easy to develop community plugins and official plugins separately without bloating CloudQuery CLI.

Last, but not least, with the upcoming new destinations we also support two new modes of operation: `overwrite` and `append-only`. Where previously we only supported `overwrite` mode, now you can achiveve history-like capabilities for compliance and other use cases in conjuction with new destinations for data-lakes and data-warehouses where storage is cheap.

## Improved SDK

As data-integration is a scale problem we've put a lot of effort to provide an SDK which will signfictaly reduce development time for both official and community source plugins. Here are some of the features:

- **Built-in Concurrency** - High performance is critical--especially for infrastructure tasks—-this is why we take advantage of the excellent concurrency support in Go. We do smart scheduling for the source plugin requests, so plugins can extract information in parallel from multiple APIs as well as multiple accounts. This is a very common use case, especially when syncing from the big cloud providers.
- **Structured logging** - ELT depends on third party APIs which can have down time, permissions errors, throttling errors and more. The key to running such pipelines successfully is monitoring and understanding errors. The new SDK provides structured logging with fast performance from the zerolog library, so it’s easy to understand errors, counts, timings, etc.
- **Cross Platform / Cross Language support** - Plugins are now communicating via simple gRPC protocol which means plugins can be compiled to any OS and architecture and can be written in any language (although we are currently only supporting Go). Previously they also worked over via gRPC but we used HashiCorp go-plugin abstraction on top of gRPC which turned out to be not a great fit for us and unnecessary size and documentation bloat, the new dependency light CLI size was reduced from 40MB->21MB and made development and debugging a breath.

## Code Generation

To speed up development and scale maintainability we invested much in code generation which we also open-sourced as part of CloudQuery [plugin-sdk](https://github.com/cloudquery/plugin-sdk/tree/main/codegen).

The new SDK codegen capabilities provides the following advantages for source plugins that use Go clients to communicate with the APIs:

- Autogenerating CloudQuery tables from Go structs.
- Auto-detect underlying API changes and re-generate CloudQuery tables.

A good example is our [GCP](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/gcp/codegenmain/main.go) plugin which is 95% auto-generated.

## Auto Migrations

As data integration platform users build their own views and queries on top so we want to make the maximum effort not to create backward incompatible changes to schemas when we release new features for our source plugins.

For this New official destination plugins support [auto migrations](https://v1.cloudquery.io/docs/core-concepts/migrations) and [release stages](https://v1.cloudquery.io/docs/plugins/source_plugins_release_stages)

## Monorepo

As we were looking to expand our source plugins we quickly noticed a few issues in discoverability, issue management, and DevOps and development duplication across growing >20 repos.

We have now two main repos:

- [github.com/cloudquery/cloudquery](https://github.com/cloudquery/cloudquery): CloudQuery main repo containing the CLI, official source and destination plugins.
- [github.com/cloudquery/plugin-sdk](https://github.com/cloudquery/plugin-sdk): SDK for source and destination plugins.

## What's coming up next

- Over the next few weeks we will work to fully stabilize the new release. After that, we will double down on new destination plugins, write modes and new source plugins with code generation.
- New destinations: BigQuery, Snowflake, Kafka coming up! Feel free to open an issue if you have other requests!

## Thanks!

Huge thanks for the whole CloudQuery team that put day and night in the last 3 months to re-write CloudQuery from scratch. And big thanks to every single user that used/is using CloudQuery and gave great feedback and was patient with us while we fix, re-write and incorporate the feedback!
