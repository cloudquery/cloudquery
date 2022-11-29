# Table: aws_xray_encryption_configs

https://docs.aws.amazon.com/xray/latest/api/API_EncryptionConfig.html

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|key_id|String|
|status|String|
|type|String|