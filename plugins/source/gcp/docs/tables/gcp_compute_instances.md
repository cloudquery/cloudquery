# Table: gcp_compute_instances



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
|advanced_machine_features|JSON|
|can_ip_forward|Bool|
|confidential_instance_config|JSON|
|cpu_platform|String|
|creation_timestamp|String|
|deletion_protection|Bool|
|description|String|
|disks|JSON|
|display_device|JSON|
|fingerprint|String|
|guest_accelerators|JSON|
|hostname|String|
|id|Int|
|key_revocation_action_type|String|
|kind|String|
|label_fingerprint|String|
|labels|JSON|
|last_start_timestamp|String|
|last_stop_timestamp|String|
|last_suspended_timestamp|String|
|machine_type|String|
|metadata|JSON|
|min_cpu_platform|String|
|name|String|
|network_interfaces|JSON|
|network_performance_config|JSON|
|params|JSON|
|private_ipv6_google_access|String|
|reservation_affinity|JSON|
|resource_policies|StringArray|
|resource_status|JSON|
|satisfies_pzs|Bool|
|scheduling|JSON|
|service_accounts|JSON|
|shielded_instance_config|JSON|
|shielded_instance_integrity_policy|JSON|
|source_machine_image|String|
|source_machine_image_encryption_key|JSON|
|start_restricted|Bool|
|status|String|
|status_message|String|
|tags|JSON|
|zone|String|