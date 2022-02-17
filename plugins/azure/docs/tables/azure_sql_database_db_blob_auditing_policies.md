
# Table: azure_sql_database_db_blob_auditing_policies
DatabaseBlobAuditingPolicy a database blob auditing policy
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of azure_sql_databases table (FK)|
|kind|text|Resource kind|
|state|text|Specifies the state of the policy|
|storage_endpoint|text|Specifies the blob storage endpoint (eg|
|storage_account_access_key|text|Specifies the identifier key of the auditing storage account If state is Enabled and storageEndpoint is specified, not specifying the storageAccountAccessKey will use SQL server system-assigned managed identity to access the storage Prerequisites for using managed identity authentication: 1|
|retention_days|integer|Specifies the number of days to keep in the audit logs in the storage account|
|audit_actions_and_groups|text[]|this will audit all the queries and stored procedures executed against the database, as well as successful and failed logins:  BATCH_COMPLETED_GROUP, SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP, FAILED_DATABASE_AUTHENTICATION_GROUP  This above combination is also the set that is configured by default when enabling auditing from the Azure portal  The supported action groups to audit are (note: choose only specific groups that cover your auditing needs|
|storage_account_subscription_id|uuid|Specifies the blob storage subscription Id|
|is_storage_secondary_key_in_use|boolean|Specifies whether storageAccountAccessKey value is the storage's secondary key|
|is_azure_monitor_target_enabled|boolean|Specifies whether audit events are sent to Azure Monitor|
|queue_delay_ms|integer|Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed The default minimum value is 1000 (1 second)|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
