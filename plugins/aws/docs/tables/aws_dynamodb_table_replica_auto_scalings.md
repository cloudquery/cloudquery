
# Table: aws_dynamodb_table_replica_auto_scalings
Represents the auto scaling settings of the replica.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|table_cq_id|uuid|Unique CloudQuery ID of aws_dynamodb_tables table (FK)|
|global_secondary_indexes|jsonb|Replica-specific global secondary index auto scaling settings.|
|region_name|text|The Region where the replica exists.|
|read_capacity|jsonb|Represents the auto scaling settings for a global table or global secondary index.|
|write_capacity|jsonb|Represents the auto scaling settings for a global table or global secondary index.|
|replica_status|text|The current state of the replica:  * CREATING - The replica is being created.  * UPDATING - The replica is being updated.  * DELETING - The replica is being deleted.  * ACTIVE - The replica is ready for use.|
