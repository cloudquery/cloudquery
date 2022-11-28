# Table: aws_lambda_layer_version_policies



The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_lambda_layer_versions](aws_lambda_layer_versions.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|layer_version_arn|String|
|layer_version|Int|
|policy|String|
|revision_id|String|
|result_metadata|JSON|