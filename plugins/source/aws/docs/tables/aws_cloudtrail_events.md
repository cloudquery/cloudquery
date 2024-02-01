# Table: aws_cloudtrail_events

This table shows data for AWS CloudTrail Events.

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_Event.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **event_id**.It supports incremental syncs based on the **event_time** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cloud_trail_event|`json`|
|event_time (Incremental Key)|`timestamp[us, tz=UTC]`|
|access_key_id|`utf8`|
|event_id|`utf8`|
|event_name|`utf8`|
|event_source|`utf8`|
|read_only|`utf8`|
|resources|`json`|
|username|`utf8`|