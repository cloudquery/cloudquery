# Table: azure_container_managed_clusters

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice#ManagedCluster

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|provisioning_state|String|
|power_state|JSON|
|max_agent_pools|Int|
|kubernetes_version|String|
|dns_prefix|String|
|fqdn_subdomain|String|
|fqdn|String|
|private_fqdn|String|
|azure_portal_fqdn|String|
|agent_pool_profiles|JSON|
|linux_profile|JSON|
|windows_profile|JSON|
|service_principal_profile|JSON|
|addon_profiles|JSON|
|pod_identity_profile|JSON|
|node_resource_group|String|
|enable_rbac|Bool|
|enable_pod_security_policy|Bool|
|network_profile|JSON|
|aad_profile|JSON|
|auto_upgrade_profile|JSON|
|auto_scaler_profile|JSON|
|api_server_access_profile|JSON|
|disk_encryption_set_id|String|
|identity_profile|JSON|
|private_link_resources|JSON|
|disable_local_accounts|Bool|
|http_proxy_config|JSON|
|identity|JSON|
|sku|JSON|
|extended_location|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|