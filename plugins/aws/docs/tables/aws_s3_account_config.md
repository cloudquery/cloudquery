
# Table: aws_s3_account_config
Account configurations for S3
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|config_exists|boolean|Specifies whether Amazon S3 public access control config exists|
|block_public_acls|boolean|Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account|
|block_public_policy|boolean|Specifies whether Amazon S3 should block public bucket policies for buckets in this account.|
|ignore_public_acls|boolean|Specifies whether Amazon S3 should ignore public ACLs for buckets in this account|
|restrict_public_buckets|boolean|Specifies whether Amazon S3 should restrict public bucket policies for buckets in this account.|
