# Table: aws_s3_bucket_public_access_blocks

This table shows data for S3 Bucket Public Access Blocks.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_s3_buckets](aws_s3_buckets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|bucket_arn|`utf8`|
|public_access_block_configuration|`json`|