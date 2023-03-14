# Table: aws_resiliencehub_app_versions

This table shows data for AWS Resiliencehub App Versions.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppVersionSummary.html

The composite primary key for this table is (**app_arn**, **app_version**).

## Relations

This table depends on [aws_resiliencehub_apps](aws_resiliencehub_apps).

The following tables depend on aws_resiliencehub_app_versions:
  - [aws_resiliencehub_app_version_resource_mappings](aws_resiliencehub_app_version_resource_mappings)
  - [aws_resiliencehub_app_version_resources](aws_resiliencehub_app_version_resources)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|app_arn (PK)|String|
|app_version (PK)|String|