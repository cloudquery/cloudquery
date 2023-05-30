# Table: aws_s3_buckets

This table shows data for S3 Buckets.

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_s3_buckets:
  - [aws_s3_bucket_cors_rules](aws_s3_bucket_cors_rules)
  - [aws_s3_bucket_encryption_rules](aws_s3_bucket_encryption_rules)
  - [aws_s3_bucket_grants](aws_s3_bucket_grants)
  - [aws_s3_bucket_lifecycles](aws_s3_bucket_lifecycles)
  - [aws_s3_bucket_websites](aws_s3_bucket_websites)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|replication_role|`utf8`|
|replication_rules|`json`|
|region|`utf8`|
|logging_target_bucket|`utf8`|
|logging_target_prefix|`utf8`|
|policy|`json`|
|versioning_status|`utf8`|
|versioning_mfa_delete|`utf8`|
|block_public_acls|`bool`|
|block_public_policy|`bool`|
|ignore_public_acls|`bool`|
|restrict_public_buckets|`bool`|
|tags|`json`|
|ownership_controls|`list<item: utf8, nullable>`|