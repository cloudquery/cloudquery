# Table: aws_iam_users

This table shows data for IAM Users.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html

The composite primary key for this table is (**account_id**, **arn**).

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
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id (PK)|utf8|
|arn (PK)|utf8|
|tags|json|
|create_date|timestamp[us, tz=UTC]|
|path|utf8|
|user_id|utf8|
|user_name|utf8|
|password_last_used|timestamp[us, tz=UTC]|
|permissions_boundary|json|