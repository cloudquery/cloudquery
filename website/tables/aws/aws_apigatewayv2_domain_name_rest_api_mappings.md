# Table: aws_apigatewayv2_domain_name_rest_api_mappings

This table shows data for Amazon API Gateway v2 Domain Name Rest API Mappings.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/domainnames-domainname-apimappings.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_domain_names](aws_apigatewayv2_domain_names).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|domain_name_arn|`utf8`|
|arn (PK)|`utf8`|
|api_id|`utf8`|
|stage|`utf8`|
|api_mapping_id|`utf8`|
|api_mapping_key|`utf8`|