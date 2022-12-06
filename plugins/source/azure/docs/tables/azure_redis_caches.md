# Table: azure_redis_caches

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis#ResourceType

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|provisioning_state|String|
|host_name|String|
|port|Int|
|ssl_port|Int|
|access_keys|JSON|
|linked_servers|JSON|
|instances|JSON|
|private_endpoint_connections|JSON|
|sku|JSON|
|subnet_id|String|
|static_ip|String|
|redis_configuration|JSON|
|redis_version|String|
|enable_non_ssl_port|Bool|
|replicas_per_master|Int|
|replicas_per_primary|Int|
|tenant_settings|JSON|
|shard_count|Int|
|minimum_tls_version|String|
|public_network_access|String|
|zones|StringArray|
|tags|JSON|
|location|String|
|id (PK)|String|
|name|String|
|type|String|