# Table: gcp_sql_instances

This table shows data for GCP SQL Instances.

https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/instances#DatabaseInstance

The primary key for this table is **self_link**.

## Relations

The following tables depend on gcp_sql_instances:
  - [gcp_sql_users](gcp_sql_users)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|available_maintenance_versions|`list<item: utf8, nullable>`|
|backend_type|`utf8`|
|connection_name|`utf8`|
|create_time|`utf8`|
|current_disk_size|`int64`|
|database_installed_version|`utf8`|
|database_version|`utf8`|
|disk_encryption_configuration|`json`|
|disk_encryption_status|`json`|
|etag|`utf8`|
|failover_replica|`json`|
|gce_zone|`utf8`|
|instance_type|`utf8`|
|ip_addresses|`json`|
|ipv6_address|`utf8`|
|kind|`utf8`|
|maintenance_version|`utf8`|
|master_instance_name|`utf8`|
|max_disk_size|`int64`|
|name|`utf8`|
|on_premises_configuration|`json`|
|out_of_disk_report|`json`|
|project|`utf8`|
|region|`utf8`|
|replica_configuration|`json`|
|replica_names|`list<item: utf8, nullable>`|
|root_password|`utf8`|
|satisfies_pzs|`bool`|
|scheduled_maintenance|`json`|
|secondary_gce_zone|`utf8`|
|self_link (PK)|`utf8`|
|server_ca_cert|`json`|
|service_account_email_address|`utf8`|
|settings|`json`|
|state|`utf8`|
|suspension_reason|`list<item: utf8, nullable>`|