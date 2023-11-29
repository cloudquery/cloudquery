# Table: aws_s3_bucket_loggings

This table shows data for S3 Bucket Loggings.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLogging.html

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
|logging_enabled|`json`|