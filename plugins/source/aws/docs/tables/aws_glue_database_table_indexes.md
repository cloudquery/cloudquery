
# Table: aws_glue_database_table_indexes
A descriptor for a partition index in a table
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_table_cq_id|uuid|Unique CloudQuery ID of aws_glue_database_tables table (FK)|
|index_name|text|The name of the partition index|
|index_status|text|The status of the partition index|
|keys|jsonb|A list of one or more keys, as KeySchemaElement structures, for the partition index|
|backfill_errors|jsonb|A list of errors that can occur when registering partition indexes for an existing table|
