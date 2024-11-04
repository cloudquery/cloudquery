---
title: Syncs
description: CloudQuery syncs data from cloud providers to destinations. This page describes the various modes.
---

# Syncs

When you run `cloudquery sync <config>`, the CloudQuery CLI fetches data from all the source integrations matched by the config and delivers it to the matched destination integrations. This might mean fetching data from AWS, GCP and Azure and delivering it to PostgreSQL, or it could mean fetching data from AWS, Cloudflare and Datadog and delivering it to BigQuery, Kafka and Neo4j. It all depends on the configuration provided, and there is a near-endless array of possible combinations that grows every time a new source or destination is created. (Configuration is described in the [Configuration section](/docs/core-concepts/configuration).)      

Data is synced to the destination in a streaming fashion. This means that as soon as data is received for a source integration resource, it is delivered to the destination integration. Destination integrations may batch writes for performance reasons, but generally data will be delivered to the destination as the sync progresses.

## Table Sync Modes

Table syncs come in two flavors: `full` and `incremental`. A single `cloudquery sync` command invocation can combine both these types, and which type is used for a particular table depends on the table definition.

## Full Table Syncs

This is the normal mode of operation for most tables. For tables in this mode, a snapshot of all data is fetched from the corresponding APIs on every sync. Depending on the destination write mode, the data is then appended (`write_mode`: `append`), overwritten while keeping stale rows from previous syncs (`write_mode`: `overwrite`) or overwritten and rows from previous syncs deleted at the end of the sync (`write_mode`: `overwrite-delete-stale`). 

## Incremental Table Syncs

Some APIs lend themselves to being synced incrementally. Rather than fetch all past data on every sync, an incremental table will only fetch data that has changed since the last sync. This is done by storing some metadata in a state **backend**. The metadata is known as a **cursor**, and it marks where the last sync ended, so that the next sync can resume from the same point. Incremental syncs can be vastly more efficient than full syncs, especially for tables with large amounts of data. This is because only the data that's changed since the last sync needs to be retrieved, and in many cases this is a small subset of the overall dataset.

Incremental tables are always clearly marked as "incremental" in integration table documentation, along with an indication of which columns are used for the value of the cursor. Because they use state, incremental tables require a little more management. For more details, see [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables) under the Advanced Topics section.  
