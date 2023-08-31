# Table: aws_apigatewayv2_api_integration_responses

This table shows data for Amazon API Gateway v2 API Integration Responses.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-integrations-integrationid-integrationresponses-integrationresponseid.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_api_integrations](aws_apigatewayv2_api_integrations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|api_integration_arn|`utf8`|
|integration_id|`utf8`|
|arn (PK)|`utf8`|
|integration_response_key|`utf8`|
|content_handling_strategy|`utf8`|
|integration_response_id|`utf8`|
|response_parameters|`json`|
|response_templates|`json`|
|template_selection_expression|`utf8`|