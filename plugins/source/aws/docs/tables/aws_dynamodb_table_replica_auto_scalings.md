# Table: aws_dynamodb_table_replica_auto_scalings

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ReplicaAutoScalingDescription.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_dynamodb_tables](aws_dynamodb_tables.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|table_arn|String|
|global_secondary_indexes|JSON|
|region_name|String|
|replica_provisioned_read_capacity_auto_scaling_settings|JSON|
|replica_provisioned_write_capacity_auto_scaling_settings|JSON|
|replica_status|String|