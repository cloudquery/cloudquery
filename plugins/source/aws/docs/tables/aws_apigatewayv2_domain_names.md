# Table: aws_apigatewayv2_domain_names

This table shows data for Amazon API Gateway v2 Domain Names.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/domainnames.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

The following tables depend on aws_apigatewayv2_domain_names:
  - [aws_apigatewayv2_domain_name_rest_api_mappings](aws_apigatewayv2_domain_name_rest_api_mappings.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|domain_name|`utf8`|
|api_mapping_selection_expression|`utf8`|
|domain_name_configurations|`json`|
|mutual_tls_authentication|`json`|
|tags|`json`|