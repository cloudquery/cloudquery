# Table: azure_monitor_activity_log_alerts

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|properties|JSON|
|tags|JSON|
|id|String|
|name|String|
|type|String|