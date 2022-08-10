
# Table: aws_glue_crawler_targets_mongo_db_targets
Specifies an Amazon DocumentDB or MongoDB data store to crawl
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|crawler_cq_id|uuid|Unique CloudQuery ID of aws_glue_crawlers table (FK)|
|connection_name|text|The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target|
|path|text|The path of the Amazon DocumentDB or MongoDB target (database/collection)|
|scan_all|boolean|Indicates whether to scan all the records, or to sample rows from the table Scanning all the records can take a long time when the table is not a high throughput table|
