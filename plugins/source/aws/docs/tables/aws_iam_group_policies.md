# Table: aws_iam_group_policies


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_iam_groups`](aws_iam_groups.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|group_arn|String|
|group_id|String|
|policy_document|JSON|
|group_name|String|
|policy_name|String|
|result_metadata|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|