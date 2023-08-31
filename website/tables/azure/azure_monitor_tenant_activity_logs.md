# Table: azure_monitor_tenant_activity_logs

This table shows data for Azure Monitor Tenant Activity Logs.

https://learn.microsoft.com/en-us/rest/api/monitor/tenant-activity-logs/list?tabs=HTTP#eventdata

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|authorization|`json`|
|caller|`utf8`|
|category|`json`|
|claims|`json`|
|correlation_id|`utf8`|
|description|`utf8`|
|event_data_id|`utf8`|
|event_name|`json`|
|event_timestamp|`timestamp[us, tz=UTC]`|
|http_request|`json`|
|id (PK)|`utf8`|
|level|`utf8`|
|operation_id|`utf8`|
|operation_name|`json`|
|properties|`json`|
|resource_group_name|`utf8`|
|resource_id|`utf8`|
|resource_provider_name|`json`|
|resource_type|`json`|
|status|`json`|
|sub_status|`json`|
|submission_timestamp|`timestamp[us, tz=UTC]`|
|tenant_id|`utf8`|