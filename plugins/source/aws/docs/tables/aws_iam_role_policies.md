# Table: aws_iam_role_policies



The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_iam_roles](aws_iam_roles.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|role_arn|String|
|policy_document|JSON|
|policy_name|String|
|role_name|String|
|result_metadata|JSON|