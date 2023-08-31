# Table: aws_apigateway_domain_name_base_path_mappings

This table shows data for Amazon API Gateway Domain Name Base Path Mappings.

https://docs.aws.amazon.com/apigateway/latest/api/API_BasePathMapping.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_domain_names](aws_apigateway_domain_names).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|domain_name_arn|`utf8`|
|arn (PK)|`utf8`|
|base_path|`utf8`|
|rest_api_id|`utf8`|
|stage|`utf8`|