# Table: facebookmarketing_adrules

This table shows data for Facebook Marketing Adrules.

https://developers.facebook.com/docs/marketing-api/reference/ad-rule#Reading

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|created_by|`json`|
|created_time|`timestamp[us, tz=UTC]`|
|evaluation_spec|`json`|
|execution_spec|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|schedule_spec|`json`|
|status|`utf8`|
|updated_time|`timestamp[us, tz=UTC]`|