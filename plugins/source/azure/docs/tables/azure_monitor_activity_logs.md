# Table: azure_monitor_activity_logs

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor#EventData

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|authorization|JSON|
|caller|String|
|category|JSON|
|claims|JSON|
|correlation_id|String|
|description|String|
|event_data_id|String|
|event_name|JSON|
|event_timestamp|Timestamp|
|http_request|JSON|
|id (PK)|String|
|level|String|
|operation_id|String|
|operation_name|JSON|
|properties|JSON|
|resource_group_name|String|
|resource_id|String|
|resource_provider_name|JSON|
|resource_type|JSON|
|status|JSON|
|sub_status|JSON|
|submission_timestamp|Timestamp|
|subscription_id|String|
|tenant_id|String|