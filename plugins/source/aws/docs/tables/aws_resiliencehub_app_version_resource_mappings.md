# Table: aws_resiliencehub_app_version_resource_mappings

This table shows data for AWS Resilience Hub App Version Resource Mappings.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ResourceMapping.html

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
|mapping_type|`utf8`|
|physical_resource_id|`json`|
|app_registry_app_name|`utf8`|
|eks_source_name|`utf8`|
|logical_stack_name|`utf8`|
|resource_group_name|`utf8`|
|resource_name|`utf8`|
|terraform_source_name|`utf8`|