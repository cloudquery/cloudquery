# Table: aws_firehose_delivery_streams

This table shows data for Firehose Delivery Streams.

https://docs.aws.amazon.com/firehose/latest/APIReference/API_DeliveryStreamDescription.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|delivery_stream_arn|`utf8`|
|delivery_stream_name|`utf8`|
|delivery_stream_status|`utf8`|
|delivery_stream_type|`utf8`|
|destinations|`json`|
|has_more_destinations|`bool`|
|version_id|`utf8`|
|create_timestamp|`timestamp[us, tz=UTC]`|
|delivery_stream_encryption_configuration|`json`|
|failure_description|`json`|
|last_update_timestamp|`timestamp[us, tz=UTC]`|
|source|`json`|