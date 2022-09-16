
# Table: aws_dynamodb_table_global_secondary_indexes
Represents the properties of a global secondary index.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|table_cq_id|uuid|Unique CloudQuery ID of aws_dynamodb_tables table (FK)|
|backfilling|boolean|Indicates whether the index is currently backfilling|
|arn|text|The Amazon Resource Name (ARN) that uniquely identifies the index.|
|name|text|The name of the global secondary index.|
|index_size_bytes|bigint|The total size of the specified index, in bytes|
|status|text|The current state of the global secondary index:  * CREATING - The index is being created.  * UPDATING - The index is being updated.  * DELETING - The index is being deleted.  * ACTIVE - The index is ready for use.|
|item_count|bigint|The number of items in the specified index|
|key_schema|jsonb|The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:  * HASH - partition key  * RANGE - sort key  The partition key of an item is also known as its hash attribute|
|projection_non_key_attributes|text[]|Represents the non-key attribute names which will be projected into the index. For local secondary indexes, the total count of NonKeyAttributes summed across all of the local secondary indexes, must not exceed 20|
|projection_type|text|The set of attributes that are projected into the index:  * KEYS_ONLY - Only the index and primary keys are projected into the index.  * INCLUDE - In addition to the attributes described in KEYS_ONLY, the secondary index will include other non-key attributes that you specify.  * ALL - All of the table attributes are projected into the index.|
|provisioned_throughput_last_decrease_date_time|timestamp without time zone|The date and time of the last provisioned throughput decrease for this table.|
|provisioned_throughput_last_increase_date_time|timestamp without time zone|The date and time of the last provisioned throughput increase for this table.|
|provisioned_throughput_number_of_decreases_today|bigint|The number of provisioned throughput decreases for this table during this UTC calendar day|
|provisioned_throughput_read_capacity_units|bigint|The maximum number of strongly consistent reads consumed per second before DynamoDB returns a ThrottlingException|
|provisioned_throughput_write_capacity_units|bigint|The maximum number of writes consumed per second before DynamoDB returns a ThrottlingException.|
