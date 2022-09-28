# Table: digitalocean_images


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|Int|
|name|String|
|type|String|
|distribution|String|
|slug|String|
|public|Bool|
|regions|StringArray|
|min_disk_size|Int|
|size_gigabytes|Float|
|created_at|String|
|description|String|
|tags|StringArray|
|status|String|
|error_message|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|