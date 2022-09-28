# Table: heroku_add_on_region_capabilities
https://devcenter.heroku.com/articles/platform-api-reference#add-on-region-capability-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|addon_service|JSON|
|id (PK)|String|
|region|JSON|
|supports_private_networking|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|