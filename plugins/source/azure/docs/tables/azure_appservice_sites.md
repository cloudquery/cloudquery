# Table: azure_appservice_sites

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2#Site

The primary key for this table is **id**.

## Relations

The following tables depend on azure_appservice_sites:
  - [azure_appservice_functions](azure_appservice_functions.md)
  - [azure_appservice_site_auth_settings](azure_appservice_site_auth_settings.md)
  - [azure_appservice_site_auth_settings_v2](azure_appservice_site_auth_settings_v2.md)
  - [azure_appservice_vnet_connections](azure_appservice_vnet_connections.md)

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
|kind|String|
|client_affinity_enabled|Bool|
|client_cert_enabled|Bool|
|client_cert_exclusion_paths|String|
|client_cert_mode|String|
|cloning_info|JSON|
|container_size|Int|
|custom_domain_verification_id|String|
|daily_memory_time_quota|Int|
|enabled|Bool|
|https_only|Bool|
|host_name_ssl_states|JSON|
|host_names_disabled|Bool|
|hosting_environment_profile|JSON|
|hyper_v|Bool|
|is_xenon|Bool|
|key_vault_reference_identity|String|
|public_network_access|String|
|redundancy_mode|String|
|reserved|Bool|
|scm_site_also_stopped|Bool|
|server_farm_id|String|
|site_config|JSON|
|storage_account_required|Bool|
|virtual_network_subnet_id|String|
|vnet_content_share_enabled|Bool|
|vnet_image_pull_enabled|Bool|
|vnet_route_all_enabled|Bool|
|availability_state|String|
|default_host_name|String|
|enabled_host_names|StringArray|
|host_names|StringArray|
|in_progress_operation_id|String|
|is_default_container|Bool|
|last_modified_time_utc|Timestamp|
|max_number_of_workers|Int|
|outbound_ip_addresses|String|
|possible_outbound_ip_addresses|String|
|repository_site_name|String|
|resource_group|String|
|slot_swap_status|JSON|
|state|String|
|suspended_till|Timestamp|
|target_swap_slot|String|
|traffic_manager_host_names|StringArray|
|usage_state|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|