# Table: aws_ssoadmin_instances

This table shows data for Ssoadmin Instances.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_InstanceMetadata.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_ssoadmin_instances:
  - [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id (PK)|uuid|
|_cq_parent_id|uuid|
|identity_store_id|utf8|
|instance_arn|utf8|