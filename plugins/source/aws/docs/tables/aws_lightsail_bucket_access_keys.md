# Table: aws_lightsail_bucket_access_keys

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_AccessKey.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_lightsail_buckets](aws_lightsail_buckets.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|bucket_arn|String|
|access_key_id|String|
|created_at|Timestamp|
|last_used|JSON|
|secret_access_key|String|
|status|String|