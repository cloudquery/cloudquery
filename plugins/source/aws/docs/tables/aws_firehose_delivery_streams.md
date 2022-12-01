# Table: aws_firehose_delivery_streams

https://docs.aws.amazon.com/firehose/latest/APIReference/API_DeliveryStreamDescription.html

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
|tags|JSON|
|arn (PK)|String|
|delivery_stream_name|String|
|delivery_stream_status|String|
|delivery_stream_type|String|
|destinations|JSON|
|has_more_destinations|Bool|
|version_id|String|
|create_timestamp|Timestamp|
|delivery_stream_encryption_configuration|JSON|
|failure_description|JSON|
|last_update_timestamp|Timestamp|
|source|JSON|