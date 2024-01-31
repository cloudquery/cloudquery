# Table: aws_iam_group_policies

This table shows data for IAM Group Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetGroupPolicy.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **group_arn**, **policy_name**).
## Relations

This table depends on [aws_iam_groups](aws_iam_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|group_arn|`utf8`|
|policy_document|`json`|
|group_name|`utf8`|
|policy_name|`utf8`|