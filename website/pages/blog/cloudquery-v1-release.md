---
title: CloudQuery V1 Release
tag: announcement
date: 2022/10/04
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

Data integration is a scale problem of developing and maintaining quality plugins for huge number of APIs. This problem grows exponentially if you want also to support multiple destinations such as databases, data-warehouses, data lakes, message queues and storage systems.

To be able to support a growing number of destinations, we decoupled the source plugins from the destination layer.

Now, any CloudQuery source plugins will get support of new destinations plugins out-of-the-box without any need to update source plugins.

Destinations are also designed in the same pluggable (gRPC) way, so it will be easy to develop community plugins and official plugins separately without bloating CloudQuery CLI.

Last, but not least, with the upcoming new destinations we also support three new modes of operation: `overwrite`, `overwrite-delete-stale` and `append-only`. Where previously we only supported `overwrite-delete-stale` mode, now you can achieve history-like capabilities for compliance and other use cases in conjunction with new destinations for data-lakes and data-warehouses where storage is cheap.

## Improved SDK

As data-integration is a scale problem we've put a lot of effort to provide an SDK which will significantly reduce development time for both official and community source plugins. Here are some of the features:

- **Built-in Concurrency** - High performance is critical--especially for infrastructure tasks—-this is why we take advantage of the excellent concurrency support in Go. We do smart scheduling for the source plugin requests, so plugins can extract information in parallel from multiple APIs as well as multiple accounts. This is a common use case, especially when syncing from the big cloud providers.
- **Structured logging** - ELT depends on third party APIs which can have down time, permissions errors, throttling errors and more. The key to running such pipelines successfully is monitoring and understanding errors. The new SDK provides structured logging with fast performance from the `zerolog` library, so it’s easy to understand errors, counts, timings, etc.
- **Cross Platform / Cross Language support** - Plugins are now communicating via simple gRPC protocol, which means plugins can be compiled to any OS and architecture and can be written in any language (although we are currently only supporting Go). Previously they also streamed over gRPC, but we used HashiCorp's go-plugin abstraction on top of gRPC, which turned out not to be a great fit for us, causing unnecessary size and documentation bloat. The new dependency-light CLI size is now 21MB, down from 40MB, and makes development and debugging a breeze.

## Code Generation

To speed up development and scale maintainability we invested much in code generation which we also open-sourced as part of CloudQuery [plugin-sdk](https://github.com/cloudquery/plugin-sdk/tree/main/codegen).

The new SDK code generation capabilities provide the following advantages for source plugins that use Go clients to communicate with the APIs:

- Automatically generating CloudQuery tables from Go structs.
- Auto-detect underlying API changes and re-generate CloudQuery tables.

A good example is our [GCP](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/gcp/codegen/main.go) plugin which is now 95% auto-generated.

## Auto Migrations

As data integration platform users build their own views and queries on top, we made the maximum effort not to create backward incompatible changes to schemas when we release new features for our source plugins.

For this New official destination plugins support [auto migrations](/docs/core-concepts/migrations) and [release stages](/docs/plugins/sources/overview#source-plugin-release-stages)

## Monorepo

As we were looking to expand our source plugins we quickly noticed a few issues in discoverability, issue management, DevOps and development duplication across a growing number of more than 20 repositories.

We have now two main repositories:

- [github.com/cloudquery/cloudquery](https://github.com/cloudquery/cloudquery): CloudQuery main repository containing the CLI, official source and destination plugins.
- [github.com/cloudquery/plugin-sdk](https://github.com/cloudquery/plugin-sdk): SDK for source and destination plugins.

## Policies

As we are a big believer in data-engineering best practices we provide a set of standard [SQL queries](https://www.cloudquery.io/docs/core-concepts/policies) for popular benchmarks. Most importantly, they are all open source and not abstracted behind any custom policy language so you can re-use them, customize and apply to your needs while enjoying the SQL eco-system.

## What's coming up next

- Over the next few weeks we will work to fully stabilize the new release. After that, we will double down on new destination plugins, write modes and new source plugins with code generation.
- New destinations: BigQuery, Snowflake, Kafka coming up! Feel free to open an issue if you have other requests!

## Migrating from v0 to v1

With the release of v1 comes a promise of increased stability. However, if you were using v0 before, there will be some breaking changes that need to be handled. See our [V1 Migration guide](/blog/migrating-from-v0-to-v1) for guidance. And if you have any questions, we are always happy to help on [Discord](https://www.cloudquery.io/discord)!

## Thanks!

Huge thanks for the whole CloudQuery team that put day and night in the last 3 months to re-write CloudQuery from scratch. And big thanks to every single user that used/is using CloudQuery and gave great feedback and was patient with us while we fix, re-write and incorporate the feedback!
