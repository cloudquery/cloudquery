# Table: aws_cloudtrail_events

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_Event.html

The primary key for this table is **event_id**.
It supports incremental syncs based on the **event_time** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|cloud_trail_event|JSON|
|event_time (Incremental Key)|Timestamp|
|access_key_id|String|
|event_id (PK)|String|
|event_name|String|
|event_source|String|
|read_only|String|
|resources|JSON|
|username|String|