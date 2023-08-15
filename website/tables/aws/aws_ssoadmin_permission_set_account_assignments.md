# Table: aws_ssoadmin_permission_set_account_assignments

This table shows data for Ssoadmin Permission Set Account Assignments.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AccountAssignment.html
The 'request_account_id' and 'request_region' columns are added to show the account_id and region of where the request was made from.

The composite primary key for this table is (**instance_arn**, **account_id**, **permission_set_arn**, **principal_id**, **principal_type**).

## Relations

This table depends on [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|instance_arn (PK)|`utf8`|
|account_id (PK)|`utf8`|
|permission_set_arn (PK)|`utf8`|
|principal_id (PK)|`utf8`|
|principal_type (PK)|`utf8`|