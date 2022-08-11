
# Table: aws_lightsail_database_log_events

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_databases table (FK)|
|created_at|timestamp without time zone|The timestamp when the database log event was created|
|message|text|The message of the database log event|
|log_stream_name|text|An object describing the result of your get relational database log streams request|
