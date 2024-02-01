# Table: aws_s3_buckets

This table shows data for S3 Buckets.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_s3_buckets:
  - [aws_s3_bucket_cors_rules](aws_s3_bucket_cors_rules.md)
  - [aws_s3_bucket_encryption_rules](aws_s3_bucket_encryption_rules.md)
  - [aws_s3_bucket_grants](aws_s3_bucket_grants.md)
  - [aws_s3_bucket_lifecycles](aws_s3_bucket_lifecycles.md)
  - [aws_s3_bucket_loggings](aws_s3_bucket_loggings.md)
  - [aws_s3_bucket_notification_configurations](aws_s3_bucket_notification_configurations.md)
  - [aws_s3_bucket_object_lock_configurations](aws_s3_bucket_object_lock_configurations.md)
  - [aws_s3_bucket_ownership_controls](aws_s3_bucket_ownership_controls.md)
  - [aws_s3_bucket_policies](aws_s3_bucket_policies.md)
  - [aws_s3_bucket_public_access_blocks](aws_s3_bucket_public_access_blocks.md)
  - [aws_s3_bucket_replications](aws_s3_bucket_replications.md)
  - [aws_s3_bucket_versionings](aws_s3_bucket_versionings.md)
  - [aws_s3_bucket_websites](aws_s3_bucket_websites.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|region|`utf8`|
|policy_status|`json`|
|tags|`json`|