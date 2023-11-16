# Table: aws_resiliencehub_app_versions

This table shows data for AWS Resilience Hub App Versions.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppVersionSummary.html

The composite primary key for this table is (**app_arn**, **app_version**).

## Relations

This table depends on [aws_resiliencehub_apps](aws_resiliencehub_apps.md).

The following tables depend on aws_resiliencehub_app_versions:
  - [aws_resiliencehub_app_version_resource_mappings](aws_resiliencehub_app_version_resource_mappings.md)
  - [aws_resiliencehub_app_version_resources](aws_resiliencehub_app_version_resources.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|app_arn (PK)|`utf8`|
|app_version (PK)|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|identifier|`int64`|
|version_name|`utf8`|