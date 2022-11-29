# Table: azure_monitor_log_profiles

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights#LogProfileResource

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|storage_account_id|String|
|service_bus_rule_id|String|
|locations|StringArray|
|categories|StringArray|
|retention_policy|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|
|kind|String|
|etag|String|