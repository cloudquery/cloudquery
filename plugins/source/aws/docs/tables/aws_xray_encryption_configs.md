# Table: aws_xray_encryption_configs

This table shows data for AWS X-Ray Encryption Configs.

https://docs.aws.amazon.com/xray/latest/api/API_EncryptionConfig.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|key_id|`utf8`|
|status|`utf8`|
|type|`utf8`|