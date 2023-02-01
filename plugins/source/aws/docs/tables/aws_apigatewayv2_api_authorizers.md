# Table: aws_apigatewayv2_api_authorizers

https://docs.aws.amazon.com/apigateway/latest/api/API_Authorizer.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|api_arn|String|
|api_id|String|
|arn (PK)|String|
|name|String|
|authorizer_credentials_arn|String|
|authorizer_id|String|
|authorizer_payload_format_version|String|
|authorizer_result_ttl_in_seconds|Int|
|authorizer_type|String|
|authorizer_uri|String|
|enable_simple_responses|Bool|
|identity_source|StringArray|
|identity_validation_expression|String|
|jwt_configuration|JSON|