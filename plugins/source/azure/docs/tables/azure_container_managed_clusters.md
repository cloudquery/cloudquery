# Table: azure_container_managed_clusters

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2#ManagedCluster

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|extended_location|JSON|
|identity|JSON|
|aad_profile|JSON|
|api_server_access_profile|JSON|
|addon_profiles|JSON|
|agent_pool_profiles|JSON|
|auto_scaler_profile|JSON|
|auto_upgrade_profile|JSON|
|dns_prefix|String|
|disable_local_accounts|Bool|
|disk_encryption_set_id|String|
|enable_pod_security_policy|Bool|
|enable_rbac|Bool|
|fqdn_subdomain|String|
|http_proxy_config|JSON|
|identity_profile|JSON|
|kubernetes_version|String|
|linux_profile|JSON|
|network_profile|JSON|
|node_resource_group|String|
|oidc_issuer_profile|JSON|
|pod_identity_profile|JSON|
|private_link_resources|JSON|
|public_network_access|String|
|security_profile|JSON|
|service_principal_profile|JSON|
|storage_profile|JSON|
|windows_profile|JSON|
|azure_portal_fqdn|String|
|current_kubernetes_version|String|
|fqdn|String|
|max_agent_pools|Int|
|power_state|JSON|
|private_fqdn|String|
|provisioning_state|String|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|