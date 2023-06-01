# Table: datadog_role_users

This table shows data for Datadog Role Users.

The composite primary key for this table is (**role_id**, **id**).

## Relations

This table depends on [datadog_roles](datadog_roles).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name|`utf8`|
|role_id (PK)|`utf8`|
|attributes|`json`|
|id (PK)|`utf8`|
|relationships|`json`|
|type|`utf8`|
|additional_properties|`json`|