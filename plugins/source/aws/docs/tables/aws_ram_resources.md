# Table: aws_ram_resources

This table shows data for RAM Resources.

https://docs.aws.amazon.com/ram/latest/APIReference/API_Resource.html

The composite primary key for this table is (**account_id**, **region**, **arn**, **resource_share_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|arn (PK)|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|resource_group_arn|`utf8`|
|resource_region_scope|`utf8`|
|resource_share_arn (PK)|`utf8`|
|status|`utf8`|
|status_message|`utf8`|
|type|`utf8`|