# Table: aws_s3_bucket_encryption_rules

https://docs.aws.amazon.com/AmazonS3/latest/API/API_ServerSideEncryptionRule.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_s3_buckets](aws_s3_buckets.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|bucket_arn|String|
|apply_server_side_encryption_by_default|JSON|
|bucket_key_enabled|Bool|