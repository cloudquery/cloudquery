# Table: aws_ssoadmin_instances

This table shows data for Ssoadmin Instances.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_InstanceMetadata.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **instance_arn**.
## Relations

The following tables depend on aws_ssoadmin_instances:
  - [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|created_date|`timestamp[us, tz=UTC]`|
|identity_store_id|`utf8`|
|instance_arn|`utf8`|
|name|`utf8`|
|owner_account_id|`utf8`|
|status|`utf8`|