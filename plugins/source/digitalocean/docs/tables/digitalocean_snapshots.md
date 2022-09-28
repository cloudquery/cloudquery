# Table: digitalocean_snapshots


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|name|String|
|resource_id|String|
|resource_type|String|
|regions|StringArray|
|min_disk_size|Int|
|size_gigabytes|Float|
|created_at|String|
|tags|StringArray|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|