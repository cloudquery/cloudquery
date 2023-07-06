# Table: aws_lightsail_bucket_access_keys

This table shows data for Lightsail Bucket Access Keys.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_AccessKey.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lightsail_buckets](aws_lightsail_buckets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|bucket_arn|`utf8`|
|access_key_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|last_used|`json`|
|secret_access_key|`utf8`|
|status|`utf8`|