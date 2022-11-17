# Table: azure_redis_caches

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2#ResourceInfo

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|sku|JSON|
|enable_non_ssl_port|Bool|
|minimum_tls_version|String|
|public_network_access|String|
|redis_configuration|JSON|
|redis_version|String|
|replicas_per_master|Int|
|replicas_per_primary|Int|
|shard_count|Int|
|static_ip|String|
|subnet_id|String|
|tenant_settings|JSON|
|access_keys|JSON|
|host_name|String|
|instances|JSON|
|linked_servers|JSON|
|port|Int|
|private_endpoint_connections|JSON|
|provisioning_state|String|
|ssl_port|Int|
|identity|JSON|
|tags|JSON|
|zones|StringArray|
|id (PK)|String|
|name|String|
|type|String|