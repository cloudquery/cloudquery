# Table: azure_logic_workflows


The primary key for this table is **id**.

## Relations
The following tables depend on `azure_logic_workflows`:
  - [`azure_logic_diagnostic_settings`](azure_logic_diagnostic_settings.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|provisioning_state|String|
|created_time|Timestamp|
|changed_time|Timestamp|
|state|String|
|version|String|
|access_endpoint|String|
|endpoints_configuration|JSON|
|access_control|JSON|
|sku|JSON|
|integration_account|JSON|
|integration_service_environment|JSON|
|parameters|JSON|
|identity|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|