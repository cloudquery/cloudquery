# Table: aws_xray_encryption_configs

This table shows data for AWS X-Ray Encryption Configs.

https://docs.aws.amazon.com/xray/latest/api/API_EncryptionConfig.html

The composite primary key for this table is (**account_id**, **region**, **key_id**, **type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|key_id (PK)|`utf8`|
|status|`utf8`|
|type (PK)|`utf8`|