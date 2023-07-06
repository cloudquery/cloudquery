# Table: facebookmarketing_ad_place_page_sets

This table shows data for Facebook Marketing Ad Place Page Sets.

https://developers.facebook.com/docs/marketing-api/reference/ad-place-page-set#Reading

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id (PK)|`utf8`|
|location_types|`list<item: utf8, nullable>`|
|name|`utf8`|
|pages_count|`int64`|
|parent_page|`json`|