# Table: aws_emr_block_public_access_configs

This table shows data for Amazon EMR Block Public Access Configs.

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
|block_public_access_configuration|JSON|
|block_public_access_configuration_metadata|JSON|
|result_metadata|JSON|