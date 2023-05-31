# Table: aws_appsync_graphql_apis

This table shows data for Appsync Graphql APIs.

https://docs.aws.amazon.com/appsync/latest/APIReference/API_GraphqlApi.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|additional_authentication_providers|`json`|
|api_id|`utf8`|
|authentication_type|`utf8`|
|lambda_authorizer_config|`json`|
|log_config|`json`|
|name|`utf8`|
|open_id_connect_config|`json`|
|tags|`json`|
|uris|`json`|
|user_pool_config|`json`|
|waf_web_acl_arn|`utf8`|
|xray_enabled|`bool`|