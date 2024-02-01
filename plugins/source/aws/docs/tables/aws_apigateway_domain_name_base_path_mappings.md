# Table: aws_apigateway_domain_name_base_path_mappings

This table shows data for Amazon API Gateway Domain Name Base Path Mappings.

https://docs.aws.amazon.com/apigateway/latest/api/API_BasePathMapping.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

This table depends on [aws_apigateway_domain_names](aws_apigateway_domain_names.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|domain_name_arn|`utf8`|
|arn|`utf8`|
|base_path|`utf8`|
|rest_api_id|`utf8`|
|stage|`utf8`|