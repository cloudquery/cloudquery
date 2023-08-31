# Table: aws_s3_bucket_lifecycles

This table shows data for S3 Bucket Lifecycles.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_LifecycleRule.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_s3_buckets](aws_s3_buckets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|bucket_arn|`utf8`|
|status|`utf8`|
|abort_incomplete_multipart_upload|`json`|
|expiration|`json`|
|id|`utf8`|
|noncurrent_version_expiration|`json`|
|noncurrent_version_transitions|`json`|
|prefix|`utf8`|
|transitions|`json`|