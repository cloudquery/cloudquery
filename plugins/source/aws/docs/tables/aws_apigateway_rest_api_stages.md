# Table: aws_apigateway_rest_api_stages

https://docs.aws.amazon.com/apigateway/latest/api/API_Stage.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_apigateway_rest_apis](aws_apigateway_rest_apis.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|rest_api_arn|String|
|arn|String|
|access_log_settings|JSON|
|cache_cluster_enabled|Bool|
|cache_cluster_size|String|
|cache_cluster_status|String|
|canary_settings|JSON|
|client_certificate_id|String|
|created_date|Timestamp|
|deployment_id|String|
|description|String|
|documentation_version|String|
|last_updated_date|Timestamp|
|method_settings|JSON|
|stage_name|String|
|tags|JSON|
|tracing_enabled|Bool|
|variables|JSON|
|web_acl_arn|String|