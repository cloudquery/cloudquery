# Table: aws_iam_ssh_public_keys

https://docs.aws.amazon.com/IAM/latest/APIReference/API_SSHPublicKeyMetadata.html

The primary key for this table is **ssh_public_key_id**.

## Relations
This table depends on [aws_iam_users](aws_iam_users.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|user_arn|String|
|user_id|String|
|ssh_public_key_id (PK)|String|
|status|String|
|upload_date|Timestamp|
|user_name|String|