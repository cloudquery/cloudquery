# Table: digitalocean_images

This table shows data for DigitalOcean Images.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Images

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`int64`|
|name|`utf8`|
|type|`utf8`|
|distribution|`utf8`|
|slug|`utf8`|
|public|`bool`|
|regions|`list<item: utf8, nullable>`|
|min_disk_size|`int64`|
|size_gigabytes|`float64`|
|created_at|`utf8`|
|description|`utf8`|
|tags|`list<item: utf8, nullable>`|
|status|`utf8`|
|error_message|`utf8`|