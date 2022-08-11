
# Table: aws_glue_crawler_targets_delta_targets
Specifies a Delta data store to crawl one or more Delta tables
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|crawler_cq_id|uuid|Unique CloudQuery ID of aws_glue_crawlers table (FK)|
|connection_name|text|The name of the connection to use to connect to the Delta table target|
|delta_tables|text[]|A list of the Amazon S3 paths to the Delta tables|
|write_manifest|boolean|Specifies whether to write the manifest files to the Delta table path|
