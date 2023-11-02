# Table: facebookmarketing_ad_studies

This table shows data for Facebook Marketing Ad Studies.

https://developers.facebook.com/docs/marketing-api/reference/ad-study/#Reading

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|business|`json`|
|canceled_time|`timestamp[us, tz=UTC]`|
|client_business|`json`|
|cooldown_start_time|`timestamp[us, tz=UTC]`|
|created_by|`json`|
|created_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|end_time|`timestamp[us, tz=UTC]`|
|id (PK)|`utf8`|
|measurement_contact|`json`|
|name|`utf8`|
|observation_end_time|`timestamp[us, tz=UTC]`|
|results_first_available_date|`utf8`|
|sales_contact|`json`|
|start_time|`timestamp[us, tz=UTC]`|
|type|`utf8`|
|updated_by|`json`|
|updated_time|`timestamp[us, tz=UTC]`|