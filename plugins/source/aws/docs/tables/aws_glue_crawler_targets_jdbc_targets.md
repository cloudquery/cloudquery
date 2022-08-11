
# Table: aws_glue_crawler_targets_jdbc_targets
Specifies a JDBC data store to crawl
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|crawler_cq_id|uuid|Unique CloudQuery ID of aws_glue_crawlers table (FK)|
|connection_name|text|The name of the connection to use to connect to the JDBC target|
|exclusions|text[]|A list of glob patterns used to exclude from the crawl|
|path|text|The path of the JDBC target|
