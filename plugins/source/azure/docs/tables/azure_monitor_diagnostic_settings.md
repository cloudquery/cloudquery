# Table: azure_monitor_diagnostic_settings

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights#DiagnosticSettingsResource

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
|monitor_resource_id|String|
|storage_account_id|String|
|service_bus_rule_id|String|
|event_hub_authorization_rule_id|String|
|event_hub_name|String|
|metrics|JSON|
|logs|JSON|
|workspace_id|String|
|log_analytics_destination_type|String|
|id (PK)|String|
|name|String|
|type|String|
|resource_uri|String|