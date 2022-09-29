# Table: azure_logic_diagnostic_settings


The primary key for this table is **id**.

## Relations
This table depends on [`azure_logic_workflows`](azure_logic_workflows.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|logic_workflow_id|String|
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|