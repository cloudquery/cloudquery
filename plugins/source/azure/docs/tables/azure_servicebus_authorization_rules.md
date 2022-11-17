# Table: azure_servicebus_authorization_rules

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2#SBAuthorizationRule

The primary key for this table is **id**.

## Relations
This table depends on [azure_servicebus_topics](azure_servicebus_topics.md).

The following tables depend on azure_servicebus_authorization_rules:
  - [azure_servicebus_access_keys](azure_servicebus_access_keys.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|rights|StringArray|
|id (PK)|String|
|location|String|
|name|String|
|system_data|JSON|
|type|String|
|topic_id|String|