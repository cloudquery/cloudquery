# Table: azure_monitor_activity_log_alerts


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|scopes|StringArray|
|enabled|Bool|
|condition|JSON|
|actions|JSON|
|description|String|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|