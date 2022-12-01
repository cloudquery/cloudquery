# Table: aws_apigateway_api_keys

https://docs.aws.amazon.com/apigateway/latest/api/API_ApiKey.html

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
|created_date|Timestamp|
|customer_id|String|
|description|String|
|enabled|Bool|
|id|String|
|last_updated_date|Timestamp|
|name|String|
|stage_keys|StringArray|
|tags|JSON|
|value|String|