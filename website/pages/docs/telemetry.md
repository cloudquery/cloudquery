---
title: Telemetry
---

# Telemetry

By default, the CloudQuery CLI collects some anonymous usage data. There are two types of data we collect: **errors** and **stats**. These are described below. The [Controlling what is sent](#controlling-what-is-sent) section describes how to control what you send to CloudQuery. 
 
## Errors

Errors are stack traces sent whenever a panic occurs in the CLI or an official plugin. Having this data allows us to be notified when there is a bug that needs to be prioritized.    
 
## Stats

Stats data are numbers about the sync that was performed, such as the number of errors and number of resources fetched, as well as the plugin versions used. These are sent at the end of a sync. They contain no identifying information. We use this data to understand which plugins are being used and how much, which helps guide our roadmap and development efforts.  

## Controlling what is sent

The CLI supports two methods of controlling the telemetry that gets sent to CloudQuery.

### Command-line Flag

The `--telemetry-level` can be passed to the `cloudquery` CLI. It supports four options: `none`, `stats`, `errors`, `all` (default).

### Environment Variable

A `CQ_TELEMETRY_LEVEL` environment variable can also be used to control the telemetry being sent. It supports the same options as the `--telemetry-level` flag.

### More Information

See the [`cloudquery` command reference](reference/cli/cloudquery) for all available command-line options.

