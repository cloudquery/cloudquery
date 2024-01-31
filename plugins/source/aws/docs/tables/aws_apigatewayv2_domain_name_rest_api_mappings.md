# Table: aws_apigatewayv2_domain_name_rest_api_mappings

This table shows data for Amazon API Gateway v2 Domain Name Rest API Mappings.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/domainnames-domainname-apimappings.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

This table depends on [aws_apigatewayv2_domain_names](aws_apigatewayv2_domain_names.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|domain_name_arn|`utf8`|
|arn|`utf8`|
|api_id|`utf8`|
|stage|`utf8`|
|api_mapping_id|`utf8`|
|api_mapping_key|`utf8`|