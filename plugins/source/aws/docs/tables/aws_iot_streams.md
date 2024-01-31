# Table: aws_iot_streams

This table shows data for AWS IoT Streams.

https://docs.aws.amazon.com/iot/latest/apireference/API_StreamInfo.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|files|`json`|
|last_updated_at|`timestamp[us, tz=UTC]`|
|role_arn|`utf8`|
|stream_arn|`utf8`|
|stream_id|`utf8`|
|stream_version|`int64`|