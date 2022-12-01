# Table: digitalocean_storage_volumes



The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|droplet_ids|IntArray|
|region|JSON|
|name|String|
|size_gigabytes|Int|
|description|String|
|created_at|Timestamp|
|filesystem_type|String|
|filesystem_label|String|
|tags|StringArray|