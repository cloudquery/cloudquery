---
title: Glossary
---

# Glossary

## CLI

The CloudQuery CLI (command-line interface) is the core [open-source](https://github.com/cloudquery/cloudquery) project that expects to be run on your local machine as CLI or on a remote machine.

## SDK

The [CloudQuery SDK](https://github.com/cloudquery/cq-provider-sdk) is the open-source SDK library used by official and community plugins to integrate with the CloudQuery ecosystem.
The SDK makes it easy to write new plugins (both source and destination plugins) and takes care of the TL in ETL (extract-transform-load).

## Fetch

Fetch is both the CLI command and the process when CloudQuery extracts all the configured resources in a directory
with `*.cq.yml`, transforms them, and loads them into the destination (database).

## Policy

Policy compliance is a broad term and can refer to any kind of policy, from internal standards to regulatory requirements.
A CloudQuery Policy is a set of SQL queries used extract and transform results to a table that can later on be visualized and analyzed
by any BI tools and standard set of data engineering tools.

## Plugin

CloudQuery supports two plugins: source & destination.

`source` plugin is responsible for extracting data from remote APIs, transforming it and sending it to CloudQuery for further handling.

`destination` plugin is responsible for getting data from one the `source` plugin and save it to a database, a data lake or a subscription (Kafka) according to configuration.
