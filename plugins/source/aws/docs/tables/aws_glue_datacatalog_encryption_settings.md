# Table: aws_glue_datacatalog_encryption_settings

This table shows data for Glue Datacatalog Encryption Settings.

https://docs.aws.amazon.com/glue/latest/webapi/API_GetDataCatalogEncryptionSettings.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|connection_password_encryption|`json`|
|encryption_at_rest|`json`|