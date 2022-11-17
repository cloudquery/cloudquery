# Table: azure_sql_database_blob_auditing_policies

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql#DatabaseBlobAuditingPolicy

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
|state|String|
|audit_actions_and_groups|StringArray|
|is_azure_monitor_target_enabled|Bool|
|is_managed_identity_in_use|Bool|
|is_storage_secondary_key_in_use|Bool|
|queue_delay_ms|Int|
|retention_days|Int|
|storage_account_access_key|String|
|storage_account_subscription_id|String|
|storage_endpoint|String|
|id (PK)|String|
|kind|String|
|name|String|
|type|String|
|database_id|String|