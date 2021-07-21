
# Table: azure_sql_server_devops_audit_settings
ServerDevOpsAuditingSettings a server DevOps auditing settings
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_sql_servers table (FK)|
|created_by|text|A string identifier for the identity that created the resource|
|created_by_type|text|The type of identity that created the resource.|
|created_at_time|timestamp without time zone|The timestamp of resource creation (UTC).|
|last_modified_by|text|A string identifier for the identity that last modified the resource|
|last_modified_by_type|text|The type of identity that last modified the resource|
|last_modified_at_time|timestamp without time zone|The timestamp of last modification (UTC).|
|is_azure_monitor_target_enabled|boolean|Specifies whether DevOps audit events are sent to Azure Monitor|
|state|text|Specifies the state of the audit If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required Possible values include: 'BlobAuditingPolicyStateEnabled', 'BlobAuditingPolicyStateDisabled'|
|storage_endpoint|text|Specifies the blob storage endpoint.|
|storage_account_access_key|text|Specifies the identifier key of the auditing storage account.|
|storage_account_subscription_id|uuid|Specifies the blob storage subscription Id|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
