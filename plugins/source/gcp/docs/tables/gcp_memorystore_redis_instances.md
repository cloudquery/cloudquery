
# Table: gcp_memorystore_redis_instances
A Memorystore for Redis instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project ID of the resource|
|id|text|Memorystore for Redis instance ID|
|alternative_location_id|text|If specified, at least one node will be provisioned in this zone in addition to the zone specified in `location_id`|
|auth_enabled|boolean|Indicates whether OSS Redis AUTH is enabled for the instance|
|authorized_network|text|The full name of the Google Compute Engine network (https://cloud.google.com/vpc/docs/vpc) to which the instance is connected|
|connect_mode|text|The network connect mode of the Redis instance|
|create_time|text|The time the instance was created|
|current_location_id|text|The current zone where the Redis primary node is located|
|customer_managed_key|text|The KMS key reference that the customer provides when trying to create the instance|
|display_name|text|An arbitrary and optional user-provided name for the instance|
|host|text|Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service|
|labels|jsonb|Resource labels to represent user provided metadata|
|location_id|text|The zone where the instance will be provisioned|
|maintenance_policy|jsonb|The maintenance policy for the instance|
|maintenance_schedule|jsonb|Date and time of upcoming maintenance events which have been scheduled|
|memory_size_gb|bigint|Redis memory size in GiB|
|name|text|Unique name of the resource|
|nodes|jsonb|Redis instance nodes properties|
|persistence_config|jsonb|Persistence configuration parameters|
|persistence_iam_identity|text|Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage|
|port|bigint|The port number of the exposed Redis endpoints|
|read_endpoint|text|Hostname or IP address of the exposed readonly Redis endpoint|
|read_endpoint_port|bigint|The port number of the exposed readonly redis endpoint|
|read_replicas_mode|text|Read replicas mode for the instance|
|redis_configs|jsonb|Redis configuration parameters|
|redis_version|text|The version of Redis software|
|replica_count|bigint|The number of replica nodes|
|reserved_ip_range|text|IP range for node placement|
|secondary_ip_range|text|Additional IP range for node placement|
|state|text|The current state of the instance|
|status_message|text|Additional information about the current status of the instance, if available|
|suspension_reasons|text[]|Reasons that causes instance in `SUSPENDED` state|
|tier|text|The service tier of the instance|
|transit_encryption_mode|text|The TLS mode of the Redis instance|
