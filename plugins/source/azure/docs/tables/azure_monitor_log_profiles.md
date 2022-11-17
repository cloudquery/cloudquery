# Table: azure_monitor_log_profiles

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor#LogProfileResource

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
|categories|StringArray|
|locations|StringArray|
|retention_policy|JSON|
|service_bus_rule_id|String|
|storage_account_id|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|