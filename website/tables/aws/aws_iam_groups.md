# Table: aws_iam_groups

This table shows data for AWS IAM Groups.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html

The composite primary key for this table is (**account_id**, **id**).

## Relations

The following tables depend on aws_iam_groups:
  - [aws_iam_group_last_accessed_details](aws_iam_group_last_accessed_details)
  - [aws_iam_group_policies](aws_iam_group_policies)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|policies|JSON|
|id (PK)|String|
|arn|String|
|create_date|Timestamp|
|group_id|String|
|group_name|String|
|path|String|