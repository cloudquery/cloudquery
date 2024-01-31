# Table: aws_lambda_layers

This table shows data for AWS Lambda Layers.

https://docs.aws.amazon.com/lambda/latest/dg/API_LayersListItem.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_lambda_layers:
  - [aws_lambda_layer_versions](aws_lambda_layer_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|latest_matching_version|`json`|
|layer_arn|`utf8`|
|layer_name|`utf8`|