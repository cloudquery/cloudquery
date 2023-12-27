# Table: aws_lambda_layer_version_policies

This table shows data for AWS Lambda Layer Version Policies.

https://docs.aws.amazon.com/lambda/latest/dg/API_GetLayerVersionPolicy.html

The composite primary key for this table is (**layer_version_arn**, **revision_id**).

## Relations

This table depends on [aws_lambda_layer_versions](aws_lambda_layer_versions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|layer_version_arn (PK)|`utf8`|
|layer_version|`int64`|
|policy|`utf8`|
|revision_id (PK)|`utf8`|