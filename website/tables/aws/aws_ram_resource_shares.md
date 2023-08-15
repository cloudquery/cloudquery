# Table: aws_ram_resource_shares

This table shows data for RAM Resource Shares.

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShare.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Relations

The following tables depend on aws_ram_resource_shares:
  - [aws_ram_resource_share_permissions](aws_ram_resource_share_permissions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|allow_external_principals|`bool`|
|creation_time|`timestamp[us, tz=UTC]`|
|feature_set|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|owning_account_id|`utf8`|
|resource_share_arn|`utf8`|
|status|`utf8`|
|status_message|`utf8`|