# Table: aws_emr_block_public_access_configs

This table shows data for Amazon EMR Block Public Access Configs.

https://docs.aws.amazon.com/emr/latest/APIReference/API_GetBlockPublicAccessConfiguration.html

The composite primary key for this table is (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|block_public_access_configuration|`json`|
|block_public_access_configuration_metadata|`json`|
|result_metadata|`json`|