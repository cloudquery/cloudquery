# Table: aws_apigateway_rest_api_stages

This table shows data for Amazon API Gateway Rest API Stages.

https://docs.aws.amazon.com/apigateway/latest/api/API_Stage.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

This table depends on [aws_apigateway_rest_apis](aws_apigateway_rest_apis.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|rest_api_arn|`utf8`|
|arn|`utf8`|
|access_log_settings|`json`|
|cache_cluster_enabled|`bool`|
|cache_cluster_size|`utf8`|
|cache_cluster_status|`utf8`|
|canary_settings|`json`|
|client_certificate_id|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|deployment_id|`utf8`|
|description|`utf8`|
|documentation_version|`utf8`|
|last_updated_date|`timestamp[us, tz=UTC]`|
|method_settings|`json`|
|stage_name|`utf8`|
|tags|`json`|
|tracing_enabled|`bool`|
|variables|`json`|
|web_acl_arn|`utf8`|