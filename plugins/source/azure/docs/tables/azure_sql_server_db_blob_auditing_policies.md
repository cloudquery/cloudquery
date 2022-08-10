
# Table: azure_sql_server_db_blob_auditing_policies
Database blob auditing policy
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_sql_servers table (FK)|
|state|text|Specifies the state of the policy.|
|storage_endpoint|text|Specifies the blob storage endpoint.|
|storage_account_access_key|text|Specifies the identifier key of the auditing storage account.|
|retention_days|integer|Specifies the number of days to keep in the audit logs in the storage account|
|audit_actions_and_groups|text[]|Specifies the Actions-Groups and Actions to audit.|
|storage_account_subscription_id|uuid|Specifies the blob storage subscription Id|
|is_storage_secondary_key_in_use|boolean|Specifies whether storageAccountAccessKey value is the storage's secondary key|
|is_azure_monitor_target_enabled|boolean|Specifies whether audit events are sent to Azure Monitor.|
|queue_delay_ms|integer|Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed.|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
