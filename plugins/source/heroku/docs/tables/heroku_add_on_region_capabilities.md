# Table: heroku_add_on_region_capabilities

https://devcenter.heroku.com/articles/platform-api-reference#add-on-region-capability

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|addon_service|JSON|
|region|JSON|
|supports_private_networking|Bool|