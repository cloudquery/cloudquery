
# Table: aws_athena_work_group_named_queries
A query, where QueryString contains the SQL statements that make up the query
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|work_group_cq_id|uuid|Unique CloudQuery ID of aws_athena_work_groups table (FK)|
|database|text|The database to which the query belongs|
|name|text|The query name|
|query_string|text|The SQL statements that make up the query|
|description|text|The query description|
|named_query_id|text|The unique identifier of the query|
|work_group|text|The name of the workgroup that contains the named query|
