# Table: aws_iam_roles

This table shows data for IAM Roles.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Role.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

The following tables depend on aws_iam_roles:
  - [aws_iam_role_attached_policies](aws_iam_role_attached_policies)
  - [aws_iam_role_last_accessed_details](aws_iam_role_last_accessed_details)
  - [aws_iam_role_policies](aws_iam_role_policies)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|assume_role_policy_document|`json`|
|tags|`json`|
|arn (PK)|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|path|`utf8`|
|role_id|`utf8`|
|role_name|`utf8`|
|description|`utf8`|
|max_session_duration|`int64`|
|permissions_boundary|`json`|
|role_last_used|`json`|