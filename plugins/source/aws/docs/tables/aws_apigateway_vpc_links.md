# Table: aws_apigateway_vpc_links

https://docs.aws.amazon.com/apigateway/latest/api/API_VpcLink.html

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn|String|
|description|String|
|id|String|
|name|String|
|status|String|
|status_message|String|
|tags|JSON|
|target_arns|StringArray|