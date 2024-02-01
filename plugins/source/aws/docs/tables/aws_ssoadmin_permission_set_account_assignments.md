# Table: aws_ssoadmin_permission_set_account_assignments

This table shows data for Ssoadmin Permission Set Account Assignments.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AccountAssignment.html
The 'request_account_id' and 'request_region' columns are added to show the account_id and region of where the request was made from.

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**instance_arn**, **account_id**, **permission_set_arn**, **principal_id**, **principal_type**).
## Relations

This table depends on [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|instance_arn|`utf8`|
|account_id|`utf8`|
|permission_set_arn|`utf8`|
|principal_id|`utf8`|
|principal_type|`utf8`|