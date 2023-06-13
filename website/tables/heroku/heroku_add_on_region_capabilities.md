# Table: heroku_add_on_region_capabilities

This table shows data for Heroku Add On Region Capabilities.

https://devcenter.heroku.com/articles/platform-api-reference#add-on-region-capability

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|addon_service|`json`|
|region|`json`|
|supports_private_networking|`bool`|