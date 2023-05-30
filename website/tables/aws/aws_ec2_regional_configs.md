# Table: aws_ec2_regional_configs

This table shows data for Amazon Elastic Compute Cloud (EC2) Regional Configs.

The composite primary key for this table is (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|ebs_encryption_enabled_by_default|`bool`|
|ebs_default_kms_key_id|`utf8`|