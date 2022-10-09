# Schema Changes from v0 to v1
This guide summarizes schema changes from CloudQuery v0 to v1. It is automatically generated and
not guaranteed to be complete, but we hope it helps as a starting point and reference when migrating to v1.

Last updated 2022-10-06.

## azure_account_location_paired_region
This table was removed.


## azure_account_locations
This table was removed.


## azure_authorization_role_assignments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|principal_id|text|removed|
|properties_principal_id|text|added|
|properties_role_definition_id|text|added|
|properties_scope|text|added|
|role_definition_id|text|removed|
|scope|text|removed|

## azure_authorization_role_definition_permissions
Moved to JSON column on [azure_authorization_role_definitions](#azure_authorization_role_definitions)


## azure_authorization_role_definitions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|permissions|jsonb|added|

## azure_batch_account_private_endpoint_connections
Moved to JSON column on [azure_batch_accounts](#azure_batch_accounts)


## azure_batch_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|active_job_and_job_schedule_quota|bigint|updated|Type changed from integer to bigint
|auto_storage|jsonb|added|
|auto_storage_authentication_mode|text|removed|
|auto_storage_last_key_sync_time|timestamp without time zone|removed|
|auto_storage_node_identity_reference_resource_id|text|removed|
|auto_storage_storage_account_id|text|removed|
|dedicated_core_quota|bigint|updated|Type changed from integer to bigint
|encryption|jsonb|added|
|encryption_key_source|text|removed|
|encryption_key_vault_properties_key_identifier|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|key_vault_reference|jsonb|added|
|key_vault_reference_id|text|removed|
|key_vault_reference_url|text|removed|
|low_priority_core_quota|bigint|updated|Type changed from integer to bigint
|pool_quota|bigint|updated|Type changed from integer to bigint
|private_endpoint_connections|jsonb|added|

## azure_cdn_custom_domains
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|cdn_endpoint_id|text|added|
|host_name|text|added|
|resource_state|text|added|
|custom_https_provisioning_state|text|added|
|custom_https_provisioning_substate|text|added|
|validation_data|text|added|
|provisioning_state|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|

## azure_cdn_endpoints
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|cdn_profile_id|text|added|
|host_name|text|added|
|origins|jsonb|added|
|origin_groups|jsonb|added|
|resource_state|text|added|
|provisioning_state|text|added|
|origin_path|text|added|
|content_types_to_compress|text[]|added|
|origin_host_header|text|added|
|is_compression_enabled|boolean|added|
|is_http_allowed|boolean|added|
|is_https_allowed|boolean|added|
|query_string_caching_behavior|text|added|
|optimization_type|text|added|
|probe_path|text|added|
|geo_filters|jsonb|added|
|default_origin_group|jsonb|added|
|url_signing_keys|jsonb|added|
|delivery_policy|jsonb|added|
|web_application_firewall_policy_link|jsonb|added|
|location|text|added|
|tags|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|

## azure_cdn_profile_endpoint_custom_domains
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_delivery_policy_rules
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_geo_filters
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_origin_groups
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_origins
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_routes
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_url_signing_keys
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoints
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_rule_set_rules
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_rule_sets
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_security_policies
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profiles

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|created_by|text|removed|
|created_by_type|text|removed|
|last_modified_at_time|timestamp without time zone|removed|
|last_modified_by|text|removed|
|last_modified_by_type|text|removed|
|resource_state|text|added|
|sku|jsonb|added|
|sku_name|text|removed|
|state|text|removed|
|system_data|jsonb|added|

## azure_cdn_routes
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|cdn_endpoint_id|text|added|
|custom_domains|jsonb|added|
|origin_group|jsonb|added|
|origin_path|text|added|
|rule_sets|jsonb|added|
|supported_protocols|text[]|added|
|patterns_to_match|text[]|added|
|query_string_caching_behavior|text|added|
|forwarding_protocol|text|added|
|link_to_default_domain|text|added|
|https_redirect|text|added|
|enabled_state|text|added|
|provisioning_state|text|added|
|deployment_status|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|

## azure_cdn_rule_sets
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|cdn_profile_id|text|added|
|provisioning_state|text|added|
|deployment_status|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|

## azure_cdn_rules
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|cdn_rule_set_id|text|added|
|order|bigint|added|
|conditions|jsonb|added|
|actions|jsonb|added|
|match_processing_behavior|text|added|
|provisioning_state|text|added|
|deployment_status|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|

## azure_cdn_security_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|cdn_profile_id|text|added|
|provisioning_state|text|added|
|deployment_status|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|

## azure_compute_disk_encryption_settings
Moved to JSON column on [azure_compute_disks](#azure_compute_disks)


## azure_compute_disks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|bursting_enabled|boolean|added|
|creation_data|jsonb|added|
|creation_data_create_option|text|removed|
|creation_data_gallery_image_reference_id|text|removed|
|creation_data_gallery_image_reference_lun|integer|removed|
|creation_data_image_reference_id|text|removed|
|creation_data_image_reference_lun|integer|removed|
|creation_data_source_resource_id|text|removed|
|creation_data_source_unique_id|text|removed|
|creation_data_source_uri|text|removed|
|creation_data_storage_account_id|text|removed|
|creation_data_upload_size_bytes|bigint|removed|
|disk_m_bps_read_only|bigint|added|
|disk_m_bps_read_write|bigint|added|
|disk_mbps_read_only|bigint|removed|
|disk_mbps_read_write|bigint|removed|
|disk_size_gb|bigint|updated|Type changed from integer to bigint
|encryption|jsonb|added|
|encryption_disk_encryption_set_id|text|removed|
|encryption_settings_collection|jsonb|added|
|encryption_settings_collection_enabled|boolean|removed|
|encryption_settings_collection_encryption_settings_version|text|removed|
|encryption_type|text|removed|
|extended_location|jsonb|added|
|hyper_v_generation|text|added|
|hyperv_generation|text|removed|
|max_shares|bigint|updated|Type changed from integer to bigint
|property_updates_in_progress|jsonb|added|
|purchase_plan|jsonb|added|
|security_profile|jsonb|added|
|share_info|jsonb|updated|Type changed from text[] to jsonb
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|
|supports_hibernation|boolean|added|
|tier|text|added|

## azure_compute_instance_views
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|compute_virtual_machine_id|text|added|
|platform_update_domain|bigint|added|
|platform_fault_domain|bigint|added|
|computer_name|text|added|
|os_name|text|added|
|os_version|text|added|
|hyper_v_generation|text|added|
|rdp_thumb_print|text|added|
|vm_agent|jsonb|added|
|maintenance_redeploy_status|jsonb|added|
|disks|jsonb|added|
|extensions|jsonb|added|
|vm_health|jsonb|added|
|boot_diagnostics|jsonb|added|
|assigned_host|text|added|
|statuses|jsonb|added|
|patch_status|jsonb|added|

## azure_compute_virtual_machine_extensions
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|compute_virtual_machine_id|text|added|
|force_update_tag|text|added|
|publisher|text|added|
|type_handler_version|text|added|
|auto_upgrade_minor_version|boolean|added|
|enable_automatic_upgrade|boolean|added|
|provisioning_state|text|added|
|instance_view|jsonb|added|
|id|text|added|
|name|text|added|
|location|text|added|
|tags|jsonb|added|
|type|text|added|

## azure_compute_virtual_machine_resources
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machine_scale_set_extensions
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machine_scale_set_os_profile_secrets
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machine_scale_sets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|additional_capabilities|jsonb|added|
|additional_capabilities_ultra_ssd_enabled|boolean|removed|
|automatic_repairs_policy|jsonb|added|
|automatic_repairs_policy_enabled|boolean|removed|
|automatic_repairs_policy_grace_period|text|removed|
|billing_profile_max_price|float|removed|
|diagnostics_profile|jsonb|removed|
|eviction_policy|text|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|extension_profile_extensions_time_budget|text|removed|
|host_group|jsonb|added|
|host_group_id|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|license_type|text|removed|
|network_profile|jsonb|removed|
|os_profile_admin_password|text|removed|
|os_profile_admin_username|text|removed|
|os_profile_computer_name_prefix|text|removed|
|os_profile_custom_data|text|removed|
|os_profile_linux_configuration|jsonb|removed|
|os_profile_windows_configuration|jsonb|removed|
|plan|jsonb|added|
|plan_name|text|removed|
|plan_product|text|removed|
|plan_promotion_code|text|removed|
|plan_publisher|text|removed|
|platform_fault_domain_count|bigint|updated|Type changed from integer to bigint
|priority|text|removed|
|proximity_placement_group|jsonb|added|
|proximity_placement_group_id|text|removed|
|scale_in_policy|jsonb|added|
|scale_in_policy_rules|text[]|removed|
|scheduled_events_profile|jsonb|removed|
|security_profile|jsonb|removed|
|sku|jsonb|added|
|sku_capacity|bigint|removed|
|sku_name|text|removed|
|sku_tier|text|removed|
|storage_profile|jsonb|removed|
|user_data|text|removed|
|virtual_machine_profile|jsonb|added|

## azure_compute_virtual_machine_secret_vault_certificates
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machine_secrets
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machine_win_config_rm_listeners
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machines

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|additional_capabilities|jsonb|added|
|additional_capabilities_ultra_ssd_enabled|boolean|removed|
|admin_password|text|removed|
|admin_username|text|removed|
|allow_extension_operations|boolean|removed|
|availability_set|jsonb|added|
|availability_set_id|text|removed|
|billing_profile|jsonb|added|
|billing_profile_max_price|float|removed|
|computer_name|text|removed|
|custom_data|text|removed|
|diagnostics_profile|jsonb|added|
|diagnostics_profile_boot_diagnostics_enabled|boolean|removed|
|diagnostics_profile_boot_diagnostics_storage_uri|text|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|hardware_profile|jsonb|added|
|hardware_profile_vm_size|text|removed|
|host|jsonb|added|
|host_group|jsonb|added|
|host_group_id|text|removed|
|host_id|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|linux_configuration_disable_password_authentication|boolean|removed|
|linux_configuration_patch_settings_assessment_mode|text|removed|
|linux_configuration_patch_settings_patch_mode|text|removed|
|linux_configuration_provision_vm_agent|boolean|removed|
|linux_configuration_ssh_public_keys|jsonb|removed|
|network_profile|jsonb|added|
|network_profile_network_api_version|text|removed|
|network_profile_network_interface_configurations|jsonb|removed|
|network_profile_network_interfaces|jsonb|removed|
|os_profile|jsonb|added|
|plan|jsonb|added|
|plan_name|text|removed|
|plan_product|text|removed|
|plan_promotion_code|text|removed|
|plan_publisher|text|removed|
|platform_fault_domain|bigint|updated|Type changed from integer to bigint
|proximity_placement_group|jsonb|added|
|proximity_placement_group_id|text|removed|
|require_guest_provision_signal|boolean|removed|
|resources|jsonb|added|
|security_profile|jsonb|added|
|security_profile_encryption_at_host|boolean|removed|
|security_profile_security_type|text|removed|
|security_profile_uefi_settings_secure_boot_enabled|boolean|removed|
|security_profile_uefi_settings_v_tpm_enabled|boolean|removed|
|virtual_machine_scale_set|jsonb|added|
|virtual_machine_scale_set_id|text|removed|
|windows_configuration_additional_unattend_content|jsonb|removed|
|windows_configuration_enable_automatic_updates|boolean|removed|
|windows_configuration_patch_settings_assessment_mode|text|removed|
|windows_configuration_patch_settings_enable_hotpatching|boolean|removed|
|windows_configuration_patch_settings_patch_mode|text|removed|
|windows_configuration_provision_vm_agent|boolean|removed|
|windows_configuration_time_zone|text|removed|

## azure_container_managed_cluster_agent_pool_profiles
Moved to JSON column on [azure_container_managed_clusters](#azure_container_managed_clusters)


## azure_container_managed_cluster_pip_user_assigned_id_exceptions
Moved to JSON column on [azure_container_managed_clusters](#azure_container_managed_clusters)


## azure_container_managed_cluster_pip_user_assigned_identities
Moved to JSON column on [azure_container_managed_clusters](#azure_container_managed_clusters)


## azure_container_managed_cluster_private_link_resources
Moved to JSON column on [azure_container_managed_clusters](#azure_container_managed_clusters)


## azure_container_managed_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|aad_profile|jsonb|added|
|aad_profile_admin_group_object_ids|text[]|removed|
|aad_profile_client_app_id|text|removed|
|aad_profile_enable_azure_rbac|boolean|removed|
|aad_profile_managed|boolean|removed|
|aad_profile_server_app_id|text|removed|
|aad_profile_server_app_secret|text|removed|
|aad_profile_tenant_id|text|removed|
|agent_pool_profiles|jsonb|added|
|api_server_access_profile|jsonb|added|
|api_server_access_profile_authorized_ip_ranges|text[]|removed|
|api_server_access_profile_enable_private_cluster|boolean|removed|
|api_server_access_profile_private_dns_zone|text|removed|
|auto_scaler_profile|jsonb|added|
|auto_scaler_profile_expander|text|removed|
|auto_upgrade_profile|jsonb|added|
|auto_upgrade_profile_upgrade_channel|text|removed|
|enable_pod_security_policy|boolean|added|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|http_proxy_config|jsonb|added|
|http_proxy_config_http_proxy|text|removed|
|http_proxy_config_https_proxy|text|removed|
|http_proxy_config_no_proxy|text[]|removed|
|http_proxy_config_trusted_ca|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|linux_profile|jsonb|added|
|linux_profile_admin_username|text|removed|
|max_agent_pools|bigint|updated|Type changed from integer to bigint
|network_profile|jsonb|added|
|network_profile_dns_service_ip|text|removed|
|network_profile_docker_bridge_cidr|text|removed|
|network_profile_load_balancer_allocated_outbound_ports|integer|removed|
|network_profile_load_balancer_effective_outbound_ips|text[]|removed|
|network_profile_load_balancer_idle_timeout|integer|removed|
|network_profile_load_balancer_managed_outbound_ips_count|integer|removed|
|network_profile_load_balancer_outbound_ip_prefixes|text[]|removed|
|network_profile_load_balancer_outbound_ips|text[]|removed|
|network_profile_load_balancer_sku|text|removed|
|network_profile_network_mode|text|removed|
|network_profile_network_plugin|text|removed|
|network_profile_network_policy|text|removed|
|network_profile_outbound_type|text|removed|
|network_profile_pod_cidr|text|removed|
|network_profile_service_cidr|text|removed|
|pod_identity_profile|jsonb|added|
|pod_identity_profile_allow_network_plugin_kubenet|boolean|removed|
|pod_identity_profile_enabled|boolean|removed|
|power_state|jsonb|added|
|power_state_code|text|removed|
|private_link_resources|jsonb|added|
|service_principal_profile|jsonb|added|
|service_principal_profile_client_id|text|removed|
|service_principal_profile_secret|text|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|
|windows_profile|jsonb|added|
|windows_profile_admin_password|text|removed|
|windows_profile_admin_username|text|removed|
|windows_profile_enable_csi_proxy|boolean|removed|
|windows_profile_license_type|text|removed|

## azure_container_registries

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|network_rule_set|jsonb|added|
|network_rule_set_default_action|text|removed|
|policies|jsonb|added|
|quarantine_policy_status|text|removed|
|retention_policy_days|integer|removed|
|retention_policy_last_updated_time|timestamp without time zone|removed|
|retention_policy_status|text|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|
|status|jsonb|updated|Type changed from text to jsonb
|status_message|text|removed|
|status_timestamp|timestamp without time zone|removed|
|storage_account|jsonb|added|
|storage_account_id|text|removed|
|trust_policy_status|text|removed|
|trust_policy_type|text|removed|

## azure_container_registry_network_rule_set_ip_rules
Moved to JSON column on [azure_container_registries](#azure_container_registries)


## azure_container_registry_network_rule_set_virtual_network_rules
Moved to JSON column on [azure_container_registries](#azure_container_registries)


## azure_container_registry_replications
Moved to JSON column on [azure_container_registries](#azure_container_registries)


## azure_container_replications
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|container_registry_id|text|added|
|provisioning_state|text|added|
|status|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|

## azure_cosmosdb_account_cors
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_account_failover_policies
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_account_locations
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_account_private_endpoint_connections
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_account_read_locations
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_account_write_locations
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|api_properties|jsonb|added|
|api_properties_server_version|text|removed|
|capabilities|jsonb|updated|Type changed from text[] to jsonb
|consistency_policy|jsonb|added|
|consistency_policy_default_consistency_level|text|removed|
|consistency_policy_max_interval_in_seconds|integer|removed|
|consistency_policy_max_staleness_prefix|bigint|removed|
|cors|jsonb|added|
|failover_policies|jsonb|added|
|ip_rules|jsonb|updated|Type changed from text[] to jsonb
|kind|text|added|
|locations|jsonb|added|
|private_endpoint_connections|jsonb|added|
|read_locations|jsonb|added|
|write_locations|jsonb|added|

## azure_cosmosdb_mongo_db_databases
Renamed from [azure_cosmosdb_mongodb_databases](azure_cosmosdb_mongodb_databases)


## azure_cosmosdb_mongodb_databases
Renamed to [azure_cosmosdb_mongo_db_databases](#azure_cosmosdb_mongo_db_databases)


## azure_cosmosdb_sql_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|autoscale_settings_max_throughput|integer|removed|
|cosmosdb_account_id|text|added|
|database_colls|text|removed|
|database_etag|text|removed|
|database_id|text|removed|
|database_rid|text|removed|
|database_ts|float|removed|
|database_users|text|removed|
|options|jsonb|added|
|resource|jsonb|added|
|sql_database_get_properties_throughput|integer|removed|

## azure_datalake_analytics_account_compute_policies
Moved to JSON column on [azure_datalake_analytics_accounts](#azure_datalake_analytics_accounts)


## azure_datalake_analytics_account_data_lake_store_accounts
Moved to JSON column on [azure_datalake_analytics_accounts](#azure_datalake_analytics_accounts)


## azure_datalake_analytics_account_firewall_rules
Moved to JSON column on [azure_datalake_analytics_accounts](#azure_datalake_analytics_accounts)


## azure_datalake_analytics_account_storage_accounts
Moved to JSON column on [azure_datalake_analytics_accounts](#azure_datalake_analytics_accounts)


## azure_datalake_analytics_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|compute_policies|jsonb|added|
|data_lake_store_accounts|jsonb|added|
|firewall_rules|jsonb|added|
|max_degree_of_parallelism|bigint|updated|Type changed from integer to bigint
|max_degree_of_parallelism_per_job|bigint|updated|Type changed from integer to bigint
|max_job_count|bigint|updated|Type changed from integer to bigint
|min_priority_per_job|bigint|updated|Type changed from integer to bigint
|query_store_retention|bigint|updated|Type changed from integer to bigint
|storage_accounts|jsonb|added|
|system_max_degree_of_parallelism|bigint|updated|Type changed from integer to bigint
|system_max_job_count|bigint|updated|Type changed from integer to bigint

## azure_datalake_storage_account_firewall_rules
This table was removed.


## azure_datalake_storage_account_trusted_id_providers
This table was removed.


## azure_datalake_storage_account_virtual_network_rules
This table was removed.


## azure_datalake_storage_accounts
This table was removed.


## azure_datalake_store_accounts
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|identity|jsonb|added|
|default_group|text|added|
|encryption_config|jsonb|added|
|encryption_state|text|added|
|encryption_provisioning_state|text|added|
|firewall_rules|jsonb|added|
|virtual_network_rules|jsonb|added|
|firewall_state|text|added|
|firewall_allow_azure_ips|text|added|
|trusted_id_providers|jsonb|added|
|trusted_id_provider_state|text|added|
|new_tier|text|added|
|current_tier|text|added|
|account_id|uuid|added|
|provisioning_state|text|added|
|state|text|added|
|creation_time|timestamp without time zone|added|
|last_modified_time|timestamp without time zone|added|
|endpoint|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|

## azure_eventhub_namespace_encryption_key_vault_properties
Moved to JSON column on [azure_eventhub_namespaces](#azure_eventhub_namespaces)


## azure_eventhub_namespaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|encryption|jsonb|added|
|encryption_key_source|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|maximum_throughput_units|bigint|updated|Type changed from integer to bigint
|network_rule_set|jsonb|removed|
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_name|text|removed|
|sku_tier|text|removed|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## azure_eventhub_network_rule_sets
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|eventhub_namespace_id|text|added|
|trusted_service_access_enabled|boolean|added|
|default_action|text|added|
|virtual_network_rules|jsonb|added|
|ip_rules|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_front_door_backend_pool_backends
This table was removed.


## azure_front_door_backend_pools
This table was removed.


## azure_front_door_frontend_endpoints
This table was removed.


## azure_front_door_health_probe_settings
This table was removed.


## azure_front_door_load_balancing_settings
This table was removed.


## azure_front_door_routing_rules
This table was removed.


## azure_front_door_rules_engine_rule_match_conditions
This table was removed.


## azure_front_door_rules_engine_rules
This table was removed.


## azure_front_door_rules_engines
This table was removed.


## azure_front_doors
This table was removed.


## azure_frontdoor_doors
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|resource_state|text|added|
|provisioning_state|text|added|
|cname|text|added|
|frontdoor_id|text|added|
|rules_engines|jsonb|added|
|friendly_name|text|added|
|routing_rules|jsonb|added|
|load_balancing_settings|jsonb|added|
|health_probe_settings|jsonb|added|
|backend_pools|jsonb|added|
|frontend_endpoints|jsonb|added|
|backend_pools_settings|jsonb|added|
|enabled_state|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|

## azure_iothub_devices
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|etag|text|added|
|properties_authorization_policies|jsonb|added|
|properties_disable_local_auth|boolean|added|
|properties_disable_device_sas|boolean|added|
|properties_disable_module_sas|boolean|added|
|properties_restrict_outbound_network_access|boolean|added|
|properties_allowed_fqdn_list|text[]|added|
|properties_public_network_access|text|added|
|properties_ip_filter_rules|jsonb|added|
|properties_network_rule_sets|jsonb|added|
|properties_min_tls_version|text|added|
|properties_private_endpoint_connections|jsonb|added|
|properties_provisioning_state|text|added|
|properties_state|text|added|
|properties_host_name|text|added|
|properties_event_hub_endpoints|jsonb|added|
|properties_routing|jsonb|added|
|properties_storage_endpoints|jsonb|added|
|properties_messaging_endpoints|jsonb|added|
|properties_enable_file_upload_notifications|boolean|added|
|properties_cloud_to_device|jsonb|added|
|properties_comments|text|added|
|properties_features|text|added|
|properties_locations|jsonb|added|
|properties_enable_data_residency|boolean|added|
|sku|jsonb|added|
|identity|jsonb|added|
|system_data|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|

## azure_iothub_hub_authorization_policies
This table was removed.


## azure_iothub_hub_ip_filter_rules
This table was removed.


## azure_iothub_hub_network_rule_sets_ip_rules
This table was removed.


## azure_iothub_hub_private_endpoint_connections
This table was removed.


## azure_iothub_hub_routing_endpoints_event_hubs
This table was removed.


## azure_iothub_hub_routing_endpoints_service_bus_queues
This table was removed.


## azure_iothub_hub_routing_endpoints_service_bus_topics
This table was removed.


## azure_iothub_hub_routing_endpoints_storage_containers
This table was removed.


## azure_iothub_hub_routing_routes
This table was removed.


## azure_iothub_hubs
This table was removed.


## azure_keyvault_keys
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|keyvault_vault_id|text|added|
|kid|text|added|
|attributes|jsonb|added|
|tags|jsonb|added|
|managed|boolean|added|

## azure_keyvault_managed_hsm
Moved to JSON column on [azure_keyvault_managed_hsms](#azure_keyvault_managed_hsms)


## azure_keyvault_managed_hsms
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|properties_tenant_id|uuid|added|
|properties_initial_admin_object_ids|text[]|added|
|properties_hsm_uri|text|added|
|properties_enable_soft_delete|boolean|added|
|properties_soft_delete_retention_in_days|bigint|added|
|properties_enable_purge_protection|boolean|added|
|properties_create_mode|text|added|
|properties_status_message|text|added|
|properties_provisioning_state|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|sku|jsonb|added|
|tags|jsonb|added|

## azure_keyvault_secrets
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|keyvault_vault_id|text|added|
|id|text|added|
|attributes|jsonb|added|
|tags|jsonb|added|
|content_type|text|added|
|managed|boolean|added|

## azure_keyvault_vault_access_policies
Moved to JSON column on [azure_keyvault_vaults](#azure_keyvault_vaults)


## azure_keyvault_vault_keys
Moved to JSON column on [azure_keyvault_vaults](#azure_keyvault_vaults)


## azure_keyvault_vault_private_endpoint_connections
Moved to JSON column on [azure_keyvault_vaults](#azure_keyvault_vaults)


## azure_keyvault_vault_secrets
Moved to JSON column on [azure_keyvault_vaults](#azure_keyvault_vaults)


## azure_keyvault_vaults

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|create_mode|text|removed|
|enable_purge_protection|boolean|removed|
|enable_rbac_authorization|boolean|removed|
|enable_soft_delete|boolean|removed|
|enabled_for_deployment|boolean|removed|
|enabled_for_disk_encryption|boolean|removed|
|enabled_for_template_deployment|boolean|removed|
|network_acls_bypass|text|removed|
|network_acls_default_action|text|removed|
|network_acls_ip_rules|text[]|removed|
|network_acls_virtual_network_rules|text[]|removed|
|properties_access_policies|jsonb|added|
|properties_create_mode|text|added|
|properties_enable_purge_protection|boolean|added|
|properties_enable_rbac_authorization|boolean|added|
|properties_enable_soft_delete|boolean|added|
|properties_enabled_for_deployment|boolean|added|
|properties_enabled_for_disk_encryption|boolean|added|
|properties_enabled_for_template_deployment|boolean|added|
|properties_network_acls|jsonb|added|
|properties_private_endpoint_connections|jsonb|added|
|properties_sku|jsonb|added|
|properties_soft_delete_retention_in_days|bigint|added|
|properties_tenant_id|uuid|added|
|properties_vault_uri|text|added|
|sku_family|text|removed|
|sku_name|text|removed|
|soft_delete_retention_in_days|integer|removed|
|tenant_id|uuid|removed|
|vault_uri|text|removed|

## azure_logic_app_workflows
This table was removed.


## azure_logic_diagnostic_settings
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|logic_workflow_id|text|added|
|storage_account_id|text|added|
|service_bus_rule_id|text|added|
|event_hub_authorization_rule_id|text|added|
|event_hub_name|text|added|
|metrics|jsonb|added|
|logs|jsonb|added|
|workspace_id|text|added|
|log_analytics_destination_type|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_logic_workflows
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|provisioning_state|text|added|
|created_time|timestamp without time zone|added|
|changed_time|timestamp without time zone|added|
|state|text|added|
|version|text|added|
|access_endpoint|text|added|
|endpoints_configuration|jsonb|added|
|access_control|jsonb|added|
|sku|jsonb|added|
|integration_account|jsonb|added|
|integration_service_environment|jsonb|added|
|parameters|jsonb|added|
|identity|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|

## azure_mariadb_configurations
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|mariadb_server_id|text|added|
|value|text|added|
|description|text|added|
|default_value|text|added|
|data_type|text|added|
|allowed_values|text|added|
|source|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_mariadb_server_configurations
Moved to JSON column on [azure_mariadb_servers](#azure_mariadb_servers)


## azure_mariadb_server_private_endpoint_connections
Moved to JSON column on [azure_mariadb_servers](#azure_mariadb_servers)


## azure_mariadb_servers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|backup_retention_days|integer|removed|
|earliest_restore_date|timestamp without time zone|added|
|earliest_restore_date_time|timestamp without time zone|removed|
|geo_redundant_backup|text|removed|
|private_endpoint_connections|jsonb|added|
|replica_capacity|bigint|updated|Type changed from integer to bigint
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_size|text|removed|
|sku_tier|text|removed|
|storage_autogrow|text|removed|
|storage_mb|integer|removed|
|storage_profile|jsonb|added|

## azure_monitor_activity_log_alert_action_groups
Moved to JSON column on [azure_monitor_activity_logs](#azure_monitor_activity_logs)


## azure_monitor_activity_log_alert_conditions
Moved to JSON column on [azure_monitor_activity_logs](#azure_monitor_activity_logs)


## azure_monitor_activity_log_alerts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|actions|jsonb|added|
|condition|jsonb|added|

## azure_monitor_activity_logs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|authorization|jsonb|added|
|authorization_action|text|removed|
|authorization_role|text|removed|
|authorization_scope|text|removed|
|category|jsonb|added|
|category_localized_value|text|removed|
|category_value|text|removed|
|event_name|jsonb|added|
|event_name_localized_value|text|removed|
|event_name_value|text|removed|
|event_timestamp|timestamp without time zone|added|
|event_timestamp_time|timestamp without time zone|removed|
|http_request|jsonb|added|
|http_request_client_ip_address|text|removed|
|http_request_client_request_id|text|removed|
|http_request_method|text|removed|
|http_request_uri|text|removed|
|operation_name|jsonb|added|
|operation_name_localized_value|text|removed|
|operation_name_value|text|removed|
|resource_provider_name|jsonb|added|
|resource_provider_name_localized_value|text|removed|
|resource_provider_name_value|text|removed|
|resource_type|jsonb|added|
|resource_type_localized_value|text|removed|
|resource_type_value|text|removed|
|status|jsonb|added|
|status_localized_value|text|removed|
|status_value|text|removed|
|sub_status|jsonb|added|
|sub_status_localized_value|text|removed|
|sub_status_value|text|removed|
|submission_timestamp|timestamp without time zone|added|
|submission_timestamp_time|timestamp without time zone|removed|

## azure_monitor_diagnostic_setting_logs
Moved to JSON column on [azure_monitor_diagnostic_settings](#azure_monitor_diagnostic_settings)


## azure_monitor_diagnostic_setting_metrics
Moved to JSON column on [azure_monitor_diagnostic_settings](#azure_monitor_diagnostic_settings)


## azure_monitor_diagnostic_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|logs|jsonb|added|
|metrics|jsonb|added|
|monitor_resource_id|text|added|

## azure_monitor_log_profiles

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|etag|text|added|
|kind|text|added|
|retention_policy|jsonb|added|
|retention_policy_days|integer|removed|
|retention_policy_enabled|boolean|removed|

## azure_monitor_resources
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|added|

## azure_mysql_configurations
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|mysql_server_id|text|added|
|value|text|added|
|description|text|added|
|default_value|text|added|
|data_type|text|added|
|allowed_values|text|added|
|source|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_mysql_server_configurations
Moved to JSON column on [azure_mysql_servers](#azure_mysql_servers)


## azure_mysql_server_private_endpoint_connections
Moved to JSON column on [azure_mysql_servers](#azure_mysql_servers)


## azure_mysql_servers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|earliest_restore_date|timestamp without time zone|added|
|earliest_restore_date_time|timestamp without time zone|removed|
|identity|jsonb|added|
|identity_principal_id|uuid|removed|
|identity_tenant_id|uuid|removed|
|identity_type|text|removed|
|private_endpoint_connections|jsonb|added|
|replica_capacity|bigint|updated|Type changed from integer to bigint
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_size|text|removed|
|sku_tier|text|removed|
|storage_profile|jsonb|added|
|storage_profile_backup_retention_days|integer|removed|
|storage_profile_geo_redundant_backup|text|removed|
|storage_profile_storage_autogrow|text|removed|
|storage_profile_storage_mb|integer|removed|

## azure_network_express_route_circuit_authorizations
Moved to JSON column on [azure_network_express_route_circuits](#azure_network_express_route_circuits)


## azure_network_express_route_circuit_connections
Moved to JSON column on [azure_network_express_route_circuits](#azure_network_express_route_circuits)


## azure_network_express_route_circuit_peerings
Moved to JSON column on [azure_network_express_route_circuits](#azure_network_express_route_circuits)


## azure_network_express_route_circuits

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|authorizations|jsonb|added|
|bandwidth_in_gbps|real|updated|Type changed from float to real
|express_route_port|jsonb|added|
|express_route_port_id|text|removed|
|peerings|jsonb|added|
|service_provider_properties|jsonb|added|
|service_provider_properties_bandwidth_in_mbps|integer|removed|
|service_provider_properties_peering_location|text|removed|
|service_provider_properties_service_provider_name|text|removed|
|sku|jsonb|added|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_tier|text|removed|
|stag|bigint|updated|Type changed from integer to bigint

## azure_network_express_route_connections
This table was removed.


## azure_network_express_route_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|auto_scale_configuration|jsonb|added|
|auto_scale_configuration_bound_max|integer|removed|
|auto_scale_configuration_bound_min|integer|removed|
|express_route_connections|jsonb|added|
|virtual_hub|jsonb|added|
|virtual_hub_id|text|removed|

## azure_network_express_route_links
This table was removed.


## azure_network_express_route_ports

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|bandwidth_in_gbps|bigint|updated|Type changed from integer to bigint
|circuits|jsonb|updated|Type changed from text[] to jsonb
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|links|jsonb|added|
|provisioned_bandwidth_in_gbps|real|updated|Type changed from float to real

## azure_network_flow_logs
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|network_watcher_id|text|added|
|target_resource_id|text|added|
|target_resource_guid|text|added|
|storage_id|text|added|
|enabled|boolean|added|
|retention_policy|jsonb|added|
|format|jsonb|added|
|flow_analytics_configuration|jsonb|added|
|provisioning_state|text|added|
|etag|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|

## azure_network_interface_ip_configurations
Moved to JSON column on [azure_network_interfaces](#azure_network_interfaces)


## azure_network_interfaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|dns_settings|jsonb|added|
|dns_settings_applied_dns_servers|text[]|removed|
|dns_settings_dns_servers|text[]|removed|
|dns_settings_internal_dns_name_label|text|removed|
|dns_settings_internal_domain_name_suffix|text|removed|
|dns_settings_internal_fqdn|text|removed|
|dscp_configuration|jsonb|added|
|dscp_configuration_id|text|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|ip_configurations|jsonb|added|
|network_security_group|jsonb|updated|Type changed from text to jsonb
|private_endpoint|jsonb|updated|Type changed from text to jsonb
|virtual_machine|jsonb|added|
|virtual_machine_id|text|removed|

## azure_network_peer_express_route_circuit_connections
This table was removed.


## azure_network_public_ip_addresses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|ddos_settings|jsonb|added|
|ddos_settings_ddos_custom_policy_id|text|removed|
|ddos_settings_protected_ip|boolean|removed|
|ddos_settings_protection_coverage|text|removed|
|dns_settings|jsonb|added|
|dns_settings_domain_name_label|text|removed|
|dns_settings_fqdn|text|removed|
|dns_settings_reverse_fqdn|text|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|idle_timeout_in_minutes|bigint|updated|Type changed from integer to bigint
|ip_address|text|updated|Type changed from inet to text
|public_ip_prefix|jsonb|added|
|public_ip_prefix_id|text|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|

## azure_network_route_filter_rules
Moved to JSON column on [azure_network_route_filters](#azure_network_route_filters)


## azure_network_route_filters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|rules|jsonb|added|

## azure_network_route_table_routes
Moved to JSON column on [azure_network_route_tables](#azure_network_route_tables)


## azure_network_route_tables

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|route_table_subnets|text[]|removed|
|routes|jsonb|added|
|subnets|jsonb|added|

## azure_network_security_group_default_security_rules
Moved to JSON column on [azure_network_security_groups](#azure_network_security_groups)


## azure_network_security_group_flow_logs
Moved to JSON column on [azure_network_security_groups](#azure_network_security_groups)


## azure_network_security_group_security_rules
Moved to JSON column on [azure_network_security_groups](#azure_network_security_groups)


## azure_network_security_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|default_security_rules|jsonb|added|
|flow_logs|jsonb|added|
|network_interfaces|jsonb|added|
|security_rules|jsonb|added|
|subnets|jsonb|added|

## azure_network_virtual_network_gateway_connections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|local_network_gateway2|jsonb|added|
|local_network_gateway_2|jsonb|removed|
|network_virtual_network_gateway_id|text|added|
|peer|jsonb|added|
|peer_id|text|removed|
|routing_weight|bigint|updated|Type changed from integer to bigint
|subscription_id|text|added|
|virtual_network_gateway1|jsonb|added|
|virtual_network_gateway2|jsonb|added|
|virtual_network_gateway_1|jsonb|removed|
|virtual_network_gateway_2|jsonb|removed|
|virtual_network_gateway_cq_id|uuid|removed|

## azure_network_virtual_network_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|bgp_settings|jsonb|added|
|bgp_settings_asn|bigint|removed|
|bgp_settings_bgp_peering_address|text|removed|
|bgp_settings_bgp_peering_addresses|jsonb|removed|
|bgp_settings_peer_weight|integer|removed|
|custom_routes|jsonb|added|
|custom_routes_address_prefixes|text[]|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|gateway_default_site|jsonb|added|
|gateway_default_site_id|text|removed|
|network_virtual_network_id|text|added|
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_name|text|removed|
|sku_tier|text|removed|
|subscription_id|text|added|
|v_net_extended_location_resource_id|text|added|
|virtual_network_cq_id|uuid|removed|
|vnet_extended_location_resource_id|text|removed|
|vpn_client_configuration|jsonb|added|
|vpn_client_configuration_aad_audience|text|removed|
|vpn_client_configuration_aad_issuer|text|removed|
|vpn_client_configuration_aad_tenant|text|removed|
|vpn_client_configuration_address_pool|text[]|removed|
|vpn_client_configuration_authentication_types|text[]|removed|
|vpn_client_configuration_ipsec_policies|jsonb|removed|
|vpn_client_configuration_protocols|text[]|removed|
|vpn_client_configuration_radius_server_address|text|removed|
|vpn_client_configuration_radius_server_secret|text|removed|
|vpn_client_configuration_radius_servers|jsonb|removed|
|vpn_client_configuration_revoked_certificates|jsonb|removed|
|vpn_client_configuration_root_certificates|jsonb|removed|

## azure_network_virtual_network_peerings
Moved to JSON column on [azure_network_virtual_networks](#azure_network_virtual_networks)


## azure_network_virtual_network_subnets
Moved to JSON column on [azure_network_virtual_networks](#azure_network_virtual_networks)


## azure_network_virtual_networks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|address_space|jsonb|added|
|address_space_address_prefixes|text[]|removed|
|bgp_communities|jsonb|added|
|bgp_communities_regional_community|text|removed|
|bgp_communities_virtual_network_community|text|removed|
|ddos_protection_plan|jsonb|added|
|ddos_protection_plan_id|text|removed|
|dhcp_options|jsonb|added|
|dhcp_options_dns_servers|inet[]|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|ip_allocations|jsonb|updated|Type changed from text[] to jsonb
|subnets|jsonb|added|
|virtual_network_peerings|jsonb|added|

## azure_network_watchers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## azure_postgresql_configurations
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|postgresql_server_id|text|added|
|value|text|added|
|description|text|added|
|default_value|text|added|
|data_type|text|added|
|allowed_values|text|added|
|source|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_postgresql_firewall_rules
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|postgresql_server_id|text|added|
|start_ip_address|text|added|
|end_ip_address|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_postgresql_server_configurations
Moved to JSON column on [azure_postgresql_servers](#azure_postgresql_servers)


## azure_postgresql_server_firewall_rules
Moved to JSON column on [azure_postgresql_servers](#azure_postgresql_servers)


## azure_postgresql_server_private_endpoint_connections
Moved to JSON column on [azure_postgresql_servers](#azure_postgresql_servers)


## azure_postgresql_servers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|earliest_restore_date|timestamp without time zone|added|
|earliest_restore_date_time|timestamp without time zone|removed|
|identity|jsonb|added|
|identity_principal_id|uuid|removed|
|identity_tenant_id|uuid|removed|
|identity_type|text|removed|
|private_endpoint_connections|jsonb|added|
|replica_capacity|bigint|updated|Type changed from integer to bigint
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_size|text|removed|
|sku_tier|text|removed|
|storage_profile|jsonb|added|
|storage_profile_backup_retention_days|integer|removed|
|storage_profile_geo_redundant_backup|text|removed|
|storage_profile_storage_autogrow|text|removed|
|storage_profile_storage_mb|integer|removed|

## azure_redis_caches
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|provisioning_state|text|added|
|host_name|text|added|
|port|bigint|added|
|ssl_port|bigint|added|
|access_keys|jsonb|added|
|linked_servers|jsonb|added|
|instances|jsonb|added|
|private_endpoint_connections|jsonb|added|
|sku|jsonb|added|
|subnet_id|text|added|
|static_ip|text|added|
|redis_configuration|jsonb|added|
|redis_version|text|added|
|enable_non_ssl_port|boolean|added|
|replicas_per_master|bigint|added|
|replicas_per_primary|bigint|added|
|tenant_settings|jsonb|added|
|shard_count|bigint|added|
|minimum_tls_version|text|added|
|public_network_access|text|added|
|zones|text[]|added|
|tags|jsonb|added|
|location|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_redis_services
This table was removed.


## azure_resources_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## azure_resources_links

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|notes|text|removed|
|properties_notes|text|added|
|properties_source_id|text|added|
|properties_target_id|text|added|
|source_id|text|removed|
|target_id|text|removed|
|type|text|removed|

## azure_resources_policy_assignments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|metadata|jsonb|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|

## azure_search_service_private_endpoint_connections
Moved to JSON column on [azure_search_services](#azure_search_services)


## azure_search_service_shared_private_link_resources
Moved to JSON column on [azure_search_services](#azure_search_services)


## azure_search_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|network_rule_set|jsonb|added|
|network_rule_set_ip_rules|inet[]|removed|
|partition_count|bigint|updated|Type changed from integer to bigint
|private_endpoint_connections|jsonb|added|
|replica_count|bigint|updated|Type changed from integer to bigint
|shared_private_link_resources|jsonb|added|
|sku|jsonb|added|
|sku_name|text|removed|

## azure_security_assessments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|azure_portal_uri|text|removed|
|cause|text|removed|
|code|text|removed|
|description|text|removed|
|links|jsonb|added|
|metadata|jsonb|added|
|metadata_assessment_type|text|removed|
|metadata_categories|text[]|removed|
|metadata_description|text|removed|
|metadata_display_name|text|removed|
|metadata_implementation_effort|text|removed|
|metadata_partner_data_partner_name|text|removed|
|metadata_partner_data_product_name|text|removed|
|metadata_policy_definition_id|text|removed|
|metadata_preview|boolean|removed|
|metadata_remediation_description|text|removed|
|metadata_severity|text|removed|
|metadata_threats|text[]|removed|
|metadata_user_impact|text|removed|
|partner_name|text|removed|
|partners_data|jsonb|added|
|resource_details|jsonb|removed|
|status|jsonb|added|

## azure_security_auto_provisioning_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|resource_type|text|removed|
|type|text|added|

## azure_security_contacts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|resource_type|text|removed|
|type|text|added|

## azure_security_jit_network_access_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|requests|jsonb|added|
|virtual_machines|jsonb|added|

## azure_security_jit_network_access_policy_requests
Moved to JSON column on [azure_security_jit_network_access_policies](#azure_security_jit_network_access_policies)


## azure_security_jit_network_access_policy_virtual_machines
Moved to JSON column on [azure_security_jit_network_access_policies](#azure_security_jit_network_access_policies)


## azure_security_pricings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|free_trial_remaining_time|text|added|
|pricing_properties_free_trial_remaining_time|text|removed|
|pricing_properties_tier|text|removed|
|pricing_tier|text|added|
|resource_type|text|removed|
|type|text|added|

## azure_security_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|resource_type|text|removed|
|type|text|added|

## azure_servicebus_access_keys
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|servicebus_authorization_rule_id|text|added|
|primary_connection_string|text|added|
|secondary_connection_string|text|added|
|alias_primary_connection_string|text|added|
|alias_secondary_connection_string|text|added|
|primary_key|text|added|
|secondary_key|text|added|
|key_name|text|added|

## azure_servicebus_authorization_rules
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|servicebus_topic_id|text|added|
|rights|text[]|added|
|system_data|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_servicebus_namespace_private_endpoint_connections
Moved to JSON column on [azure_servicebus_namespaces](#azure_servicebus_namespaces)


## azure_servicebus_namespace_topic_authorization_rules
Moved to JSON column on [azure_servicebus_namespaces](#azure_servicebus_namespaces)


## azure_servicebus_namespace_topics
Moved to JSON column on [azure_servicebus_namespaces](#azure_servicebus_namespaces)


## azure_servicebus_namespaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|encryption|jsonb|added|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|key_source|text|removed|
|key_vault_properties|jsonb|removed|
|private_endpoint_connections|jsonb|added|
|require_infrastructure_encryption|boolean|removed|
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_name|text|removed|
|sku_tier|text|removed|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|
|user_assigned_identities|jsonb|removed|

## azure_servicebus_topics
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|servicebus_namespace_id|text|added|
|size_in_bytes|bigint|added|
|created_at|timestamp without time zone|added|
|updated_at|timestamp without time zone|added|
|accessed_at|timestamp without time zone|added|
|subscription_count|bigint|added|
|count_details|jsonb|added|
|default_message_time_to_live|text|added|
|max_size_in_megabytes|bigint|added|
|max_message_size_in_kilobytes|bigint|added|
|requires_duplicate_detection|boolean|added|
|duplicate_detection_history_time_window|text|added|
|enable_batched_operations|boolean|added|
|status|text|added|
|support_ordering|boolean|added|
|auto_delete_on_idle|text|added|
|enable_partitioning|boolean|added|
|enable_express|boolean|added|
|system_data|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_backup_long_term_retention_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_database_id|text|added|
|weekly_retention|text|added|
|monthly_retention|text|added|
|yearly_retention|text|added|
|week_of_year|bigint|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_database_blob_auditing_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_database_id|text|added|
|kind|text|added|
|state|text|added|
|storage_endpoint|text|added|
|storage_account_access_key|text|added|
|retention_days|bigint|added|
|audit_actions_and_groups|text[]|added|
|storage_account_subscription_id|uuid|added|
|is_storage_secondary_key_in_use|boolean|added|
|is_azure_monitor_target_enabled|boolean|added|
|queue_delay_ms|bigint|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_database_db_blob_auditing_policies
Moved to JSON column on [azure_sql_databases](#azure_sql_databases)


## azure_sql_database_db_threat_detection_policies
Moved to JSON column on [azure_sql_databases](#azure_sql_databases)


## azure_sql_database_db_vulnerability_assessment_scans
Moved to JSON column on [azure_sql_databases](#azure_sql_databases)


## azure_sql_database_db_vulnerability_assessments
Moved to JSON column on [azure_sql_databases](#azure_sql_databases)


## azure_sql_database_threat_detection_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_database_id|text|added|
|location|text|added|
|kind|text|added|
|state|text|added|
|disabled_alerts|text|added|
|email_addresses|text|added|
|email_account_admins|text|added|
|storage_endpoint|text|added|
|storage_account_access_key|text|added|
|retention_days|bigint|added|
|use_server_default|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_database_vulnerability_assessment_scans
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_database_id|text|added|
|scan_id|text|added|
|trigger_type|text|added|
|state|text|added|
|start_time|timestamp without time zone|added|
|end_time|timestamp without time zone|added|
|errors|jsonb|added|
|storage_container_path|text|added|
|number_of_failed_security_checks|bigint|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_database_vulnerability_assessments
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_database_id|text|added|
|storage_container_path|text|added|
|storage_container_sas_key|text|added|
|storage_account_access_key|text|added|
|recurring_scans|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|auto_pause_delay|bigint|updated|Type changed from integer to bigint
|backup_long_term_retention_policy|jsonb|removed|
|creation_date|timestamp without time zone|added|
|creation_date_time|timestamp without time zone|removed|
|current_sku|jsonb|added|
|current_sku_capacity|integer|removed|
|current_sku_family|text|removed|
|current_sku_name|text|removed|
|current_sku_size|text|removed|
|current_sku_tier|text|removed|
|earliest_restore_date|timestamp without time zone|added|
|earliest_restore_date_time|timestamp without time zone|removed|
|high_availability_replica_count|bigint|updated|Type changed from integer to bigint
|min_capacity|real|updated|Type changed from float to real
|paused_date|timestamp without time zone|added|
|paused_date_time|timestamp without time zone|removed|
|resumed_date|timestamp without time zone|added|
|resumed_date_time|timestamp without time zone|removed|
|server_cq_id|uuid|removed|
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_size|text|removed|
|sku_tier|text|removed|
|source_database_deletion_date|timestamp without time zone|added|
|source_database_deletion_date_time|timestamp without time zone|removed|
|sql_server_id|text|added|
|subscription_id|text|added|
|transparent_data_encryption|jsonb|removed|

## azure_sql_encryption_protectors
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_server_id|text|added|
|kind|text|added|
|location|text|added|
|subregion|text|added|
|server_key_name|text|added|
|server_key_type|text|added|
|uri|text|added|
|thumbprint|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_firewall_rules
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_server_id|text|added|
|kind|text|added|
|location|text|added|
|start_ip_address|text|added|
|end_ip_address|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_managed_database_vulnerability_assessment_scans

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|managed_database_cq_id|uuid|removed|
|number_of_failed_security_checks|bigint|updated|Type changed from integer to bigint
|sql_managed_database_id|text|added|
|subscription_id|text|added|

## azure_sql_managed_database_vulnerability_assessments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|managed_database_cq_id|uuid|removed|
|recurring_scans|jsonb|added|
|recurring_scans_email_subscription_admins|boolean|removed|
|recurring_scans_emails|text[]|removed|
|recurring_scans_is_enabled|boolean|removed|
|sql_managed_database_id|text|added|
|subscription_id|text|added|

## azure_sql_managed_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|creation_date|timestamp without time zone|added|
|creation_date_time|timestamp without time zone|removed|
|earliest_restore_point|timestamp without time zone|added|
|earliest_restore_point_time|timestamp without time zone|removed|
|sql_managed_instance_id|text|added|
|subscription_id|text|added|

## azure_sql_managed_instance_encryption_protectors

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|managed_instance_cq_id|uuid|removed|
|sql_managed_instance_id|text|added|
|subscription_id|text|added|

## azure_sql_managed_instance_private_endpoint_connections
Moved to JSON column on [azure_sql_managed_instances](#azure_sql_managed_instances)


## azure_sql_managed_instance_vulnerability_assessments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|managed_instance_cq_id|uuid|removed|
|recurring_scans|jsonb|added|
|recurring_scans_email_subscription_admins|boolean|removed|
|recurring_scans_emails|text[]|removed|
|recurring_scans_is_enabled|boolean|removed|
|sql_managed_instance_id|text|added|
|subscription_id|text|added|

## azure_sql_managed_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|administrator_login_password|text|added|
|identity|jsonb|added|
|identity_principal_id|uuid|removed|
|identity_tenant_id|uuid|removed|
|identity_type|text|removed|
|private_endpoint_connections|jsonb|added|
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_size|text|removed|
|sku_tier|text|removed|
|storage_size_in_gb|bigint|updated|Type changed from integer to bigint
|v_cores|bigint|updated|Type changed from integer to bigint

## azure_sql_server_admins

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|server_cq_id|uuid|removed|
|sql_server_id|text|added|
|subscription_id|text|added|

## azure_sql_server_blob_auditing_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_server_id|text|added|
|state|text|added|
|storage_endpoint|text|added|
|storage_account_access_key|text|added|
|retention_days|bigint|added|
|audit_actions_and_groups|text[]|added|
|storage_account_subscription_id|uuid|added|
|is_storage_secondary_key_in_use|boolean|added|
|is_azure_monitor_target_enabled|boolean|added|
|queue_delay_ms|bigint|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_server_db_blob_auditing_policies
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_dev_ops_auditing_settings
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_server_id|text|added|
|system_data|jsonb|added|
|is_azure_monitor_target_enabled|boolean|added|
|state|text|added|
|storage_endpoint|text|added|
|storage_account_access_key|text|added|
|storage_account_subscription_id|uuid|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_server_devops_audit_settings
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_encryption_protectors
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_firewall_rules
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_private_endpoint_connections
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_security_alert_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_server_id|text|added|
|state|text|added|
|disabled_alerts|text[]|added|
|email_addresses|text[]|added|
|email_account_admins|boolean|added|
|storage_endpoint|text|added|
|storage_account_access_key|text|added|
|retention_days|bigint|added|
|creation_time|timestamp without time zone|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_server_security_alert_policy
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_virtual_network_rules
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_vulnerability_assessments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|recurring_scans|jsonb|added|
|recurring_scans_email_subscription_admins|boolean|removed|
|recurring_scans_emails|text[]|removed|
|recurring_scans_is_enabled|boolean|removed|
|server_cq_id|uuid|removed|
|sql_server_id|text|added|
|subscription_id|text|added|

## azure_sql_servers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|identity|jsonb|added|
|identity_principal_id|uuid|removed|
|identity_tenant_id|uuid|removed|
|identity_type|text|removed|
|private_endpoint_connections|jsonb|added|

## azure_sql_transparent_data_encryptions
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_database_id|text|added|
|location|text|added|
|status|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_sql_virtual_network_rules
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sql_server_id|text|added|
|virtual_network_subnet_id|text|added|
|ignore_missing_vnet_service_endpoint|boolean|added|
|state|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_storage_account_network_rule_set_ip_rules
Moved to JSON column on [azure_storage_accounts](#azure_storage_accounts)


## azure_storage_account_network_rule_set_virtual_network_rules
Moved to JSON column on [azure_storage_accounts](#azure_storage_accounts)


## azure_storage_account_private_endpoint_connections
Moved to JSON column on [azure_storage_accounts](#azure_storage_accounts)


## azure_storage_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|allow_shared_key_access|boolean|added|
|azure_files_identity_based_authentication|jsonb|added|
|blob_restore_status|jsonb|updated|Type changed from text to jsonb
|blob_restore_status_failure_reason|text|removed|
|blob_restore_status_parameters_blob_ranges|jsonb|removed|
|blob_restore_status_parameters_time_to_restore_time|timestamp without time zone|removed|
|blob_restore_status_restore_id|text|removed|
|custom_domain|jsonb|added|
|custom_domain_name|text|removed|
|custom_domain_use_sub_domain_name|boolean|removed|
|enable_https_traffic_only|boolean|removed|
|encryption|jsonb|added|
|encryption_key_current_versioned_key_identifier|text|removed|
|encryption_key_last_key_rotation_timestamp_time|timestamp without time zone|removed|
|encryption_key_source|text|removed|
|encryption_key_vault_properties_key_name|text|removed|
|encryption_key_vault_properties_key_vault_uri|text|removed|
|encryption_key_vault_properties_key_version|text|removed|
|encryption_require_infrastructure_encryption|boolean|removed|
|encryption_services_blob_enabled|boolean|removed|
|encryption_services_blob_key_type|text|removed|
|encryption_services_blob_last_enabled_time|timestamp without time zone|removed|
|encryption_services_file_enabled|boolean|removed|
|encryption_services_file_key_type|text|removed|
|encryption_services_file_last_enabled_time|timestamp without time zone|removed|
|encryption_services_queue_enabled|boolean|removed|
|encryption_services_queue_key_type|text|removed|
|encryption_services_queue_last_enabled_time|timestamp without time zone|removed|
|encryption_services_table_enabled|boolean|removed|
|encryption_services_table_key_type|text|removed|
|encryption_services_table_last_enabled_time|timestamp without time zone|removed|
|extended_location|jsonb|added|
|files_identity_auth_ad_properties_azure_storage_sid|text|removed|
|files_identity_auth_ad_properties_domain_guid|text|removed|
|files_identity_auth_ad_properties_domain_name|text|removed|
|files_identity_auth_ad_properties_forest_name|text|removed|
|files_identity_auth_ad_properties_net_bios_domain_name|text|removed|
|files_identity_auth_ad_properties_net_bios_domain_sid|text|removed|
|files_identity_auth_directory_service_options|text|removed|
|geo_replication_stats|jsonb|added|
|geo_replication_stats_can_failover|boolean|removed|
|geo_replication_stats_last_sync_time|timestamp without time zone|removed|
|geo_replication_stats_status|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|is_nfs_v3_enabled|boolean|added|
|network_acls|jsonb|added|
|network_rule_set_bypass|text|removed|
|network_rule_set_default_action|text|removed|
|primary_endpoints|jsonb|added|
|primary_endpoints_blob|text|removed|
|primary_endpoints_dfs|text|removed|
|primary_endpoints_file|text|removed|
|primary_endpoints_internet_endpoints_blob|text|removed|
|primary_endpoints_internet_endpoints_dfs|text|removed|
|primary_endpoints_internet_endpoints_file|text|removed|
|primary_endpoints_internet_endpoints_web|text|removed|
|primary_endpoints_microsoft_endpoints_blob|text|removed|
|primary_endpoints_microsoft_endpoints_dfs|text|removed|
|primary_endpoints_microsoft_endpoints_file|text|removed|
|primary_endpoints_microsoft_endpoints_queue|text|removed|
|primary_endpoints_microsoft_endpoints_table|text|removed|
|primary_endpoints_microsoft_endpoints_web|text|removed|
|primary_endpoints_queue|text|removed|
|primary_endpoints_table|text|removed|
|primary_endpoints_web|text|removed|
|private_endpoint_connections|jsonb|added|
|routing_preference|jsonb|added|
|routing_preference_publish_internet_endpoints|boolean|removed|
|routing_preference_publish_microsoft_endpoints|boolean|removed|
|routing_preference_routing_choice|text|removed|
|secondary_endpoints|jsonb|added|
|secondary_endpoints_blob|text|removed|
|secondary_endpoints_dfs|text|removed|
|secondary_endpoints_file|text|removed|
|secondary_endpoints_internet_endpoints_blob|text|removed|
|secondary_endpoints_internet_endpoints_dfs|text|removed|
|secondary_endpoints_internet_endpoints_file|text|removed|
|secondary_endpoints_internet_endpoints_web|text|removed|
|secondary_endpoints_microsoft_endpoints_blob|text|removed|
|secondary_endpoints_microsoft_endpoints_dfs|text|removed|
|secondary_endpoints_microsoft_endpoints_file|text|removed|
|secondary_endpoints_microsoft_endpoints_queue|text|removed|
|secondary_endpoints_microsoft_endpoints_table|text|removed|
|secondary_endpoints_microsoft_endpoints_web|text|removed|
|secondary_endpoints_queue|text|removed|
|secondary_endpoints_table|text|removed|
|secondary_endpoints_web|text|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|
|supports_https_traffic_only|boolean|added|

## azure_storage_blob_service_cors_rules
Moved to JSON column on [azure_storage_blob_services](#azure_storage_blob_services)


## azure_storage_blob_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_cq_id|uuid|removed|
|change_feed|jsonb|added|
|change_feed_enabled|boolean|removed|
|change_feed_retention_in_days|integer|removed|
|container_delete_retention_policy|jsonb|added|
|container_delete_retention_policy_days|integer|removed|
|container_delete_retention_policy_enabled|boolean|removed|
|cors|jsonb|added|
|delete_retention_policy|jsonb|added|
|delete_retention_policy_days|integer|removed|
|delete_retention_policy_enabled|boolean|removed|
|last_access_time_tracking_policy|jsonb|added|
|last_access_time_tracking_policy_blob_type|text[]|removed|
|last_access_time_tracking_policy_enable|boolean|removed|
|last_access_time_tracking_policy_name|text|removed|
|last_access_time_tracking_policy_tracking_granularity_in_days|integer|removed|
|restore_policy|jsonb|added|
|restore_policy_days|integer|removed|
|restore_policy_enabled|boolean|removed|
|restore_policy_last_enabled_time|timestamp without time zone|removed|
|restore_policy_min_restore_time|timestamp without time zone|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|
|storage_account_id|text|added|
|subscription_id|text|added|

## azure_storage_containers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_cq_id|uuid|removed|
|account_id|text|removed|
|remaining_retention_days|bigint|updated|Type changed from integer to bigint
|storage_account_id|text|added|

## azure_streamanalytics_jobs
This table was removed.


## azure_streamanalytics_streaming_jobs
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|sku|jsonb|added|
|job_id|text|added|
|provisioning_state|text|added|
|job_state|text|added|
|job_type|text|added|
|output_start_mode|text|added|
|output_start_time|timestamp without time zone|added|
|last_output_event_time|timestamp without time zone|added|
|events_out_of_order_policy|text|added|
|output_error_policy|text|added|
|events_out_of_order_max_delay_in_seconds|bigint|added|
|events_late_arrival_max_delay_in_seconds|bigint|added|
|data_locale|text|added|
|compatibility_level|text|added|
|created_date|timestamp without time zone|added|
|inputs|jsonb|added|
|transformation|jsonb|added|
|outputs|jsonb|added|
|functions|jsonb|added|
|etag|text|added|
|job_storage_account|jsonb|added|
|content_storage_policy|text|added|
|cluster|jsonb|added|
|identity|jsonb|added|
|tags|jsonb|added|
|location|text|added|
|id|text|added|
|name|text|added|
|type|text|added|

## azure_subscription_subscriptions
Moved to JSON column on [azure_subscriptions](#azure_subscriptions)


## azure_subscription_tenants
Moved to JSON column on [azure_subscriptions](#azure_subscriptions)


## azure_subscriptions
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|authorization_source|text|added|
|managed_by_tenants|jsonb|added|
|subscription_policies|jsonb|added|
|tags|jsonb|added|
|display_name|text|added|
|id|text|added|
|state|text|added|
|tenant_id|text|added|

## azure_subscriptions_locations
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|metadata|jsonb|added|
|display_name|text|added|
|id|text|added|
|name|text|added|
|regional_display_name|text|added|
|type|text|added|

## azure_subscriptions_tenants
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|country|text|added|
|country_code|text|added|
|default_domain|text|added|
|display_name|text|added|
|domains|jsonb|added|
|id|text|added|
|tenant_branding_logo_url|text|added|
|tenant_category|text|added|
|tenant_id|text|added|
|tenant_type|text|added|

## azure_web_app_auth_settings
Moved to JSON column on [azure_web_apps](#azure_web_apps)


## azure_web_app_host_name_ssl_states
Moved to JSON column on [azure_web_apps](#azure_web_apps)


## azure_web_app_publishing_profiles
Moved to JSON column on [azure_web_apps](#azure_web_apps)


## azure_web_apps

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|cloning_info|jsonb|added|
|cloning_info_app_settings_overrides|jsonb|removed|
|cloning_info_clone_custom_host_names|boolean|removed|
|cloning_info_clone_source_control|boolean|removed|
|cloning_info_configure_load_balancing|boolean|removed|
|cloning_info_correlation_id|uuid|removed|
|cloning_info_hosting_environment|text|removed|
|cloning_info_overwrite|boolean|removed|
|cloning_info_source_web_app_id|text|removed|
|cloning_info_source_web_app_location|text|removed|
|cloning_info_traffic_manager_profile_id|text|removed|
|cloning_info_traffic_manager_profile_name|text|removed|
|container_size|bigint|updated|Type changed from integer to bigint
|daily_memory_time_quota|bigint|updated|Type changed from integer to bigint
|host_name_ssl_states|jsonb|added|
|hosting_environment_profile|jsonb|added|
|hosting_environment_profile_id|text|removed|
|hosting_environment_profile_name|text|removed|
|hosting_environment_profile_type|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|last_modified_time_utc|timestamp without time zone|added|
|last_modified_time_utc_time|timestamp without time zone|removed|
|max_number_of_workers|bigint|updated|Type changed from integer to bigint
|slot_swap_status|jsonb|added|
|slot_swap_status_destination_slot_name|text|removed|
|slot_swap_status_source_slot_name|text|removed|
|slot_swap_status_timestamp_utc_time|timestamp without time zone|removed|
|suspended_till|timestamp without time zone|added|
|suspended_till_time|timestamp without time zone|removed|
|virtual_network_subnet_id|text|added|
|vnet_connection|jsonb|removed|

## azure_web_publishing_profiles
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|web_app_id|text|added|
|publish_url|text|added|
|user_name|text|added|
|user_pwd|text|added|

## azure_web_site_auth_settings
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|web_app_id|text|added|
|enabled|boolean|added|
|runtime_version|text|added|
|unauthenticated_client_action|text|added|
|token_store_enabled|boolean|added|
|allowed_external_redirect_urls|text[]|added|
|default_provider|text|added|
|token_refresh_extension_hours|real|added|
|client_id|text|added|
|client_secret|text|added|
|client_secret_setting_name|text|added|
|client_secret_certificate_thumbprint|text|added|
|issuer|text|added|
|validate_issuer|boolean|added|
|allowed_audiences|text[]|added|
|additional_login_params|text[]|added|
|aad_claims_authorization|text|added|
|google_client_id|text|added|
|google_client_secret|text|added|
|google_client_secret_setting_name|text|added|
|google_o_auth_scopes|text[]|added|
|facebook_app_id|text|added|
|facebook_app_secret|text|added|
|facebook_app_secret_setting_name|text|added|
|facebook_o_auth_scopes|text[]|added|
|git_hub_client_id|text|added|
|git_hub_client_secret|text|added|
|git_hub_client_secret_setting_name|text|added|
|git_hub_o_auth_scopes|text[]|added|
|twitter_consumer_key|text|added|
|twitter_consumer_secret|text|added|
|twitter_consumer_secret_setting_name|text|added|
|microsoft_account_client_id|text|added|
|microsoft_account_client_secret|text|added|
|microsoft_account_client_secret_setting_name|text|added|
|microsoft_account_o_auth_scopes|text[]|added|
|is_auth_from_file|text|added|
|auth_file_path|text|added|
|config_version|text|added|
|id|text|added|
|name|text|added|
|kind|text|added|
|type|text|added|

## azure_web_vnet_connections
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subscription_id|text|added|
|web_app_id|text|added|
|vnet_resource_id|text|added|
|cert_thumbprint|text|added|
|cert_blob|text|added|
|routes|jsonb|added|
|resync_required|boolean|added|
|dns_servers|text|added|
|is_swift|boolean|added|
|id|text|added|
|name|text|added|
|kind|text|added|
|type|text|added|
