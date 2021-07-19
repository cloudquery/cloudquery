
# Table: azure_container_managed_clusters
ManagedCluster managed cluster
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|provisioning_state|text|The current deployment or provisioning state.|
|power_state_code|text|Tells whether the cluster is Running or Stopped.|
|max_agent_pools|integer|The max number of agent pools for the managed cluster|
|kubernetes_version|text|Version of Kubernetes specified when creating the managed cluster|
|dns_prefix|text|DNS prefix specified when creating the managed cluster|
|fqdn_subdomain|text|FQDN subdomain specified when creating private cluster with custom private dns zone|
|fqdn|text|FQDN for the master pool|
|private_fqdn|text|FQDN of private cluster|
|azure_portal_fqdn|text|FQDN for the master pool which used by proxy config|
|linux_profile_admin_username|text|The administrator username to use for Linux VMs|
|windows_profile_admin_username|text|Specifies the name of the administrator account.|
|windows_profile_admin_password|text|Specifies the password of the administrator account.|
|windows_profile_license_type|text|The licenseType to use for Windows VMs.|
|windows_profile_enable_csi_proxy|boolean|Whether to enable CSI proxy|
|service_principal_profile_client_id|text|The ID for the service principal|
|service_principal_profile_secret|text|The secret password associated with the service principal in plain text|
|addon_profiles|jsonb|Profile of managed cluster add-on|
|pod_identity_profile_enabled|boolean|Whether the pod identity addon is enabled|
|pod_identity_profile_allow_network_plugin_kubenet|boolean|Customer consent for enabling AAD pod identity addon in cluster using Kubenet network plugin|
|node_resource_group|text|Name of the resource group containing agent pool nodes|
|enable_rbac|boolean|Whether to enable Kubernetes Role-Based Access Control|
|network_profile_network_plugin|text|Network plugin used for building Kubernetes network.|
|network_profile_network_policy|text|Network policy used for building Kubernetes network.|
|network_profile_network_mode|text|Network mode used for building Kubernetes network.|
|network_profile_pod_cidr|text|A CIDR notation IP range from which to assign pod IPs when kubenet is used|
|network_profile_service_cidr|text|A CIDR notation IP range from which to assign service cluster IPs.|
|network_profile_dns_service_ip|text|An IP address assigned to the Kubernetes DNS service.|
|network_profile_docker_bridge_cidr|text|A CIDR notation IP range assigned to the Docker bridge network.|
|network_profile_outbound_type|text|The outbound (egress) routing method.|
|network_profile_load_balancer_sku|text|The load balancer sku for the managed cluster.|
|network_profile_load_balancer_managed_outbound_ips_count|integer|Desired number of outbound IP created/managed by Azure for the cluster load balancer.|
|network_profile_load_balancer_outbound_ip_prefixes|text[]|A list of public IP prefix resources|
|network_profile_load_balancer_outbound_ips|text[]|A list of public IP resources|
|network_profile_load_balancer_effective_outbound_ips|text[]|The effective outbound IP resources of the cluster load balancer|
|network_profile_load_balancer_allocated_outbound_ports|integer|Desired number of allocated SNAT ports per VM.|
|network_profile_load_balancer_idle_timeout|integer|Desired outbound flow idle timeout in minutes.|
|aad_profile_managed|boolean|Whether to enable managed AAD|
|aad_profile_enable_azure_rbac|boolean|Whether to enable Azure RBAC for Kubernetes authorization|
|aad_profile_admin_group_object_ids|text[]|AAD group object IDs that will have admin role of the cluster|
|aad_profile_client_app_id|text|The client AAD application ID|
|aad_profile_server_app_id|text|The server AAD application ID|
|aad_profile_server_app_secret|text|The server AAD application secret|
|aad_profile_tenant_id|text|The AAD tenant ID to use for authentication.|
|auto_upgrade_profile_upgrade_channel|text|upgrade channel for auto upgrade.|
|auto_scaler_profile_expander|text|Possible values include: 'LeastWaste', 'MostPods', 'Priority', 'Random'|
|api_server_access_profile_authorized_ip_ranges|text[]|Authorized IP Ranges to kubernetes API server|
|api_server_access_profile_enable_private_cluster|boolean|Whether to create the cluster as a private cluster or not|
|api_server_access_profile_private_dns_zone|text|Private dns zone mode for private cluster|
|disk_encryption_set_id|text|ResourceId of the disk encryption set to use for enabling encryption at rest|
|identity_profile|jsonb|Identities associated with the cluster|
|disable_local_accounts|boolean|If set to true, getting static credential will be disabled for this cluster.|
|http_proxy_config_http_proxy|text|HTTP proxy server endpoint to use|
|http_proxy_config_https_proxy|text|HTTPS proxy server endpoint to use|
|http_proxy_config_no_proxy|text[]|Endpoints that should not go through proxy|
|http_proxy_config_trusted_ca|text|Alternative CA cert to use for connecting to proxy servers|
|identity_principal_id|text|The principal id of the system assigned identity which is used by master components|
|identity_tenant_id|text|The tenant id of the system assigned identity which is used by master components|
|identity_type|text|The type of identity used for the managed cluster.|
|identity_user_assigned_identities|jsonb|The user identity associated with the managed cluster.|
|sku_name|text|Name of a managed cluster SKU.|
|sku_tier|text|Tier of a managed cluster SKU.|
|extended_location_name|text|The name of the extended location|
|extended_location_type|text|The type of the extended location.|
|id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
