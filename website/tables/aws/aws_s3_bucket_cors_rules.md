# Table: aws_s3_bucket_cors_rules

This table shows data for S3 Bucket Cors Rules.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_CORSRule.html

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
|allowed_methods|`list<item: utf8, nullable>`|
|allowed_origins|`list<item: utf8, nullable>`|
|allowed_headers|`list<item: utf8, nullable>`|
|expose_headers|`list<item: utf8, nullable>`|
|id|`utf8`|
|max_age_seconds|`int64`|