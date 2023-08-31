# Table: aws_glue_datacatalog_encryption_settings

This table shows data for Glue Datacatalog Encryption Settings.

https://docs.aws.amazon.com/glue/latest/webapi/API_GetDataCatalogEncryptionSettings.html

The composite primary key for this table is (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|connection_password_encryption|`json`|
|encryption_at_rest|`json`|