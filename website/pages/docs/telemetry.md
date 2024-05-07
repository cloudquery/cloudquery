---
title: Telemetry
description: Information about telemetry data collected by the CloudQuery CLI.
---

CloudQuery CLI and Plugins collects both anonymous and identifiable usage data. We use the anonymous and identifiable usage data to help us improve the product, while we only use the identifiable information for billing. This page describes what data is collected and how to control it.

# Identifiable Data

As part of the commercial offering, Premium CloudQuery plugins periodically send data to the CloudQuery licensing server to validate that the user or team has a valid subscription. The data in these requests include:
- **plugin_name**: the full name of the plugin (including the team that owns the plugin)
- **plugin_kind**: whether the plugin is a source or destination plugin
- **table_name**: the name of the table being synced. This is only sent for sources whose tables are statically defined. Any source that has dynamic table names (for example all database sources) will not send this field.
- **resource_count**: the number of rows being synced

This data cannot be disabled and is required for the commercial offering to function. If you have a use case that requires this data not be sent, please contact the [Sales team](https://cloudquery.typeform.com/to/UrgOydHV?typeform-source=www.cloudquery.io) to discuss your requirements.

# Telemetry Data

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

