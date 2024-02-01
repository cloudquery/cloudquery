# Table: aws_iam_users

This table shows data for IAM Users.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

The following tables depend on aws_iam_users:
  - [aws_iam_mfa_devices](aws_iam_mfa_devices.md)
  - [aws_iam_signing_certificates](aws_iam_signing_certificates.md)
  - [aws_iam_ssh_public_keys](aws_iam_ssh_public_keys.md)
  - [aws_iam_user_access_keys](aws_iam_user_access_keys.md)
  - [aws_iam_user_attached_policies](aws_iam_user_attached_policies.md)
  - [aws_iam_user_groups](aws_iam_user_groups.md)
  - [aws_iam_user_last_accessed_details](aws_iam_user_last_accessed_details.md)
  - [aws_iam_user_policies](aws_iam_user_policies.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|tags|`json`|
|create_date|`timestamp[us, tz=UTC]`|
|path|`utf8`|
|user_id|`utf8`|
|user_name|`utf8`|
|password_last_used|`timestamp[us, tz=UTC]`|
|permissions_boundary|`json`|