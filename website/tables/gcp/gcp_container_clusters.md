# Table: gcp_container_clusters

This table shows data for GCP Container Clusters.

https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|initial_node_count|`int64`|
|node_config|`json`|
|master_auth|`json`|
|logging_service|`utf8`|
|monitoring_service|`utf8`|
|network|`utf8`|
|cluster_ipv4_cidr|`utf8`|
|addons_config|`json`|
|subnetwork|`utf8`|
|node_pools|`json`|
|locations|`list<item: utf8, nullable>`|
|enable_kubernetes_alpha|`bool`|
|resource_labels|`json`|
|label_fingerprint|`utf8`|
|legacy_abac|`json`|
|network_policy|`json`|
|ip_allocation_policy|`json`|
|master_authorized_networks_config|`json`|
|maintenance_policy|`json`|
|binary_authorization|`json`|
|autoscaling|`json`|
|network_config|`json`|
|default_max_pods_constraint|`json`|
|resource_usage_export_config|`json`|
|authenticator_groups_config|`json`|
|private_cluster_config|`json`|
|database_encryption|`json`|
|vertical_pod_autoscaling|`json`|
|shielded_nodes|`json`|
|release_channel|`json`|
|workload_identity_config|`json`|
|mesh_certificates|`json`|
|cost_management_config|`json`|
|notification_config|`json`|
|confidential_nodes|`json`|
|identity_service_config|`json`|
|self_link (PK)|`utf8`|
|zone|`utf8`|
|endpoint|`utf8`|
|initial_cluster_version|`utf8`|
|current_master_version|`utf8`|
|current_node_version|`utf8`|
|create_time|`utf8`|
|status|`utf8`|
|status_message|`utf8`|
|node_ipv4_cidr_size|`int64`|
|services_ipv4_cidr|`utf8`|
|instance_group_urls|`list<item: utf8, nullable>`|
|current_node_count|`int64`|
|expire_time|`utf8`|
|location|`utf8`|
|enable_tpu|`bool`|
|tpu_ipv4_cidr_block|`utf8`|
|conditions|`json`|
|autopilot|`json`|
|id|`utf8`|
|node_pool_defaults|`json`|
|logging_config|`json`|
|monitoring_config|`json`|
|node_pool_auto_config|`json`|
|etag|`utf8`|
|fleet|`json`|