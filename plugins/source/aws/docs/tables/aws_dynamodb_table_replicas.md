
# Table: aws_dynamodb_table_replicas
Contains the details of the replica.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|table_cq_id|uuid|Unique CloudQuery ID of aws_dynamodb_tables table (FK)|
|global_secondary_indexes|jsonb|Replica-specific global secondary index settings.|
|kms_master_key_id|text|The KMS key of the replica that will be used for KMS encryption.|
|provisioned_throughput_override_read_capacity_units|bigint|Replica-specific read capacity units|
|region_name|text|The name of the Region.|
|replica_inaccessible_date_time|timestamp without time zone|The time at which the replica was first detected as inaccessible|
|replica_status|text|The current state of the replica:  * CREATING - The replica is being created.  * UPDATING - The replica is being updated.  * DELETING - The replica is being deleted.  * ACTIVE - The replica is ready for use.  * REGION_DISABLED - The replica is inaccessible because the Amazon Web Services Region has been disabled|
|replica_status_description|text|Detailed information about the replica status.|
|replica_status_percent_progress|text|Specifies the progress of a Create, Update, or Delete action on the replica as a percentage.|
|summary_last_update_date_time|timestamp without time zone|The date and time at which the table class was last updated.|
|summary_table_class|text|The table class of the specified table|
