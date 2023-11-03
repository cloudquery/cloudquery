# Table: facebookmarketing_adimages

This table shows data for Facebook Marketing Ad Images.

https://developers.facebook.com/docs/marketing-api/reference/ad-image#Reading

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|creatives|`list<item: utf8, nullable>`|
|hash|`utf8`|
|height|`int64`|
|id (PK)|`utf8`|
|is_associated_creatives_in_adgroups|`bool`|
|name|`utf8`|
|original_height|`int64`|
|original_width|`int64`|
|permalink_url|`utf8`|
|status|`utf8`|
|updated_time|`timestamp[us, tz=UTC]`|
|url|`utf8`|
|url_128|`utf8`|
|width|`int64`|