# Table: digitalocean_sizes


The primary key for this table is **slug**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|