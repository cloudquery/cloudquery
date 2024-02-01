# Table: aws_iam_groups

This table shows data for IAM Groups.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

The following tables depend on aws_iam_groups:
  - [aws_iam_group_attached_policies](aws_iam_group_attached_policies.md)
  - [aws_iam_group_last_accessed_details](aws_iam_group_last_accessed_details.md)
  - [aws_iam_group_policies](aws_iam_group_policies.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|group_id|`utf8`|
|group_name|`utf8`|
|path|`utf8`|