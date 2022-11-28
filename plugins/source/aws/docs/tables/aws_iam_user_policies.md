# Table: aws_iam_user_policies



The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_iam_users](aws_iam_users.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|user_arn|String|
|user_id|String|
|policy_document|JSON|
|policy_name|String|
|user_name|String|
|result_metadata|JSON|