# Table: gcp_container_clusters



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
|addons_config|JSON|
|authenticator_groups_config|JSON|
|autopilot|JSON|
|autoscaling|JSON|
|binary_authorization|JSON|
|cluster_ipv4_cidr|String|
|conditions|JSON|
|confidential_nodes|JSON|
|cost_management_config|JSON|
|create_time|String|
|current_master_version|String|
|current_node_count|Int|
|current_node_version|String|
|database_encryption|JSON|
|default_max_pods_constraint|JSON|
|description|String|
|enable_kubernetes_alpha|Bool|
|enable_tpu|Bool|
|endpoint|String|
|expire_time|String|
|id|String|
|identity_service_config|JSON|
|initial_cluster_version|String|
|initial_node_count|Int|
|instance_group_urls|StringArray|
|ip_allocation_policy|JSON|
|label_fingerprint|String|
|legacy_abac|JSON|
|location|String|
|locations|StringArray|
|logging_config|JSON|
|logging_service|String|
|maintenance_policy|JSON|
|master_auth|JSON|
|master_authorized_networks_config|JSON|
|mesh_certificates|JSON|
|monitoring_config|JSON|
|monitoring_service|String|
|name|String|
|network|String|
|network_config|JSON|
|network_policy|JSON|
|node_config|JSON|
|node_ipv4_cidr_size|Int|
|node_pool_auto_config|JSON|
|node_pool_defaults|JSON|
|node_pools|JSON|
|notification_config|JSON|
|private_cluster_config|JSON|
|release_channel|JSON|
|resource_labels|JSON|
|resource_usage_export_config|JSON|
|services_ipv4_cidr|String|
|shielded_nodes|JSON|
|status|String|
|status_message|String|
|subnetwork|String|
|tpu_ipv4_cidr_block|String|
|vertical_pod_autoscaling|JSON|
|workload_identity_config|JSON|
|zone|String|