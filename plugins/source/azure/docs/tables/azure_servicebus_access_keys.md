# Table: azure_servicebus_access_keys

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus#AccessKeys

The primary key for this table is **_cq_id**.

## Relations
This table depends on [azure_servicebus_authorization_rules](azure_servicebus_authorization_rules.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|servicebus_authorization_rule_id|String|
|primary_connection_string|String|
|secondary_connection_string|String|
|alias_primary_connection_string|String|
|alias_secondary_connection_string|String|
|primary_key|String|
|secondary_key|String|
|key_name|String|