# Table: aws_apigateway_rest_apis

https://docs.aws.amazon.com/apigateway/latest/api/API_RestApi.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_apigateway_rest_apis:
  - [aws_apigateway_rest_api_authorizers](aws_apigateway_rest_api_authorizers.md)
  - [aws_apigateway_rest_api_deployments](aws_apigateway_rest_api_deployments.md)
  - [aws_apigateway_rest_api_documentation_parts](aws_apigateway_rest_api_documentation_parts.md)
  - [aws_apigateway_rest_api_documentation_versions](aws_apigateway_rest_api_documentation_versions.md)
  - [aws_apigateway_rest_api_gateway_responses](aws_apigateway_rest_api_gateway_responses.md)
  - [aws_apigateway_rest_api_models](aws_apigateway_rest_api_models.md)
  - [aws_apigateway_rest_api_request_validators](aws_apigateway_rest_api_request_validators.md)
  - [aws_apigateway_rest_api_resources](aws_apigateway_rest_api_resources.md)
  - [aws_apigateway_rest_api_stages](aws_apigateway_rest_api_stages.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|api_key_source|String|
|binary_media_types|StringArray|
|created_date|Timestamp|
|description|String|
|disable_execute_api_endpoint|Bool|
|endpoint_configuration|JSON|
|id|String|
|minimum_compression_size|Int|
|name|String|
|policy|String|
|tags|JSON|
|version|String|
|warnings|StringArray|