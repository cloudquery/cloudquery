# Table: azure_cdn_security_policies

This table shows data for Azure Cdn Security Policies.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn@v1.0.0#SecurityPolicy

The primary key for this table is **id**.

## Relations

This table depends on [azure_cdn_profiles](azure_cdn_profiles).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|properties|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|