# Table: aws_iam_users

This table shows data for IAM Users.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html

The composite primary key for this table is (**account_id**, **id**).

## Relations

The following tables depend on aws_iam_users:
  - [aws_iam_signing_certificates](aws_iam_signing_certificates)
  - [aws_iam_ssh_public_keys](aws_iam_ssh_public_keys)
  - [aws_iam_user_access_keys](aws_iam_user_access_keys)
  - [aws_iam_user_attached_policies](aws_iam_user_attached_policies)
  - [aws_iam_user_groups](aws_iam_user_groups)
  - [aws_iam_user_last_accessed_details](aws_iam_user_last_accessed_details)
  - [aws_iam_user_policies](aws_iam_user_policies)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|id (PK)|String|
|tags|JSON|
|arn|String|
|create_date|Timestamp|
|path|String|
|user_id|String|
|user_name|String|
|password_last_used|Timestamp|
|permissions_boundary|JSON|