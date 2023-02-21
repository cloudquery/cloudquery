# Table: aws_iam_user_access_keys

https://docs.aws.amazon.com/IAM/latest/APIReference/API_AccessKeyMetadata.html

The composite primary key for this table is (**account_id**, **user_arn**, **access_key_id**).

## Relations

This table depends on [aws_iam_users](aws_iam_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|user_arn (PK)|String|
|access_key_id (PK)|String|
|user_id|String|
|last_used|Timestamp|
|last_used_service_name|String|
|create_date|Timestamp|
|status|String|
|user_name|String|
|last_rotated|Timestamp|