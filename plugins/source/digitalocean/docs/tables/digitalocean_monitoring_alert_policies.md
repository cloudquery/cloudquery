# Table: digitalocean_monitoring_alert_policies


The primary key for this table is **uuid**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|