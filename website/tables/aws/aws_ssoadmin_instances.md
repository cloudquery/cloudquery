# Table: aws_ssoadmin_instances

This table shows data for Ssoadmin Instances.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_InstanceMetadata.html

The primary key for this table is **instance_arn**.

## Relations

The following tables depend on aws_ssoadmin_instances:
  - [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|identity_store_id|`utf8`|
|instance_arn (PK)|`utf8`|