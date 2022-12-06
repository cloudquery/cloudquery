# Table: aws_eventbridge_archives

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Archive.html

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
|archive_name|String|
|creation_time|Timestamp|
|event_count|Int|
|event_source_arn|String|
|retention_days|Int|
|size_bytes|Int|
|state|String|
|state_reason|String|