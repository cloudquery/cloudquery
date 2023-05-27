# Table: azure_monitor_subscription_diagnostic_settings

This table shows data for Azure Monitor Subscription Diagnostic Settings.

https://learn.microsoft.com/en-us/rest/api/monitor/subscription-diagnostic-settings/list?tabs=HTTP#subscriptiondiagnosticsettingsresource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|subscription_id|utf8|
|properties|json|
|id (PK)|utf8|
|name|utf8|
|system_data|json|
|type|utf8|