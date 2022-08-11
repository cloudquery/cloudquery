
# Table: aws_glue_crawler_targets_dynamo_db_targets
Specifies an Amazon DynamoDB table to crawl
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|crawler_cq_id|uuid|Unique CloudQuery ID of aws_glue_crawlers table (FK)|
|path|text|The name of the DynamoDB table to crawl|
|scan_all|boolean|Indicates whether to scan all the records, or to sample rows from the table Scanning all the records can take a long time when the table is not a high throughput table|
|scan_rate|float|The percentage of the configured read capacity units to use by the Glue crawler Read capacity units is a term defined by DynamoDB, and is a numeric value that acts as rate limiter for the number of reads that can be performed on that table per second|
