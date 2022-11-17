# Table: azure_sql_managed_instance_encryption_protectors

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql#ManagedInstanceEncryptionProtector

The primary key for this table is **id**.

## Relations
This table depends on [azure_sql_managed_instances](azure_sql_managed_instances.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|server_key_type|String|
|auto_rotation_enabled|Bool|
|server_key_name|String|
|thumbprint|String|
|uri|String|
|id (PK)|String|
|kind|String|
|name|String|
|type|String|
|managed_instance_id|String|