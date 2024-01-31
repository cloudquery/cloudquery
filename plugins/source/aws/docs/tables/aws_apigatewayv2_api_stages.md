# Table: aws_apigatewayv2_api_stages

This table shows data for Amazon API Gateway v2 API Stages.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-stages.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|api_arn|`utf8`|
|api_id|`utf8`|
|arn|`utf8`|
|stage_name|`utf8`|
|access_log_settings|`json`|
|api_gateway_managed|`bool`|
|auto_deploy|`bool`|
|client_certificate_id|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|default_route_settings|`json`|
|deployment_id|`utf8`|
|description|`utf8`|
|last_deployment_status_message|`utf8`|
|last_updated_date|`timestamp[us, tz=UTC]`|
|route_settings|`json`|
|stage_variables|`json`|
|tags|`json`|