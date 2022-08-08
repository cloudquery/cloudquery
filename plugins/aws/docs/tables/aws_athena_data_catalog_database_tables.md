
# Table: aws_athena_data_catalog_database_tables
Contains metadata for a table
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|data_catalog_database_cq_id|uuid|Unique CloudQuery ID of aws_athena_data_catalog_databases table (FK)|
|name|text|The name of the table|
|create_time|timestamp without time zone|The time that the table was created|
|last_access_time|timestamp without time zone|The last time the table was accessed|
|parameters|jsonb|A set of custom key/value pairs for table properties|
|table_type|text|The type of table|
