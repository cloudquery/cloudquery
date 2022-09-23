# Table: aws_iam_groups


The composite primary key for this table is (**account_id**, **id**).

## Relations
The following tables depend on `aws_iam_groups`:
  - [`aws_iam_group_policies`](aws_iam_group_policies.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id (PK)|String|
|policies|JSON|
|id (PK)|String|
|arn|String|
|create_date|Timestamp|
|group_name|String|
|path|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|