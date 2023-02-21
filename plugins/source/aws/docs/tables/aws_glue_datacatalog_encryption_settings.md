# Table: aws_glue_datacatalog_encryption_settings

https://docs.aws.amazon.com/glue/latest/webapi/API_GetDataCatalogEncryptionSettings.html

The composite primary key for this table is (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|connection_password_encryption|JSON|
|encryption_at_rest|JSON|