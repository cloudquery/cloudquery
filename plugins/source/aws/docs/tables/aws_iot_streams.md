# Table: aws_iot_streams

https://docs.aws.amazon.com/iot/latest/apireference/API_StreamInfo.html

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
|created_at|Timestamp|
|description|String|
|files|JSON|
|last_updated_at|Timestamp|
|role_arn|String|
|stream_id|String|
|stream_version|Int|