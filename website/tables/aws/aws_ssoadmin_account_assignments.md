# Table: aws_ssoadmin_account_assignments

This table shows data for Ssoadmin Account Assignments.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AccountAssignment.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|permission_set_arn|`utf8`|
|principal_id|`utf8`|
|principal_type|`utf8`|