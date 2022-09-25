# Table: aws_iam_user_attached_policies


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_iam_users`](aws_iam_users.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|user_arn|String|
|user_id|String|
|policy_arn|String|
|policy_name|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|