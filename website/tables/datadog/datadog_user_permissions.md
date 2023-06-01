# Table: datadog_user_permissions

This table shows data for Datadog User Permissions.

The composite primary key for this table is (**user_id**, **id**).

## Relations

This table depends on [datadog_users](datadog_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name|`utf8`|
|user_id (PK)|`utf8`|
|attributes|`json`|
|id (PK)|`utf8`|
|type|`utf8`|
|additional_properties|`json`|