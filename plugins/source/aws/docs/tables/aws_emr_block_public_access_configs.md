# Table: aws_emr_block_public_access_configs

This table shows data for Amazon EMR Block Public Access Configs.

https://docs.aws.amazon.com/emr/latest/APIReference/API_GetBlockPublicAccessConfiguration.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|block_public_access_configuration|`json`|
|block_public_access_configuration_metadata|`json`|