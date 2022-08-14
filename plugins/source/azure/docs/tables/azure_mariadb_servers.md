
# Table: azure_mariadb_servers
Server represents a server.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|sku_name|text|Name - The name of the sku.|
|sku_tier|text|The tier of the particular SKU.|
|sku_capacity|integer|The scale up/out capacity, representing server's compute units.|
|sku_size|text|The size code, to be interpreted by resource as appropriate.|
|sku_family|text|The family of hardware.|
|administrator_login|text|The administrator's login name of a server.|
|version|text|Server version.|
|ssl_enforcement|text|Enable ssl enforcement or not when connect to server.|
|user_visible_state|text|A state of a server that is visible to user.|
|fully_qualified_domain_name|text|The fully qualified domain name of a server.|
|earliest_restore_date_time|timestamp without time zone||
|backup_retention_days|integer|Backup retention days for the server.|
|geo_redundant_backup|text|Enable Geo-redundant or not for server backup.|
|storage_mb|integer|Max storage allowed for a server.|
|storage_autogrow|text|Enable Storage Auto Grow|
|replication_role|text|The replication role of the server.|
|master_server_id|text|The master server id of a replica server.|
|replica_capacity|integer|The maximum number of replicas that a master server can have.|
|public_network_access|text|Whether or not public network access is allowed for this server.|
|tags|jsonb|Resource tags.|
|location|text|The geo-location where the resource lives|
|id|text|Fully qualified resource ID for the resource|
|name|text|The name of the resource.|
|type|text|The type of the resource.|
