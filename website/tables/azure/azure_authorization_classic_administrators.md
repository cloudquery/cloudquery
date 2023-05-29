# Table: azure_authorization_classic_administrators

This table shows data for Azure Authorization Classic Administrators.

https://learn.microsoft.com/en-us/rest/api/authorization/classic-administrators/list?tabs=HTTP#classicadministrator

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|properties|`json`|
|type|`utf8`|