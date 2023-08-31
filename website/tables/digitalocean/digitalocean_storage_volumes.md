# Table: digitalocean_storage_volumes

This table shows data for DigitalOcean Storage Volumes.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Block-Storage

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|region|`json`|
|name|`utf8`|
|size_gigabytes|`int64`|
|description|`utf8`|
|droplet_ids|`list<item: int64, nullable>`|
|created_at|`timestamp[us, tz=UTC]`|
|filesystem_type|`utf8`|
|filesystem_label|`utf8`|
|tags|`list<item: utf8, nullable>`|