# Table: azure_servicebus_access_keys

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2#AccessKeys

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
|alias_primary_connection_string|String|
|alias_secondary_connection_string|String|
|key_name|String|
|primary_connection_string|String|
|primary_key|String|
|secondary_connection_string|String|
|secondary_key|String|
|authorization_rule_id|String|