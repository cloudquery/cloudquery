# Table: aws_ram_resource_shares

This table shows data for RAM Resource Shares.

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShare.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**).
## Relations

The following tables depend on aws_ram_resource_shares:
  - [aws_ram_resource_share_permissions](aws_ram_resource_share_permissions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
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