# Table: aws_glue_crawlers



The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|classifiers|StringArray|
|configuration|String|
|crawl_elapsed_time|Int|
|crawler_security_configuration|String|
|creation_time|Timestamp|
|database_name|String|
|description|String|
|lake_formation_configuration|JSON|
|last_crawl|JSON|
|last_updated|Timestamp|
|lineage_configuration|JSON|
|name|String|
|recrawl_policy|JSON|
|role|String|
|schedule|JSON|
|schema_change_policy|JSON|
|state|String|
|table_prefix|String|
|targets|JSON|
|version|Int|