# Table: aws_s3_bucket_encryption_rules

This table shows data for S3 Bucket Encryption Rules.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_ServerSideEncryptionRule.html

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
|apply_server_side_encryption_by_default|`json`|
|bucket_key_enabled|`bool`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### S3 buckets should have server-side encryption enabled

```sql
SELECT
  'S3 buckets should have server-side encryption enabled' AS title,
  aws_s3_buckets.account_id,
  arn AS resource_id,
  CASE
  WHEN aws_s3_bucket_encryption_rules.bucket_arn IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_s3_buckets
  LEFT JOIN aws_s3_bucket_encryption_rules ON
      aws_s3_bucket_encryption_rules.bucket_arn = aws_s3_buckets.arn;
```


