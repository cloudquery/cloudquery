# Table: aws_s3_bucket_encryption_rules

This table shows data for S3 Bucket Encryption Rules.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_ServerSideEncryptionRule.html

The primary key for this table is **bucket_arn**.

## Relations

This table depends on [aws_s3_buckets](aws_s3_buckets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|bucket_arn (PK)|`utf8`|
|apply_server_side_encryption_by_default|`json`|
|bucket_key_enabled|`bool`|