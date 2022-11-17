# Table: azure_logic_diagnostic_settings

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor#DiagnosticSettingsResource

The primary key for this table is **id**.

## Relations
This table depends on [azure_logic_workflows](azure_logic_workflows.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|event_hub_authorization_rule_id|String|
|event_hub_name|String|
|log_analytics_destination_type|String|
|logs|JSON|
|marketplace_partner_id|String|
|metrics|JSON|
|service_bus_rule_id|String|
|storage_account_id|String|
|workspace_id|String|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|
|workflow_id|String|