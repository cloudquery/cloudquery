# Table: aws_apigateway_rest_api_documentation_versions

https://docs.aws.amazon.com/apigateway/latest/api/API_DocumentationVersion.html

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
|created_date|Timestamp|
|description|String|
|version|String|