# Table: azure_monitor_diagnostic_settings

The primary key for this table is **resource_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|resource_id (PK)|String|
|diagnostic_settings_resource|JSON|