# Table: aws_appsync_graphql_apis

https://docs.aws.amazon.com/appsync/latest/APIReference/API_GraphqlApi.html

The primary key for this table is **arn**.



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
|additional_authentication_providers|JSON|
|api_id|String|
|authentication_type|String|
|lambda_authorizer_config|JSON|
|log_config|JSON|
|name|String|
|open_id_connect_config|JSON|
|tags|JSON|
|uris|JSON|
|user_pool_config|JSON|
|waf_web_acl_arn|String|
|xray_enabled|Bool|