# Table: aws_iam_group_policies

This table shows data for IAM Group Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetGroupPolicy.html

The composite primary key for this table is (**account_id**, **group_arn**, **policy_name**).

## Relations

This table depends on [aws_iam_groups](aws_iam_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|group_arn (PK)|String|
|policy_document|JSON|
|group_name|String|
|policy_name (PK)|String|
|result_metadata|JSON|