# Table: aws_s3_buckets



The primary key for this table is **arn**.

## Relations

The following tables depend on aws_s3_buckets:
  - [aws_s3_bucket_encryption_rules](aws_s3_bucket_encryption_rules.md)
  - [aws_s3_bucket_lifecycles](aws_s3_bucket_lifecycles.md)
  - [aws_s3_bucket_grants](aws_s3_bucket_grants.md)
  - [aws_s3_bucket_cors_rules](aws_s3_bucket_cors_rules.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|creation_date|Timestamp|
|name|String|
|replication_role|String|
|replication_rules|JSON|
|region|String|
|logging_target_bucket|String|
|logging_target_prefix|String|
|policy|JSON|
|versioning_status|String|
|versioning_mfa_delete|String|
|block_public_acls|Bool|
|block_public_policy|Bool|
|ignore_public_acls|Bool|
|restrict_public_buckets|Bool|
|tags|JSON|
|ownership_controls|StringArray|