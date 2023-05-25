# Table: aws_ssoadmin_permission_boundaries

This table shows data for Ssoadmin Permission Boundaries.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_GetInlinePolicyForPermissionSet.html

The composite primary key for this table is (**permission_set_arn**, **instance_arn**).

## Relations

This table depends on [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|permission_set_arn (PK)|utf8|
|instance_arn (PK)|utf8|
|customer_managed_policy_reference|json|
|managed_policy_arn|utf8|