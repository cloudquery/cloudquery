# Table: aws_glue_datacatalog_encryption_settings



The primary key for this table is **account_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|connection_password_encryption|JSON|
|encryption_at_rest|JSON|