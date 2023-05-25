# Table: aws_s3_accounts

This table shows data for S3 Accounts.

The primary key for this table is **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id (PK)|utf8|
|block_public_acls|bool|
|block_public_policy|bool|
|ignore_public_acls|bool|
|restrict_public_buckets|bool|
|config_exists|bool|