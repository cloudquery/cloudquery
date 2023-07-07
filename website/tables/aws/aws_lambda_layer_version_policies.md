# Table: aws_lambda_layer_version_policies

This table shows data for AWS Lambda Layer Version Policies.

https://docs.aws.amazon.com/lambda/latest/dg/API_GetLayerVersionPolicy.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lambda_layer_versions](aws_lambda_layer_versions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|layer_version_arn|`utf8`|
|layer_version|`int64`|
|policy|`utf8`|
|revision_id|`utf8`|
|result_metadata|`json`|