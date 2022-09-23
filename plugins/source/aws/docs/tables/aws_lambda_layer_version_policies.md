# Table: aws_lambda_layer_version_policies


The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|layer_version_arn|String|
|layer_version|Int|
|policy|String|
|revision_id|String|
|result_metadata|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|