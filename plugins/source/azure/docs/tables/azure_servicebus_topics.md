# Table: azure_servicebus_topics

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus#SBTopic

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
|servicebus_namespace_id|String|
|size_in_bytes|Int|
|created_at|Timestamp|
|updated_at|Timestamp|
|accessed_at|Timestamp|
|subscription_count|Int|
|count_details|JSON|
|default_message_time_to_live|String|
|max_size_in_megabytes|Int|
|max_message_size_in_kilobytes|Int|
|requires_duplicate_detection|Bool|
|duplicate_detection_history_time_window|String|
|enable_batched_operations|Bool|
|status|String|
|support_ordering|Bool|
|auto_delete_on_idle|String|
|enable_partitioning|Bool|
|enable_express|Bool|
|system_data|JSON|
|id (PK)|String|
|name|String|
|type|String|