
# Table: aws_glue_crawler_targets_s3_targets
Specifies a data store in Amazon Simple Storage Service (Amazon S3)
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|crawler_cq_id|uuid|Unique CloudQuery ID of aws_glue_crawlers table (FK)|
|connection_name|text|The name of a connection which allows a job or crawler to access data in Amazon S3 within an Amazon Virtual Private Cloud environment (Amazon VPC)|
|dlq_event_queue_arn|text|A valid Amazon dead-letter SQS ARN|
|event_queue_arn|text|A valid Amazon SQS ARN|
|exclusions|text[]|A list of glob patterns used to exclude from the crawl|
|path|text|The path to the Amazon S3 target|
|sample_size|bigint|Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset|
