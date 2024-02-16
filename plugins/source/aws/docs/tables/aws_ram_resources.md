# Table: aws_ram_resources

This table shows data for RAM Resources.

https://docs.aws.amazon.com/ram/latest/APIReference/API_Resource.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**, **resource_share_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|resource_group_arn|`utf8`|
|resource_region_scope|`utf8`|
|resource_share_arn|`utf8`|
|status|`utf8`|
|status_message|`utf8`|
|type|`utf8`|