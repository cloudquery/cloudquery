# Table: azure_monitor_tenant_activity_log_alerts

This table shows data for Azure Monitor Tenant Activity Log Alerts.

https://learn.microsoft.com/en-us/rest/api/monitor/activity-log-alerts/list-by-subscription-id?tabs=HTTP#activitylogalertresource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|