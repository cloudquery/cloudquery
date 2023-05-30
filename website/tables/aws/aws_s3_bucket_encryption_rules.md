# Table: aws_s3_bucket_encryption_rules

This table shows data for S3 Bucket Encryption Rules.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_ServerSideEncryptionRule.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_s3_buckets](aws_s3_buckets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|bucket_arn|`utf8`|
|apply_server_side_encryption_by_default|`json`|
|bucket_key_enabled|`bool`|