# Table: cloudflare_images

This table shows data for Cloudflare Images.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id (PK)|`utf8`|
|filename|`utf8`|
|metadata|`json`|
|require_signed_urls|`bool`|
|variants|`list<item: utf8, nullable>`|
|uploaded|`timestamp[us, tz=UTC]`|