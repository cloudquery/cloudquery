# Table: aws_apigatewayv2_domain_names

This table shows data for Amazon API Gateway v2 Domain Names.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/domainnames.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

The following tables depend on aws_apigatewayv2_domain_names:
  - [aws_apigatewayv2_domain_name_rest_api_mappings](aws_apigatewayv2_domain_name_rest_api_mappings)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|domain_name|`utf8`|
|api_mapping_selection_expression|`utf8`|
|domain_name_configurations|`json`|
|mutual_tls_authentication|`json`|
|tags|`json`|