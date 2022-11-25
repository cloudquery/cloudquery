# Table: azure_security_jit_network_access_policies

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security#JitNetworkAccessPolicy

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|name|String|
|type|String|
|kind|String|
|location|String|
|virtual_machines|JSON|
|requests|JSON|
|provisioning_state|String|