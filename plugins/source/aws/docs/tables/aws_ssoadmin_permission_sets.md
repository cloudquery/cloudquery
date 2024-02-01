# Table: aws_ssoadmin_permission_sets

This table shows data for Ssoadmin Permission Sets.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_PermissionSet.html. 
The 'request_account_id' and 'request_region' columns are added to show the account_id and region of where the request was made from.

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**instance_arn**, **permission_set_arn**).
## Relations

This table depends on [aws_ssoadmin_instances](aws_ssoadmin_instances.md).

The following tables depend on aws_ssoadmin_permission_sets:
  - [aws_ssoadmin_permission_set_account_assignments](aws_ssoadmin_permission_set_account_assignments.md)
  - [aws_ssoadmin_permission_set_customer_managed_policies](aws_ssoadmin_permission_set_customer_managed_policies.md)
  - [aws_ssoadmin_permission_set_inline_policies](aws_ssoadmin_permission_set_inline_policies.md)
  - [aws_ssoadmin_permission_set_managed_policies](aws_ssoadmin_permission_set_managed_policies.md)
  - [aws_ssoadmin_permission_set_permissions_boundaries](aws_ssoadmin_permission_set_permissions_boundaries.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|instance_arn|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|name|`utf8`|
|permission_set_arn|`utf8`|
|relay_state|`utf8`|
|session_duration|`utf8`|