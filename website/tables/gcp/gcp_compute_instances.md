# Table: gcp_compute_instances

This table shows data for GCP Compute Instances.

https://cloud.google.com/compute/docs/reference/rest/v1/instances#Instance

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id|utf8|
|advanced_machine_features|json|
|can_ip_forward|bool|
|confidential_instance_config|json|
|cpu_platform|utf8|
|creation_timestamp|utf8|
|deletion_protection|bool|
|description|utf8|
|disks|json|
|display_device|json|
|fingerprint|utf8|
|guest_accelerators|json|
|hostname|utf8|
|id|int64|
|key_revocation_action_type|utf8|
|kind|utf8|
|label_fingerprint|utf8|
|labels|json|
|last_start_timestamp|utf8|
|last_stop_timestamp|utf8|
|last_suspended_timestamp|utf8|
|machine_type|utf8|
|metadata|json|
|min_cpu_platform|utf8|
|name|utf8|
|network_interfaces|json|
|network_performance_config|json|
|params|json|
|private_ipv6_google_access|utf8|
|reservation_affinity|json|
|resource_policies|list<item: utf8, nullable>|
|resource_status|json|
|satisfies_pzs|bool|
|scheduling|json|
|self_link (PK)|utf8|
|service_accounts|json|
|shielded_instance_config|json|
|shielded_instance_integrity_policy|json|
|source_machine_image|utf8|
|source_machine_image_encryption_key|json|
|start_restricted|bool|
|status|utf8|
|status_message|utf8|
|tags|json|
|zone|utf8|