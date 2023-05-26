# Table: azure_monitor_resources

This table shows data for Azure Monitor Resources.

https://learn.microsoft.com/en-us/rest/api/resources/resources/list#genericresourceexpanded

The primary key for this table is **id**.

## Relations

The following tables depend on azure_monitor_resources:
  - [azure_monitor_diagnostic_settings](azure_monitor_diagnostic_settings)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|subscription_id|utf8|
|extended_location|json|
|identity|json|
|kind|utf8|
|location|utf8|
|managed_by|utf8|
|plan|json|
|sku|json|
|tags|json|
|changed_time|timestamp[us, tz=UTC]|
|created_time|timestamp[us, tz=UTC]|
|id (PK)|utf8|
|name|utf8|
|provisioning_state|utf8|
|type|utf8|