# Table: azure_monitor_activity_logs

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights#EventData

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|authorization|JSON|
|claims|JSON|
|caller|String|
|description|String|
|id (PK)|String|
|event_data_id|String|
|correlation_id|String|
|event_name|JSON|
|category|JSON|
|http_request|JSON|
|level|String|
|resource_group_name|String|
|resource_provider_name|JSON|
|resource_id|String|
|resource_type|JSON|
|operation_id|String|
|operation_name|JSON|
|properties|JSON|
|status|JSON|
|sub_status|JSON|
|event_timestamp|Timestamp|
|submission_timestamp|Timestamp|
|tenant_id|String|