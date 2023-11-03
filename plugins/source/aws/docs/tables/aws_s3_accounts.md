# Table: aws_s3_accounts

This table shows data for S3 Accounts.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PublicAccessBlockConfiguration.html

The primary key for this table is **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|block_public_acls|`bool`|
|block_public_policy|`bool`|
|ignore_public_acls|`bool`|
|restrict_public_buckets|`bool`|
|config_exists|`bool`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### S3 Block Public Access setting should be enabled

```sql
SELECT
  'S3 Block Public Access setting should be enabled' AS title,
  aws_iam_accounts.account_id,
  aws_iam_accounts.account_id AS resource_id,
  CASE
  WHEN config_exists IS NOT true
  OR block_public_acls IS NOT true
  OR block_public_policy IS NOT true
  OR ignore_public_acls IS NOT true
  OR restrict_public_buckets IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_accounts
  LEFT JOIN aws_s3_accounts ON
      aws_iam_accounts.account_id = aws_s3_accounts.account_id;
```


