# Table: aws_apigatewayv2_api_integration_responses

https://docs.aws.amazon.com/apigateway/latest/api/API_IntegrationResponse.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_api_integrations](aws_apigatewayv2_api_integrations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|api_integration_arn|String|
|integration_id|String|
|arn (PK)|String|
|integration_response_key|String|
|content_handling_strategy|String|
|integration_response_id|String|
|response_parameters|JSON|
|response_templates|JSON|
|template_selection_expression|String|