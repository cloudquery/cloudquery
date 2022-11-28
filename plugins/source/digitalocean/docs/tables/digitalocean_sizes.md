# Table: digitalocean_sizes



The primary key for this table is **slug**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|slug (PK)|String|
|memory|Int|
|vcpus|Int|
|disk|Int|
|price_monthly|Float|
|price_hourly|Float|
|regions|StringArray|
|available|Bool|
|transfer|Float|
|description|String|