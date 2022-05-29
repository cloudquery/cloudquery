
# Table: gcp_kubernetes_cluster_node_pools
NodePool contains the name and configuration for a cluster's node pool
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of gcp_kubernetes_clusters table (FK)|
|autoscaling_autoprovisioned|boolean|Can this node pool be deleted automatically|
|autoscaling_enabled|boolean|Is autoscaling enabled for this node pool|
|autoscaling_max_node_count|bigint|Maximum number of nodes in the NodePool|
|autoscaling_min_node_count|bigint|Minimum number of nodes in the NodePool|
|conditions|jsonb|Which conditions caused the current node pool state|
|config_accelerators|jsonb|A list of hardware accelerators to be attached to each node|
|config_boot_disk_kms_key|text|The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool|
|config_disk_size_gb|bigint|Size of the disk attached to each node, specified in GB The smallest allowed disk size is 10GB|
|config_disk_type|text|Type of the disk attached to each node (eg|
|config_gvnic_enabled|boolean|Whether gVNIC features are enabled in the node pool|
|config_image_type|text|The image type to use for this node|
|config_kubelet_config_cpu_cfs_quota|boolean|Enable CPU CFS quota enforcement for containers that specify CPU limits|
|config_kubelet_config_cpu_cfs_quota_period|text|Set the CPU CFS quota period value 'cpucfs_period_us'|
|config_kubelet_config_cpu_manager_policy|text|Control the CPU management policy on the node|
|config_labels|jsonb|The map of Kubernetes labels (key/value pairs) to be applied to each node|
|config_linux_node_config_sysctls|jsonb|The Linux kernel parameters to be applied to the nodes and all pods running on the nodes|
|config_local_ssd_count|bigint|The number of local SSD disks to be attached to the node|
|config_machine_type|text|The name of a Google Compute Engine machine type (https://cloudgooglecom/compute/docs/machine-types) If unspecified, the default machine type is `e2-medium`|
|config_metadata|jsonb|The metadata key/value pairs assigned to instances in the cluster|
|config_min_cpu_platform|text|Minimum CPU platform to be used by this instance|
|config_node_group|text|Setting this field will assign instances of this pool to run on the specified node group|
|config_oauth_scopes|text[]|The set of Google API scopes to be made available on all of the node VMs under the "default" service account|
|config_preemptible|boolean|Whether the nodes are created as preemptible VM instances|
|config_reservation_affinity_consume_reservation_type|text|"UNSPECIFIED" - Default value|
|config_reservation_affinity_key|text|Corresponds to the label key of a reservation resource|
|config_reservation_affinity_values|text[]|Corresponds to the label value(s) of reservation resource(s)|
|config_sandbox_config_type|text|"UNSPECIFIED" - Default value|
|config_service_account|text|The Google Cloud Platform Service Account to be used by the node VMs|
|config_shielded_instance_config_enable_integrity_monitoring|boolean|Defines whether the instance has integrity monitoring enabled|
|config_shielded_instance_config_enable_secure_boot|boolean|Defines whether the instance has Secure Boot enabled|
|config_tags|text[]|The list of instance tags applied to all nodes|
|config_taints|jsonb|List of kubernetes taints to be applied to each node|
|config_workload_metadata_config_mode|text|"MODE_UNSPECIFIED" - Not set   "GCE_METADATA" - Expose all Compute Engine metadata to pods   "GKE_METADATA" - Run the GKE Metadata Server on this node|
|initial_node_count|bigint|The initial node count for the pool|
|instance_group_urls|text[]|The resource URLs of the managed instance groups (https://cloudgooglecom/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool|
|locations|text[]|The list of Google Compute Engine zones (https://cloudgooglecom/compute/docs/zones#available) in which the NodePool's nodes should be located|
|management_auto_repair|boolean|A flag that specifies whether the node auto-repair is enabled for the node pool|
|management_auto_upgrade|boolean|A flag that specifies whether node auto-upgrade is enabled for the node pool|
|management_upgrade_options_auto_upgrade_start_time|timestamp without time zone|This field is set when upgrades are about to commence with the approximate start time for the upgrades, in RFC3339 (https://wwwietforg/rfc/rfc3339txt) text format|
|management_upgrade_options_description|text|This field is set when upgrades are about to commence with the description of the upgrade|
|max_pods_constraint_max_pods_per_node|bigint|Constraint enforced on the max num of pods per node|
|name|text|The name of the node pool|
|network_config_create_pod_range|boolean|Input only|
|network_config_pod_ipv4_cidr_block|cidr|The IP address range for pod IPs in this node pool Only applicable if `create_pod_range` is true|
|network_config_pod_range|text|The ID of the secondary range for pod IPs|
|pod_ipv4_cidr_size|bigint|The pod CIDR block size per node in this node pool|
|self_link|text|Server-defined URL for the resource|
|status|text|"STATUS_UNSPECIFIED" - Not set   "PROVISIONING" - The PROVISIONING state indicates the node pool is being created   "RUNNING" - The RUNNING state indicates the node pool has been created and is fully usable   "RUNNING_WITH_ERROR" - The RUNNING_WITH_ERROR state indicates the node pool has been created and is partially usable|
|status_message|text|Deprecated|
|upgrade_settings_max_surge|bigint|The maximum number of nodes that can be created beyond the current size of the node pool during the upgrade process|
|upgrade_settings_max_unavailable|bigint|The maximum number of nodes that can be simultaneously unavailable during the upgrade process|
|version|text|The version of the Kubernetes of this node|
