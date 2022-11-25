# Table: aws_apigatewayv2_api_integrations

https://docs.aws.amazon.com/apigateway/latest/api/API_Integration.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis.md).

The following tables depend on aws_apigatewayv2_api_integrations:
  - [aws_apigatewayv2_api_integration_responses](aws_apigatewayv2_api_integration_responses.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|api_arn|String|
|api_id|String|
|arn|String|
|api_gateway_managed|Bool|
|connection_id|String|
|connection_type|String|
|content_handling_strategy|String|
|credentials_arn|String|
|description|String|
|integration_id|String|
|integration_method|String|
|integration_response_selection_expression|String|
|integration_subtype|String|
|integration_type|String|
|integration_uri|String|
|passthrough_behavior|String|
|payload_format_version|String|
|request_parameters|JSON|
|request_templates|JSON|
|response_parameters|JSON|
|template_selection_expression|String|
|timeout_in_millis|Int|
|tls_config|JSON|