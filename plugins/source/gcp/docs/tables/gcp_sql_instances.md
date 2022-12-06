# Table: gcp_sql_instances



The primary key for this table is **self_link**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|self_link (PK)|String|
|available_maintenance_versions|StringArray|
|backend_type|String|
|connection_name|String|
|create_time|String|
|current_disk_size|Int|
|database_installed_version|String|
|database_version|String|
|disk_encryption_configuration|JSON|
|disk_encryption_status|JSON|
|etag|String|
|failover_replica|JSON|
|gce_zone|String|
|instance_type|String|
|ip_addresses|JSON|
|ipv6_address|String|
|kind|String|
|maintenance_version|String|
|master_instance_name|String|
|max_disk_size|Int|
|name|String|
|on_premises_configuration|JSON|
|out_of_disk_report|JSON|
|project|String|
|region|String|
|replica_configuration|JSON|
|replica_names|StringArray|
|root_password|String|
|satisfies_pzs|Bool|
|scheduled_maintenance|JSON|
|secondary_gce_zone|String|
|server_ca_cert|JSON|
|service_account_email_address|String|
|settings|JSON|
|state|String|
|suspension_reason|StringArray|