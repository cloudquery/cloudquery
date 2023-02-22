# Table: aws_apigateway_rest_api_resource_methods

https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_rest_api_resources](aws_apigateway_rest_api_resources.md).

The following tables depend on aws_apigateway_rest_api_resource_methods:
  - [aws_apigateway_rest_api_resource_method_integrations](aws_apigateway_rest_api_resource_method_integrations.md)

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
|arn (PK)|String|
|api_key_required|Bool|
|authorization_scopes|StringArray|
|authorization_type|String|
|authorizer_id|String|
|http_method|String|
|method_integration|JSON|
|method_responses|JSON|
|operation_name|String|
|request_models|JSON|
|request_parameters|JSON|
|request_validator_id|String|