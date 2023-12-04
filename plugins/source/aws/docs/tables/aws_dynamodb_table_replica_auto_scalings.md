# Table: aws_dynamodb_table_replica_auto_scalings

This table shows data for Amazon DynamoDB Table Replica Auto Scalings.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ReplicaAutoScalingDescription.html

The composite primary key for this table is (**table_arn**, **region_name**).

## Relations

This table depends on [aws_dynamodb_tables](aws_dynamodb_tables.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|table_arn (PK)|`utf8`|
|global_secondary_indexes|`json`|
|region_name (PK)|`utf8`|
|replica_provisioned_read_capacity_auto_scaling_settings|`json`|
|replica_provisioned_write_capacity_auto_scaling_settings|`json`|
|replica_status|`utf8`|