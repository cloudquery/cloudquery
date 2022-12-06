# Table: azure_sql_server_dev_ops_auditing_settings

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#ServerDevOpsAuditingSettings

The primary key for this table is **id**.

## Relations
This table depends on [azure_sql_servers](azure_sql_servers.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sql_server_id|String|
|system_data|JSON|
|is_azure_monitor_target_enabled|Bool|
|state|String|
|storage_endpoint|String|
|storage_account_access_key|String|
|storage_account_subscription_id|UUID|
|id (PK)|String|
|name|String|
|type|String|