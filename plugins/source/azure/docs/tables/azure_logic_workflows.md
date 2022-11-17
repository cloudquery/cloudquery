# Table: azure_logic_workflows

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic#Workflow

The primary key for this table is **id**.

## Relations

The following tables depend on azure_logic_workflows:
  - [azure_logic_diagnostic_settings](azure_logic_diagnostic_settings.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|identity|JSON|
|location|String|
|access_control|JSON|
|endpoints_configuration|JSON|
|integration_account|JSON|
|integration_service_environment|JSON|
|parameters|JSON|
|state|String|
|access_endpoint|String|
|changed_time|Timestamp|
|created_time|Timestamp|
|provisioning_state|String|
|sku|JSON|
|version|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|