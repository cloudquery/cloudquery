# Table: gcp_redis_instances

This table shows data for GCP Redis Instances.

https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance

The primary key for this table is **name**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|labels|`json`|
|location_id|`utf8`|
|alternative_location_id|`utf8`|
|redis_version|`utf8`|
|reserved_ip_range|`utf8`|
|secondary_ip_range|`utf8`|
|host|`utf8`|
|port|`int64`|
|current_location_id|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|status_message|`utf8`|
|redis_configs|`json`|
|tier|`utf8`|
|memory_size_gb|`int64`|
|authorized_network|`utf8`|
|persistence_iam_identity|`utf8`|
|connect_mode|`utf8`|
|auth_enabled|`bool`|
|server_ca_certs|`json`|
|transit_encryption_mode|`utf8`|
|maintenance_policy|`json`|
|maintenance_schedule|`json`|
|replica_count|`int64`|
|nodes|`json`|
|read_endpoint|`utf8`|
|read_endpoint_port|`int64`|
|read_replicas_mode|`utf8`|