# Table: aws_lambda_layers

https://docs.aws.amazon.com/lambda/latest/dg/API_LayersListItem.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lambda_layers:
  - [aws_lambda_layer_versions](aws_lambda_layer_versions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|latest_matching_version|JSON|
|layer_arn|String|
|layer_name|String|