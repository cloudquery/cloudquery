# Table: gcp_redis_instances



The primary key for this table is **name**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|name (PK)|String|
|display_name|String|
|labels|JSON|
|location_id|String|
|alternative_location_id|String|
|redis_version|String|
|reserved_ip_range|String|
|secondary_ip_range|String|
|host|String|
|port|Int|
|current_location_id|String|
|create_time|Timestamp|
|state|String|
|status_message|String|
|redis_configs|JSON|
|tier|String|
|memory_size_gb|Int|
|authorized_network|String|
|persistence_iam_identity|String|
|connect_mode|String|
|auth_enabled|Bool|
|server_ca_certs|JSON|
|transit_encryption_mode|String|
|maintenance_policy|JSON|
|maintenance_schedule|JSON|
|replica_count|Int|
|nodes|JSON|
|read_endpoint|String|
|read_endpoint_port|Int|
|read_replicas_mode|String|