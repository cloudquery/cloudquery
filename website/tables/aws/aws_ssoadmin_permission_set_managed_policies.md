# Table: aws_ssoadmin_permission_set_managed_policies

This table shows data for Ssoadmin Permission Set Managed Policies.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ListManagedPoliciesInPermissionSet.html

The composite primary key for this table is (**permission_set_arn**, **instance_arn**, **arn**).

## Relations

This table depends on [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|permission_set_arn (PK)|`utf8`|
|instance_arn (PK)|`utf8`|
|arn (PK)|`utf8`|
|name|`utf8`|