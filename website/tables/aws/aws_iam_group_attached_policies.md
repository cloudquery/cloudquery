# Table: aws_iam_group_attached_policies

This table shows data for IAM Group Attached Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html

The composite primary key for this table is (**group_arn**, **policy_name**).

## Relations

This table depends on [aws_iam_groups](aws_iam_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|group_arn (PK)|String|
|policy_name (PK)|String|
|group_id|String|
|policy_arn|String|