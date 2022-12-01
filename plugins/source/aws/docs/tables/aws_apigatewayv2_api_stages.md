# Table: aws_apigatewayv2_api_stages

https://docs.aws.amazon.com/apigateway/latest/api/API_Stage.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis.md).


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
|stage_name|String|
|access_log_settings|JSON|
|api_gateway_managed|Bool|
|auto_deploy|Bool|
|client_certificate_id|String|
|created_date|Timestamp|
|default_route_settings|JSON|
|deployment_id|String|
|description|String|
|last_deployment_status_message|String|
|last_updated_date|Timestamp|
|route_settings|JSON|
|stage_variables|JSON|
|tags|JSON|