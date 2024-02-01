# Table: aws_resiliencehub_app_version_resources

This table shows data for AWS Resilience Hub App Version Resources.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_PhysicalResource.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**app_arn**, **app_version**, **physical_resource_identifier**).
## Relations

This table depends on [aws_resiliencehub_app_versions](aws_resiliencehub_app_versions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|app_arn|`utf8`|
|app_version|`utf8`|
|physical_resource_identifier|`utf8`|
|logical_resource_id|`json`|
|physical_resource_id|`json`|
|resource_type|`utf8`|
|additional_info|`json`|
|app_components|`json`|
|excluded|`bool`|
|parent_resource_name|`utf8`|
|resource_name|`utf8`|
|source_type|`utf8`|