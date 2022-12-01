# Table: aws_eventbridge_connections

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Connection.html

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
|authorization_type|String|
|connection_state|String|
|creation_time|Timestamp|
|last_authorized_time|Timestamp|
|last_modified_time|Timestamp|
|name|String|
|state_reason|String|