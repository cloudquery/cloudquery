# Table: aws_apigateway_rest_api_resource_method_integrations

https://docs.aws.amazon.com/apigateway/latest/api/API_Integration.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_rest_api_resource_methods](aws_apigateway_rest_api_resource_methods.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|rest_api_arn|String|
|resource_arn|String|
|method_arn|String|
|arn (PK)|String|
|cache_key_parameters|StringArray|
|cache_namespace|String|
|connection_id|String|
|connection_type|String|
|content_handling|String|
|credentials|String|
|http_method|String|
|integration_responses|JSON|
|passthrough_behavior|String|
|request_parameters|JSON|
|request_templates|JSON|
|timeout_in_millis|Int|
|tls_config|JSON|
|type|String|
|uri|String|