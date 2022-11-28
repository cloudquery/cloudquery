# Table: azure_sql_database_blob_auditing_policies

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#DatabaseBlobAuditingPolicy

The primary key for this table is **id**.

## Relations
This table depends on [azure_sql_databases](azure_sql_databases.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sql_database_id|String|
|kind|String|
|state|String|
|storage_endpoint|String|
|storage_account_access_key|String|
|retention_days|Int|
|audit_actions_and_groups|StringArray|
|storage_account_subscription_id|UUID|
|is_storage_secondary_key_in_use|Bool|
|is_azure_monitor_target_enabled|Bool|
|queue_delay_ms|Int|
|id (PK)|String|
|name|String|
|type|String|