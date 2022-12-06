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
|name|String|
|description|String|
|initial_node_count|Int|
|node_config|JSON|
|master_auth|JSON|
|logging_service|String|
|monitoring_service|String|
|network|String|
|cluster_ipv4_cidr|String|
|addons_config|JSON|
|subnetwork|String|
|node_pools|JSON|
|locations|StringArray|
|enable_kubernetes_alpha|Bool|
|resource_labels|JSON|
|label_fingerprint|String|
|legacy_abac|JSON|
|network_policy|JSON|
|ip_allocation_policy|JSON|
|master_authorized_networks_config|JSON|
|maintenance_policy|JSON|
|binary_authorization|JSON|
|autoscaling|JSON|
|network_config|JSON|
|default_max_pods_constraint|JSON|
|resource_usage_export_config|JSON|
|authenticator_groups_config|JSON|
|private_cluster_config|JSON|
|database_encryption|JSON|
|vertical_pod_autoscaling|JSON|
|shielded_nodes|JSON|
|release_channel|JSON|
|workload_identity_config|JSON|
|mesh_certificates|JSON|
|cost_management_config|JSON|
|notification_config|JSON|
|confidential_nodes|JSON|
|identity_service_config|JSON|
|zone|String|
|endpoint|String|
|initial_cluster_version|String|
|current_master_version|String|
|current_node_version|String|
|create_time|String|
|status|String|
|status_message|String|
|node_ipv4_cidr_size|Int|
|services_ipv4_cidr|String|
|instance_group_urls|StringArray|
|current_node_count|Int|
|expire_time|String|
|location|String|
|enable_tpu|Bool|
|tpu_ipv4_cidr_block|String|
|conditions|JSON|
|autopilot|JSON|
|id|String|
|node_pool_defaults|JSON|
|logging_config|JSON|
|monitoring_config|JSON|
|node_pool_auto_config|JSON|