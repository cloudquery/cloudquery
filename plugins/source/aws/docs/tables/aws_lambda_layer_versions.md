# Table: aws_lambda_layer_versions

https://docs.aws.amazon.com/lambda/latest/dg/API_LayerVersionsListItem.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_lambda_layers](aws_lambda_layers.md).

The following tables depend on aws_lambda_layer_versions:
  - [aws_lambda_layer_version_policies](aws_lambda_layer_version_policies.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn|String|
|layer_arn|String|
|compatible_architectures|StringArray|
|compatible_runtimes|StringArray|
|created_date|String|
|description|String|
|license_info|String|
|version|Int|