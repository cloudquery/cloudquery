# Table: azure_monitor_diagnostic_settings

https://learn.microsoft.com/en-us/rest/api/monitor/diagnostic-settings/list?tabs=HTTP#diagnosticsettingsresource

The primary key for this table is **id**.

## Relations

This table depends on [azure_monitor_resources](azure_monitor_resources.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|properties|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|
|resource_id|String|