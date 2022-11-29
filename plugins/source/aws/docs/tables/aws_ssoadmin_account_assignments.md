# Table: aws_ssoadmin_account_assignments

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AccountAssignment.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|permission_set_arn|String|
|principal_id|String|
|principal_type|String|