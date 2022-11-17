# Table: azure_servicebus_topics

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2#SBTopic

The primary key for this table is **id**.

## Relations
This table depends on [azure_servicebus_namespaces](azure_servicebus_namespaces.md).

The following tables depend on azure_servicebus_topics:
  - [azure_servicebus_authorization_rules](azure_servicebus_authorization_rules.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|auto_delete_on_idle|String|
|default_message_time_to_live|String|
|duplicate_detection_history_time_window|String|
|enable_batched_operations|Bool|
|enable_express|Bool|
|enable_partitioning|Bool|
|max_message_size_in_kilobytes|Int|
|max_size_in_megabytes|Int|
|requires_duplicate_detection|Bool|
|status|String|
|support_ordering|Bool|
|accessed_at|Timestamp|
|count_details|JSON|
|created_at|Timestamp|
|size_in_bytes|Int|
|subscription_count|Int|
|updated_at|Timestamp|
|id (PK)|String|
|location|String|
|name|String|
|system_data|JSON|
|type|String|
|namespace_id|String|