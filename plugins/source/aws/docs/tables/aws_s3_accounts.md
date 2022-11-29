# Table: aws_s3_accounts



The primary key for this table is **account_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|block_public_acls|Bool|
|block_public_policy|Bool|
|ignore_public_acls|Bool|
|restrict_public_buckets|Bool|
|config_exists|Bool|