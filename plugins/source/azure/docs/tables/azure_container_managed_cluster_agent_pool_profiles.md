
# Table: azure_container_managed_cluster_agent_pool_profiles
ManagedClusterAgentPoolProfile profile for the container service agent pool
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|managed_cluster_cq_id|uuid|Unique CloudQuery ID of azure_container_managed_clusters table (FK)|
|name|text|Unique name of the agent pool profile in the context of the subscription and resource group|
|count|integer|Number of agents (VMs) to host docker containers.|
|vm_size|text|Size of agent VMs|
|os_disk_size_gb|integer|OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool.|
|os_disk_type|text|OS disk type to be used for machines in a given agent pool.|
|kubelet_disk_type|text|KubeletDiskType determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.|
|vnet_subnet_id|text|VNet SubnetID specifies the VNet's subnet identifier for nodes and maybe pods|
|pod_subnet_id|text|Pod SubnetID specifies the VNet's subnet identifier for pods|
|max_pods|integer|Maximum number of pods that can run on a node|
|os_type|text|OsType to be used to specify os type.|
|os_sku|text|OsSKU to be used to specify os sku.|
|max_count|integer|Maximum number of nodes for auto-scaling|
|min_count|integer|Minimum number of nodes for auto-scaling|
|enable_auto_scaling|boolean|Whether to enable auto-scaler|
|type|text|AgentPoolType represents types of an agent pool.|
|mode|text|AgentPoolMode represents mode of an agent pool.|
|orchestrator_version|text|Version of orchestrator specified when creating the managed cluster|
|node_image_version|text|Version of node image|
|upgrade_settings_max_surge|text|Count or percentage of additional nodes to be added during upgrade.|
|provisioning_state|text|The current deployment or provisioning state, which only appears in the response|
|power_state_code|text|Tells whether the cluster is Running or Stopped.|
|availability_zones|text[]|Availability zones for nodes.|
|enable_node_public_ip|boolean|Enable public IP for nodes|
|node_public_ip_prefix_id|text|Public IP Prefix ID VM nodes use IPs assigned from this Public IP Prefix|
|scale_set_priority|text|ScaleSetPriority to be used to specify virtual machine scale set priority.|
|scale_set_eviction_policy|text|ScaleSetEvictionPolicy to be used to specify eviction policy for Spot virtual machine scale set.|
|spot_max_price|float|SpotMaxPrice to be used to specify the maximum price you are willing to pay in US Dollars.|
|tags|jsonb|Agent pool tags to be persisted on the agent pool virtual machine scale set|
|node_labels|jsonb|Agent pool node labels to be persisted across all nodes in agent pool|
|node_taints|text[]|Taints added to new nodes during node pool create and scale.|
|proximity_placement_group_id|text|The ID for Proximity Placement Group|
|kubelet_config_cpu_manager_policy|text|CPU Manager policy to use|
|kubelet_config_cpu_cfs_quota|boolean|Enable CPU CFS quota enforcement for containers that specify CPU limits|
|kubelet_config_cpu_cfs_quota_period|text|Sets CPU CFS quota period value|
|kubelet_config_image_gc_high_threshold|integer|The percent of disk usage after which image garbage collection is always run|
|kubelet_config_image_gc_low_threshold|integer|The percent of disk usage before which image garbage collection is never run|
|kubelet_config_topology_manager_policy|text|Topology Manager policy to use|
|kubelet_config_allowed_unsafe_sysctls|text[]|Allowlist of unsafe sysctls or unsafe sysctl patterns (ending in `*`)|
|kubelet_config_fail_swap_on|boolean|If set to true it will make the Kubelet fail to start if swap is enabled on the node|
|kubelet_config_container_log_max_size_mb|integer|The maximum size (eg 10Mi) of container log file before it is rotated|
|kubelet_config_container_log_max_files|integer|The maximum number of container log files that can be present for a container.|
|kubelet_config_pod_max_pids|integer|The maximum number of processes per pod|
|linux_os_config|jsonb|LinuxOSConfig specifies the OS configuration of linux agent nodes|
|enable_encryption_at_host|boolean|Whether to enable EncryptionAtHost|
|enable_fips|boolean|Whether to use FIPS enabled OS|
|gpu_instance_profile|text|GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.|
