
# Table: aws_dynamodb_table_local_secondary_indexes
Represents the properties of a local secondary index.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|table_cq_id|uuid|Unique CloudQuery ID of aws_dynamodb_tables table (FK)|
|arn|text|The Amazon Resource Name (ARN) that uniquely identifies the index.|
|name|text|Represents the name of the local secondary index.|
|index_size_bytes|bigint|The total size of the specified index, in bytes|
|item_count|bigint|The number of items in the specified index|
|key_schema|jsonb|The complete key schema for the local secondary index, consisting of one or more pairs of attribute names and key types:  * HASH - partition key  * RANGE - sort key  The partition key of an item is also known as its hash attribute|
|projection_non_key_attributes|text[]|Represents the non-key attribute names which will be projected into the index. For local secondary indexes, the total count of NonKeyAttributes summed across all of the local secondary indexes, must not exceed 20|
|projection_type|text|The set of attributes that are projected into the index:  * KEYS_ONLY - Only the index and primary keys are projected into the index.  * INCLUDE - In addition to the attributes described in KEYS_ONLY, the secondary index will include other non-key attributes that you specify.  * ALL - All of the table attributes are projected into the index.|
