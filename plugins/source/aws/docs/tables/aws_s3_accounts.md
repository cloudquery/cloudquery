# Table: aws_s3_accounts

This table shows data for S3 Accounts.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PublicAccessBlockConfiguration.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|block_public_acls|`bool`|
|block_public_policy|`bool`|
|ignore_public_acls|`bool`|
|restrict_public_buckets|`bool`|
|config_exists|`bool`|