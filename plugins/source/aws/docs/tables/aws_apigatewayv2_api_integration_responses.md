# Table: aws_apigatewayv2_api_integration_responses

This table shows data for Amazon API Gateway v2 API Integration Responses.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-integrations-integrationid-integrationresponses-integrationresponseid.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

This table depends on [aws_apigatewayv2_api_integrations](aws_apigatewayv2_api_integrations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|api_integration_arn|`utf8`|
|integration_id|`utf8`|
|arn|`utf8`|
|integration_response_key|`utf8`|
|content_handling_strategy|`utf8`|
|integration_response_id|`utf8`|
|response_parameters|`json`|
|response_templates|`json`|
|template_selection_expression|`utf8`|