# Table: azure_sql_virtual_network_rules


The primary key for this table is **id**.

## Relations
This table depends on [`azure_sql_servers`](azure_sql_servers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|sql_server_id|UUID|
|virtual_network_subnet_id|String|
|ignore_missing_vnet_service_endpoint|Bool|
|state|String|
|id (PK)|String|
|name|String|
|type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|