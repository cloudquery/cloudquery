
# Table: azure_redis_services
Azure Redis service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|tags|jsonb|Resource tags.|
|location|text|The geo-location where the resource lives.|
|id|text|Fully qualified resource ID for the resource.|
|name|text|The name of the resource|
|type|text|The type of the resource|
|provisioning_state|text|Redis instance provisioning status.|
|hostname|text|Redis host name.|
|port|integer|Redis non-SSL port.|
|ssl_port|integer|Redis SSL port.|
|linked_server_ids|text[]|List of the linked servers associated with the cache.|
|instances|jsonb|List of the Redis instances associated with the cache.|
|private_endpoint_connections|jsonb|List of private endpoint connection associated with the specified redis cache.|
|sku_name|text|The type of Redis cache to deploy.|
|sku_family|text|The SKU family to use.|
|sku_capacity|integer|The size of the Redis cache to deploy.|
|subnet_id|text|The full resource ID of a subnet in a virtual network to deploy the Redis cache in.|
|static_ip|inet|Static IP address.|
|configuration|jsonb|All Redis Settings.|
|version|text|Redis version.|
|enable_non_ssl_port|boolean|Specifies whether the non-ssl Redis server port (6379) is enabled.|
|replicas_per_master|integer|The number of replicas to be created per primary.|
|replicas_per_primary|integer|The number of replicas to be created per primary.|
|tenant_settings|jsonb|A dictionary of tenant settings.|
|shard_count|integer|The number of shards to be created on a Premium Cluster Cache.|
|minimum_tls_version|text|Requires clients to use a specified TLS version (or higher) to connect.|
|public_network_access|boolean|Whether or not public endpoint access is allowed for this cache.|
