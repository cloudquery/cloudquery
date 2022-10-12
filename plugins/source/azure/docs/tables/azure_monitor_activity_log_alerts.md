# Table: azure_monitor_activity_log_alerts



The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
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