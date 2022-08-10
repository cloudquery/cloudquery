
# Table: aws_glue_crawler_targets_catalog_targets
Specifies an Glue Data Catalog target
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|crawler_cq_id|uuid|Unique CloudQuery ID of aws_glue_crawlers table (FK)|
|database_name|text|The name of the database to be synchronized|
|tables|text[]|A list of the tables to be synchronized|
|connection_name|text|The name of the connection for an Amazon S3-backed Data Catalog table to be a target of the crawl when using a Catalog connection type paired with a NETWORK Connection type|
