# Table: azure_web_apps

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web#Site

The primary key for this table is **id**.

## Relations

The following tables depend on azure_web_apps:
  - [azure_web_site_auth_settings](azure_web_site_auth_settings.md)
  - [azure_web_vnet_connections](azure_web_vnet_connections.md)
  - [azure_web_publishing_profiles](azure_web_publishing_profiles.md)
  - [azure_web_site_auth_settings_v2](azure_web_site_auth_settings_v2.md)
  - [azure_web_functions](azure_web_functions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|state|String|
|host_names|StringArray|
|repository_site_name|String|
|usage_state|String|
|enabled|Bool|
|enabled_host_names|StringArray|
|availability_state|String|
|host_name_ssl_states|JSON|
|server_farm_id|String|
|reserved|Bool|
|is_xenon|Bool|
|hyper_v|Bool|
|last_modified_time_utc|Timestamp|
|site_config|JSON|
|traffic_manager_host_names|StringArray|
|scm_site_also_stopped|Bool|
|target_swap_slot|String|
|hosting_environment_profile|JSON|
|client_affinity_enabled|Bool|
|client_cert_enabled|Bool|
|client_cert_mode|String|
|client_cert_exclusion_paths|String|
|host_names_disabled|Bool|
|custom_domain_verification_id|String|
|outbound_ip_addresses|String|
|possible_outbound_ip_addresses|String|
|container_size|Int|
|daily_memory_time_quota|Int|
|suspended_till|Timestamp|
|max_number_of_workers|Int|
|cloning_info|JSON|
|resource_group|String|
|is_default_container|Bool|
|default_host_name|String|
|slot_swap_status|JSON|
|https_only|Bool|
|redundancy_mode|String|
|in_progress_operation_id|UUID|
|storage_account_required|Bool|
|key_vault_reference_identity|String|
|virtual_network_subnet_id|String|
|identity|JSON|
|id (PK)|String|
|name|String|
|kind|String|
|location|String|
|type|String|
|tags|JSON|