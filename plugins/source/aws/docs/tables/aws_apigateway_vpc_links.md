# Table: aws_apigateway_vpc_links


The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|