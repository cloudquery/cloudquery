
# Table: aws_dynamodb_tables
Information about a DynamoDB table.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|The tags associated with the table.|
|archival_summary|jsonb|Contains information about the table archive.|
|attribute_definitions|jsonb|An array of AttributeDefinition objects|
|billing_mode_summary|jsonb|Contains the details for the read/write capacity mode.|
|creation_date_time|timestamp without time zone|The date and time when the table was created, in UNIX epoch time (http://www.epochconverter.com/) format.|
|global_table_version|text|Represents the version of global tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html) in use, if the table is replicated across Amazon Web Services Regions.|
|item_count|bigint|The number of items in the specified table|
|key_schema|jsonb|The primary key structure for the table|
|latest_stream_arn|text|The Amazon Resource Name (ARN) that uniquely identifies the latest stream for this table.|
|latest_stream_label|text|A timestamp, in ISO 8601 format, for this stream|
|provisioned_throughput_last_decrease_date_time|timestamp without time zone|The date and time of the last provisioned throughput decrease for this table.|
|provisioned_throughput_last_increase_date_time|timestamp without time zone|The date and time of the last provisioned throughput increase for this table.|
|provisioned_throughput_number_of_decreases_today|bigint|The number of provisioned throughput decreases for this table during this UTC calendar day|
|provisioned_throughput_read_capacity_units|bigint|The maximum number of strongly consistent reads consumed per second before DynamoDB returns a ThrottlingException|
|provisioned_throughput_write_capacity_units|bigint|The maximum number of writes consumed per second before DynamoDB returns a ThrottlingException.|
|restore_summary|jsonb|Contains details for the restore.|
|inaccessible_encryption_date_time|timestamp without time zone|Indicates the time, in UNIX epoch date format, when DynamoDB detected that the table's KMS key was inaccessible|
|kms_master_key_arn|text|The KMS key ARN used for the KMS encryption.|
|sse_type|text|Server-side encryption type|
|sse_status|text|Represents the current state of server-side encryption|
|stream_specification|jsonb|The current DynamoDB Streams configuration for the table.|
|arn|text|The Amazon Resource Name (ARN) that uniquely identifies the table.|
|table_class_last_update|timestamp without time zone|The date and time at which the table class was last updated.|
|table_class|text|The table class of the specified table|
|id|text|Unique identifier for the table for which the backup was created.|
|name|text|The name of the table.|
|size_bytes|bigint|The total size of the specified table, in bytes|
|status|text|The current state of the table:  * CREATING - The table is being created.  * UPDATING - The table is being updated.  * DELETING - The table is being deleted.  * ACTIVE - The table is ready for use.  * INACCESSIBLE_ENCRYPTION_CREDENTIALS - The KMS key used to encrypt the table in inaccessible|
