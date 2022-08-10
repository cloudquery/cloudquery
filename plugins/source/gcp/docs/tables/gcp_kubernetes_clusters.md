
# Table: gcp_kubernetes_clusters
A Google Kubernetes Engine cluster
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|addons_config_cloud_run_config_disabled|boolean|Whether Cloud Run addon is enabled for this cluster|
|addons_config_cloud_run_config_load_balancer_type|text|"LOAD_BALANCER_TYPE_UNSPECIFIED" - Load balancer type for Cloud Run is unspecified   "LOAD_BALANCER_TYPE_EXTERNAL" - Install external load balancer for Cloud Run   "LOAD_BALANCER_TYPE_INTERNAL" - Install internal load balancer for Cloud Run|
|addons_config_config_connector_config_enabled|boolean|Whether Cloud Connector is enabled for this cluster|
|addons_config_dns_cache_config_enabled|boolean|Whether NodeLocal DNSCache is enabled for this cluster|
|addons_config_gce_persistent_disk_csi_driver_config_enabled|boolean|Whether the Compute Engine PD CSI driver is enabled for this cluster|
|addons_config_horizontal_pod_autoscaling_disabled|boolean|Whether the Horizontal Pod Autoscaling feature is enabled in the cluster|
|addons_config_http_load_balancing_disabled|boolean|Whether the HTTP Load Balancing controller is enabled in the cluster|
|addons_config_kubernetes_dashboard_disabled|boolean|Whether the Kubernetes Dashboard is enabled for this cluster|
|addons_config_network_policy_config_disabled|boolean|Whether NetworkPolicy is enabled for this cluster|
|authenticator_groups_config_enabled|boolean|Whether this cluster should return group membership lookups during authentication using a group of security groups|
|authenticator_groups_config_security_group|text|The name of the security group-of-groups to be used Only relevant if enabled = true|
|autopilot_enabled|boolean|Enable Autopilot|
|autoscaling_autoprovisioning_locations|text[]|The list of Google Compute Engine zones (https://cloudgooglecom/compute/docs/zones#available) in which the NodePool's nodes can be created by NAP|
|autoscaling_autoprovisioning_node_pool_defaults|jsonb|AutoprovisioningNodePoolDefaults contains defaults for a node pool created by NAP|
|autoscaling_profile|text|"PROFILE_UNSPECIFIED" - No change to autoscaling configuration   "OPTIMIZE_UTILIZATION" - Prioritize optimizing utilization of resources   "BALANCED" - Use default (balanced) autoscaling configuration|
|autoscaling_enable_node_autoprovisioning|boolean|Enables automatic node pool creation and deletion|
|autoscaling_resource_limits|jsonb|Contains global constraints regarding minimum and maximum amount of resources in the cluster|
|binary_authorization_enabled|boolean|Enable Binary Authorization for this cluster|
|cluster_ipv4_cidr|cidr|The IP address range of the container pods in this cluster, in CIDR (http://enwikipediaorg/wiki/Classless_Inter-Domain_Routing) notation (eg|
|conditions|jsonb|Which conditions caused the current cluster state|
|confidential_nodes_enabled|boolean|Whether Confidential Nodes feature is enabled for all nodes in this cluster|
|create_time|timestamp without time zone|The time the cluster was created, in RFC3339 (https://wwwietforg/rfc/rfc3339txt) text format|
|current_master_version|text|The current software version of the master endpoint|
|current_node_count|bigint|The number of nodes currently in the cluster|
|current_node_version|text|Deprecated, use NodePoolsversion (https://cloudgooglecom/kubernetes-engine/docs/reference/rest/v1/projectslocationsclustersnodePools) instead|
|database_encryption_key_name|text|Name of CloudKMS key to use for the encryption of secrets in etcd|
|database_encryption_state|text|"UNKNOWN" - Should never be set   "ENCRYPTED" - Secrets in etcd are encrypted   "DECRYPTED" - Secrets in etcd are stored in plain text (at etcd level) - this is unrelated to Compute Engine level full disk encryption|
|default_max_pods_constraint_max_pods_per_node|bigint|Constraint enforced on the max num of pods per node|
|description|text|An optional description of this cluster|
|enable_kubernetes_alpha|boolean|Kubernetes alpha features are enabled on this cluster|
|enable_tpu|boolean|Enable the ability to use Cloud TPUs in this cluster|
|endpoint|inet|The IP address of this cluster's master endpoint|
|expire_time|timestamp without time zone|The time the cluster will be automatically deleted in RFC3339 (https://wwwietforg/rfc/rfc3339txt) text format|
|id|text|Output only|
|initial_cluster_version|text|The initial Kubernetes version for this cluster|
|initial_node_count|bigint|The number of nodes to create in this cluster|
|instance_group_urls|text[]|Deprecated|
|ip_allocation_policy_cluster_ipv4_cidr|cidr|This field is deprecated, use cluster_ipv4_cidr_block|
|ip_allocation_policy_cluster_ipv4_cidr_block|cidr|The IP address range for the cluster pod IPs If this field is set, then `clustercluster_ipv4_cidr` must be left blank|
|ip_allocation_policy_cluster_secondary_range_name|text|The name of the secondary range to be used for the cluster CIDR block|
|ip_allocation_policy_create_subnetwork|boolean|Whether a new subnetwork will be created automatically for the cluster|
|ip_allocation_policy_node_ipv4_cidr|cidr|This field is deprecated, use node_ipv4_cidr_block|
|ip_allocation_policy_node_ipv4_cidr_block|cidr|The IP address range of the instance IPs in this cluster|
|ip_allocation_policy_services_ipv4_cidr|cidr|This field is deprecated, use services_ipv4_cidr_block|
|ip_allocation_policy_services_ipv4_cidr_block|cidr|The IP address range of the services IPs in this cluster|
|ip_allocation_policy_services_secondary_range_name|text|The name of the secondary range to be used as for the services CIDR block|
|ip_allocation_policy_subnetwork_name|text|A custom subnetwork name to be used if `create_subnetwork` is true|
|ip_allocation_policy_tpu_ipv4_cidr_block|cidr|The IP address range of the Cloud TPUs in this cluster|
|ip_allocation_policy_use_ip_aliases|boolean|Whether alias IPs will be used for pod IPs in the cluster|
|ip_allocation_policy_use_routes|boolean|Whether routes will be used for pod IPs in the cluster This is used in conjunction with use_ip_aliases|
|label_fingerprint|text|The fingerprint of the set of labels for this cluster|
|legacy_abac_enabled|boolean|Whether the ABAC authorizer is enabled for this cluster When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM|
|location|text|The name of the Google Compute Engine zone (https://cloudgooglecom/compute/docs/regions-zones/regions-zones#available) or region (https://cloudgooglecom/compute/docs/regions-zones/regions-zones#available) in which the cluster resides|
|locations|text[]|The list of Google Compute Engine zones (https://cloudgooglecom/compute/docs/zones#available) in which the cluster's nodes should be located|
|logging_config_component_config_enable_components|text[]|Select components to collect logs|
|logging_service|text|The logging service the cluster should use to write logs|
|maintenance_policy_resource_version|text|A hash identifying the version of this policy, so that updates to fields of the policy won't accidentally undo intermediate changes (and so that users of the API unaware of some fields won't accidentally remove other fields)|
|maintenance_policy_window_daily_maintenance_window_duration|text|Duration of the time window, automatically chosen to be smallest possible in the given scenario|
|maintenance_policy_window_daily_maintenance_window_start_time|timestamp without time zone|Time within the maintenance window to start the maintenance operations|
|maintenance_policy_window_maintenance_exclusions|jsonb|Exceptions to maintenance window Non-emergency maintenance should not occur in these windows|
|maintenance_policy_window_recurring_window_recurrence|text|An RRULE (https://toolsietforg/html/rfc5545#section-3853) for how this window reccurs|
|maintenance_policy_window_recurring_window_window_end_time|timestamp without time zone|The time that the window ends|
|maintenance_policy_window_recurring_window_window_start_time|timestamp without time zone|The time that the window first starts|
|master_auth_client_certificate|text|Base64-encoded public certificate used by clients to authenticate to the cluster endpoint|
|master_auth_client_certificate_config_issue_client_certificate|boolean|Issue a client certificate|
|master_auth_client_key|text|Base64-encoded private key used by clients to authenticate to the cluster endpoint|
|master_auth_cluster_ca_certificate|text|Base64-encoded public certificate that is the root of trust for the cluster|
|master_auth_password|text|The password to use for HTTP basic authentication to the master endpoint|
|master_auth_username|text|The username to use for HTTP basic authentication to the master endpoint|
|master_authorized_networks_config_cidr_blocks|jsonb|cidr_blocks define up to 50 external networks that could access Kubernetes master through HTTPS|
|master_authorized_networks_config_enabled|boolean|Whether or not master authorized networks is enabled|
|monitoring_config_component_config_enable_components|text[]|Select components to collect metrics|
|monitoring_service|text|The monitoring service the cluster should use to write metrics|
|name|text|The name of this cluster|
|network|text|The name of the Google Compute Engine network (https://cloudgooglecom/compute/docs/networks-and-firewalls#networks) to which the cluster is connected|
|network_config_datapath_provider|text|The desired datapath provider for this cluster|
|network_config_default_snat_status_disabled|boolean|Disables cluster default sNAT rules|
|network_config_enable_intra_node_visibility|boolean|Whether Intra-node visibility is enabled for this cluster|
|network_config_enable_l4ilb_subsetting|boolean|Whether L4ILB Subsetting is enabled for this cluster|
|network_config_network|text|Output only|
|network_config_private_ipv6_google_access|text|The desired state of IPv6 connectivity to Google Services|
|network_config_subnetwork|text|Output only|
|network_policy_enabled|boolean|Whether network policy is enabled on the cluster|
|network_policy_provider|text|"PROVIDER_UNSPECIFIED" - Not set   "CALICO" - Tigera (Calico Felix)|
|node_config|jsonb|Parameters used in creating the cluster's nodes|
|node_ipv4_cidr_size|bigint|The size of the address space on each node for hosting containers|
|notification_config_pubsub_enabled|boolean|Enable notifications for Pub/Sub|
|notification_config_pubsub_topic|text|The desired Pub/Sub topic to which notifications will be sent by GKE|
|private_cluster_config_enable_private_endpoint|boolean|Whether the master's internal IP address is used as the cluster endpoint|
|private_cluster_config_enable_private_nodes|boolean|Whether nodes have internal IP addresses only|
|private_cluster_config_master_global_access_config_enabled|boolean|Whenever master is accessible globally or not|
|private_cluster_config_master_ipv4_cidr_block|cidr|The IP range in CIDR notation to use for the hosted master network|
|private_cluster_config_peering_name|text|Output only|
|private_cluster_config_private_endpoint|text|Output only|
|private_cluster_config_public_endpoint|text|Output only|
|release_channel|text|"UNSPECIFIED" - No channel specified   "RAPID" - RAPID channel is offered on an early access basis for customers who want to test new releases|
|resource_labels|jsonb|The resource labels for the cluster to use to annotate any related Google Compute Engine resources|
|resource_usage_export_config|jsonb|Configuration for exporting resource usages|
|self_link|text|Server-defined URL for the resource|
|services_ipv4_cidr|cidr|The IP address range of the Kubernetes services in this cluster, in CIDR (http://enwikipediaorg/wiki/Classless_Inter-Domain_Routing) notation (eg|
|shielded_nodes_enabled|boolean|Whether Shielded Nodes features are enabled on all nodes in this cluster|
|status|text|"STATUS_UNSPECIFIED" - Not set   "PROVISIONING" - The PROVISIONING state indicates the cluster is being created   "RUNNING" - The RUNNING state indicates the cluster has been created and is fully usable   "RECONCILING" - The RECONCILING state indicates that some work is actively being done on the cluster, such as upgrading the master or node software|
|status_message|text|Deprecated|
|subnetwork|text|The name of the Google Compute Engine subnetwork (https://cloudgooglecom/compute/docs/subnetworks) to which the cluster is connected|
|tpu_ipv4_cidr_block|cidr|The IP address range of the Cloud TPUs in this cluster, in CIDR (http://enwikipediaorg/wiki/Classless_Inter-Domain_Routing) notation (eg|
|vertical_pod_autoscaling_enabled|boolean|Enables vertical pod autoscaling|
|workload_identity_config_workload_pool|text|The workload pool to attach all Kubernetes service accounts to|
|zone|text|The name of the Google Compute Engine zone (https://cloudgooglecom/compute/docs/zones#available) in which the cluster resides|
