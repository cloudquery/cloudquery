# Table: aws_apigateway_domain_name_base_path_mappings

https://docs.aws.amazon.com/apigateway/latest/api/API_BasePathMapping.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_apigateway_domain_names](aws_apigateway_domain_names.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|domain_name_arn|String|
|arn|String|
|base_path|String|
|rest_api_id|String|
|stage|String|