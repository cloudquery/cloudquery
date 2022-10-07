# Table: aws_apigateway_rest_api_authorizers



The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_apigateway_rest_apis`](aws_apigateway_rest_apis.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|rest_api_arn|String|
|arn|String|
|auth_type|String|
|authorizer_credentials|String|
|authorizer_result_ttl_in_seconds|Int|
|authorizer_uri|String|
|id|String|
|identity_source|String|
|identity_validation_expression|String|
|name|String|
|provider_ar_ns|StringArray|
|type|String|