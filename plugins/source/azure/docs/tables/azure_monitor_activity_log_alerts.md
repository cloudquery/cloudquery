# Table: azure_monitor_activity_log_alerts

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor#ActivityLogAlertResource

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|actions|JSON|
|condition|JSON|
|scopes|StringArray|
|description|String|
|enabled|Bool|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|