# Table: aws_iam_groups

This table shows data for IAM Groups.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

The following tables depend on aws_iam_groups:
  - [aws_iam_group_attached_policies](aws_iam_group_attached_policies)
  - [aws_iam_group_last_accessed_details](aws_iam_group_last_accessed_details)
  - [aws_iam_group_policies](aws_iam_group_policies)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id (PK)|utf8|
|arn (PK)|utf8|
|create_date|timestamp[us, tz=UTC]|
|group_id|utf8|
|group_name|utf8|
|path|utf8|