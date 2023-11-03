# Table: facebookmarketing_adlabels

This table shows data for Facebook Marketing Ad Labels.

https://developers.facebook.com/docs/marketing-api/reference/ad-label#Reading

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|account|`json`|
|created_time|`timestamp[us, tz=UTC]`|
|id (PK)|`utf8`|
|name|`utf8`|
|updated_time|`timestamp[us, tz=UTC]`|