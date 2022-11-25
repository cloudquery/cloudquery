# Table: aws_ssoadmin_instances

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_InstanceMetadata.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_ssoadmin_instances:
  - [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|identity_store_id|String|
|instance_arn|String|