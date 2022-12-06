# Table: azure_sql_virtual_network_rules

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#VirtualNetworkRule

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
|virtual_network_subnet_id|String|
|ignore_missing_vnet_service_endpoint|Bool|
|state|String|
|id (PK)|String|
|name|String|
|type|String|