
# Table: aws_lightsail_database_events
Describes an event for a database
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_databases table (FK)|
|created_at|timestamp without time zone|The timestamp when the database event was created|
|event_categories|text[]|The category that the database event belongs to|
|message|text|The message of the database event|
|resource|text|The database that the database event relates to|
