# Table: aws_eventbridge_endpoints

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Endpoint.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|creation_time|Timestamp|
|description|String|
|endpoint_id|String|
|endpoint_url|String|
|event_buses|JSON|
|last_modified_time|Timestamp|
|name|String|
|replication_config|JSON|
|role_arn|String|
|routing_config|JSON|
|state|String|
|state_reason|String|