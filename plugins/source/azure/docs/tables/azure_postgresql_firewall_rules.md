# Table: azure_postgresql_firewall_rules



The primary key for this table is **id**.

## Relations
This table depends on [azure_postgresql_servers](azure_postgresql_servers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|postgresql_server_id|String|
|start_ip_address|String|
|end_ip_address|String|
|id (PK)|String|
|name|String|
|type|String|