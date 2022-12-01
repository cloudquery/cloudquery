# Table: digitalocean_monitoring_alert_policies



The primary key for this table is **uuid**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|uuid (PK)|String|
|type|String|
|description|String|
|compare|String|
|value|Float|
|window|String|
|entities|StringArray|
|tags|StringArray|
|alerts|JSON|
|enabled|Bool|