# Table: digitalocean_storage_volumes

This table shows data for DigitalOcean Storage Volumes.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Block-Storage

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|region|JSON|
|name|String|
|size_gigabytes|Int|
|description|String|
|droplet_ids|IntArray|
|created_at|Timestamp|
|filesystem_type|String|
|filesystem_label|String|
|tags|StringArray|