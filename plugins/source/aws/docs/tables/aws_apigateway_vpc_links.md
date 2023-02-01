# Table: aws_apigateway_vpc_links

https://docs.aws.amazon.com/apigateway/latest/api/API_VpcLink.html

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|arn (PK)|String|
|description|String|
|id|String|
|name|String|
|status|String|
|status_message|String|
|tags|JSON|
|target_arns|StringArray|