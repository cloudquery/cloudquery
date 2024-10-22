---
title: Telemetry
description: Information about telemetry data collected by the CloudQuery CLI.
---

# Telemetry

CloudQuery CLI and integrations collects both anonymous and identifiable usage data. We use the anonymous and identifiable usage data to help us improve the product, while we only use the identifiable information for billing. This page describes what data is collected and how to control it.

## Identifiable Data

As part of the commercial offering, Premium CloudQuery integrations periodically send data to the CloudQuery licensing server to validate that the user or team has a valid subscription. The data in these requests include:
- `plugin_name`: the full name of the integration (including the team that owns the integration)
- `plugin_kind`: whether the integration is a source or destination integration
- `table_name`: the name of the table being synced. This is only sent for sources whose tables are statically defined. Any source that has dynamic table names (for example all database sources) will not send this field.
- `resource_count`: the number of rows being synced

This data cannot be disabled and is required for the commercial offering to function. If you have a use case that requires this data not be sent, please contact the [Sales team](https://cloudquery.typeform.com/to/UrgOydHV?typeform-source=www.cloudquery.io) to discuss your requirements.

## Telemetry Data

By default, the CloudQuery CLI collects usage data. There are two types of data we collect: **errors** and **stats**. These are described below. The [Controlling what is sent](#controlling-what-is-sent) section describes how to control what you send to CloudQuery. 
 
### Errors

Errors are stack traces sent whenever a panic occurs in the CLI or an official integration. Having this data allows us to be notified when there is a bug that needs to be prioritized.    
 
## Stats

#### Anonymous Stats

Anonymous stats are numbers about the sync that was performed, such as the number of errors and number of resources fetched, as well as the integration versions used. These are sent at the end of a sync. They contain no identifying information. We use this data to understand which integrations are being used and how much, which helps guide our roadmap and development efforts.
We also anonymously track command that are run, and if they result in an error. This helps us understand how the CLI is being used and how many errors are being encountered.

#### Identifiable Stats

Identifiable stats are the same as anonymous stats, but they also include a unique identifier for the user or team.
Identifiable stats are collected when the CLI is authenticated via `cloudquery login` or using an [API key](/docs/deployment/generate-api-key).
We track the following user event types, `Login`, `Sync Started` and `Sync Finished`.
We use this information to understand usage patterns, which helps us improve the product and our support.

## Controlling what is sent

The CLI supports two methods of controlling the telemetry that gets sent to CloudQuery.

#### Command-line Flag

The `--telemetry-level` can be passed to the `cloudquery` CLI. It supports four options: `none`, `stats`, `errors`, `all` (default).

#### Environment Variable

A `CQ_TELEMETRY_LEVEL` environment variable can also be used to control the telemetry being sent. It supports the same options as the `--telemetry-level` flag.

## More Information

See the [`cloudquery` command reference](reference/cli/cloudquery) for all available command-line options.

