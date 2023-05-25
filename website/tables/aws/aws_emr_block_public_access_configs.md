# Table: aws_emr_block_public_access_configs

This table shows data for Amazon EMR Block Public Access Configs.

The composite primary key for this table is (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id (PK)|utf8|
|region (PK)|utf8|
|block_public_access_configuration|json|
|block_public_access_configuration_metadata|json|
|result_metadata|json|