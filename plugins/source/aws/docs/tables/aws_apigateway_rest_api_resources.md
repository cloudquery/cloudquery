# Table: aws_apigateway_rest_api_resources

https://docs.aws.amazon.com/apigateway/latest/api/API_Resource.html

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
|id|String|
|parent_id|String|
|path|String|
|path_part|String|
|resource_methods|JSON|