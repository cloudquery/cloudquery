# Table: aws_ssoadmin_permission_set_permissions_boundaries

This table shows data for Ssoadmin Permission Set Permissions Boundaries.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_GetPermissionsBoundaryForPermissionSet.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**permission_set_arn**, **instance_arn**).
## Relations

This table depends on [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|permission_set_arn|`utf8`|
|instance_arn|`utf8`|
|customer_managed_policy_reference|`json`|
|managed_policy_arn|`utf8`|