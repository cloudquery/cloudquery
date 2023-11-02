# Table: digitalocean_snapshots

This table shows data for DigitalOcean Snapshots.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Snapshots

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|resource_id|`utf8`|
|resource_type|`utf8`|
|regions|`list<item: utf8, nullable>`|
|min_disk_size|`int64`|
|size_gigabytes|`float64`|
|created_at|`utf8`|
|tags|`list<item: utf8, nullable>`|