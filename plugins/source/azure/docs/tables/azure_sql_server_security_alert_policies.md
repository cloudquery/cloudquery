# Table: azure_sql_server_security_alert_policies

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#ServerSecurityAlertPolicy

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
|state|String|
|disabled_alerts|StringArray|
|email_addresses|StringArray|
|email_account_admins|Bool|
|storage_endpoint|String|
|storage_account_access_key|String|
|retention_days|Int|
|creation_time|Timestamp|
|id (PK)|String|
|name|String|
|type|String|