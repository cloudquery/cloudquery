# Table: aws_lambda_layers


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_lambda_layers`:
  - [`aws_lambda_layer_versions`](aws_lambda_layer_versions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|latest_matching_version|JSON|
|layer_arn|String|
|layer_name|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|