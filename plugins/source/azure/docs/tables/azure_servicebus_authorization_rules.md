# Table: azure_servicebus_authorization_rules

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus#SBAuthorizationRule

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
|servicebus_topic_id|String|
|rights|StringArray|
|system_data|JSON|
|id (PK)|String|
|name|String|
|type|String|