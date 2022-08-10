
# Table: aws_athena_work_group_prepared_statements
A prepared SQL statement for use with Athena
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|work_group_cq_id|uuid|Unique CloudQuery ID of aws_athena_work_groups table (FK)|
|description|text|The description of the prepared statement|
|last_modified_time|timestamp without time zone|The last modified time of the prepared statement|
|query_statement|text|The query string for the prepared statement|
|statement_name|text|The name of the prepared statement|
|work_group_name|text|The name of the workgroup to which the prepared statement belongs|
