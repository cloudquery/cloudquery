# Table: azure_datalake_store_accounts

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore#Account

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|identity|JSON|
|location|String|
|name|String|
|account_id|String|
|creation_time|Timestamp|
|current_tier|String|
|default_group|String|
|encryption_config|JSON|
|encryption_provisioning_state|String|
|encryption_state|String|
|endpoint|String|
|firewall_allow_azure_ips|String|
|firewall_rules|JSON|
|firewall_state|String|
|last_modified_time|Timestamp|
|new_tier|String|
|provisioning_state|String|
|state|String|
|trusted_id_provider_state|String|
|trusted_id_providers|JSON|
|virtual_network_rules|JSON|
|tags|JSON|
|type|String|