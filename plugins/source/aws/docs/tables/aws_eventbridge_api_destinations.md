# Table: aws_eventbridge_api_destinations

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ApiDestination.html

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
|api_destination_state|String|
|connection_arn|String|
|creation_time|Timestamp|
|http_method|String|
|invocation_endpoint|String|
|invocation_rate_limit_per_second|Int|
|last_modified_time|Timestamp|
|name|String|