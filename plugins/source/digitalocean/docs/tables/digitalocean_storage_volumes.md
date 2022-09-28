# Table: digitalocean_storage_volumes


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|