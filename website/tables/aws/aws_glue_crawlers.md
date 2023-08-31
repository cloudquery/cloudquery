# Table: aws_glue_crawlers

This table shows data for Glue Crawlers.

https://docs.aws.amazon.com/glue/latest/webapi/API_Crawler.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|classifiers|`list<item: utf8, nullable>`|
|configuration|`utf8`|
|crawl_elapsed_time|`int64`|
|crawler_security_configuration|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|database_name|`utf8`|
|description|`utf8`|
|lake_formation_configuration|`json`|
|last_crawl|`json`|
|last_updated|`timestamp[us, tz=UTC]`|
|lineage_configuration|`json`|
|name|`utf8`|
|recrawl_policy|`json`|
|role|`utf8`|
|schedule|`json`|
|schema_change_policy|`json`|
|state|`utf8`|
|table_prefix|`utf8`|
|targets|`json`|
|version|`int64`|