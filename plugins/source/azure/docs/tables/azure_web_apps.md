
# Table: azure_web_apps
Site a web app, a mobile app backend, or an API app
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|state|text|Current state of the app|
|host_names|text[]|Hostnames associated with the app|
|repository_site_name|text|Name of the repository site|
|usage_state|text|State indicating whether the app has exceeded its quota usage Read-only Possible values include: 'UsageStateNormal', 'UsageStateExceeded'|
|enabled|boolean|otherwise, <code>false</code> Setting this value to false disables the app (takes the app offline)|
|enabled_host_names|text[]|Enabled hostnames for the appHostnames need to be assigned (see HostNames) AND enabled Otherwise, the app is not served on those hostnames|
|availability_state|text|Management information availability state for the app Possible values include: 'Normal', 'Limited', 'DisasterRecoveryMode'|
|server_farm_id|text|Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/MicrosoftWeb/serverfarms/{appServicePlanName}"|
|reserved|boolean|otherwise, <code>false</code>|
|is_xenon|boolean|Obsolete: Hyper-V sandbox|
|hyper_v|boolean|Hyper-V sandbox|
|last_modified_time_utc_time|timestamp without time zone||
|site_config|jsonb|Configuration of the app|
|traffic_manager_host_names|text[]|Azure Traffic Manager hostnames associated with the app Read-only|
|scm_site_also_stopped|boolean|otherwise, <code>false</code> The default is <code>false</code>|
|target_swap_slot|text|Specifies which deployment slot this app will swap into Read-only|
|hosting_environment_profile_id|text|Resource ID of the App Service Environment|
|hosting_environment_profile_name|text|Name of the App Service Environment|
|hosting_environment_profile_type|text|Resource type of the App Service Environment|
|client_affinity_enabled|boolean|Set to true to enable client affinity.|
|client_cert_enabled|boolean|Set to true to enable client certificate authentication (TLS mutual authentication).|
|client_cert_mode|text|This composes with ClientCertEnabled setting - ClientCertEnabled: false means ClientCert is ignored - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted Possible values include: 'Required', 'Optional', 'OptionalInteractiveUser'|
|client_cert_exclusion_paths|text|client certificate authentication comma-separated exclusion paths|
|host_names_disabled|boolean|otherwise, <code>false</code>  If <code>true</code>, the app is only accessible via API management process|
|custom_domain_verification_id|text|Unique identifier that verifies the custom domains assigned to the app Customer will add this id to a txt record for verification|
|outbound_ip_addresses|text|List of IP addresses that the app uses for outbound connections (eg database access) Includes VIPs from tenants that site can be hosted with current settings Read-only|
|possible_outbound_ip_addresses|text|List of IP addresses that the app uses for outbound connections (eg database access) Includes VIPs from all tenants except dataComponent Read-only|
|container_size|integer|Size of the function container|
|daily_memory_time_quota|integer|Maximum allowed daily memory-time quota (applicable on dynamic apps only)|
|suspended_till_time|timestamp without time zone||
|max_number_of_workers|integer|Maximum number of workers This only applies to Functions container|
|cloning_info_correlation_id|uuid|Correlation ID of cloning operation This ID ties multiple cloning operations together to use the same snapshot|
|cloning_info_overwrite|boolean|otherwise, <code>false</code>|
|cloning_info_clone_custom_host_names|boolean|otherwise, <code>false</code>|
|cloning_info_clone_source_control|boolean|otherwise, <code>false</code>|
|cloning_info_source_web_app_id|text|ARM resource ID of the source app App resource ID is of the form /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftWeb/sites/{siteName} for production slots and /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftWeb/sites/{siteName}/slots/{slotName} for other slots|
|cloning_info_source_web_app_location|text|Location of source app ex: West US or North Europe|
|cloning_info_hosting_environment|text|App Service Environment|
|cloning_info_app_settings_overrides|jsonb|Application setting overrides for cloned app If specified, these settings override the settings cloned from source app Otherwise, application settings from source app are retained|
|cloning_info_configure_load_balancing|boolean|<code>true</code> to configure load balancing for source and destination app|
|cloning_info_traffic_manager_profile_id|text|ARM resource ID of the Traffic Manager profile to use, if it exists Traffic Manager resource ID is of the form /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftNetwork/trafficManagerProfiles/{profileName}|
|cloning_info_traffic_manager_profile_name|text|Name of Traffic Manager profile to create This is only needed if Traffic Manager profile does not already exist|
|resource_group|text|Name of the resource group the app belongs to Read-only|
|is_default_container|boolean|<code>true</code> if the app is a default container; otherwise, <code>false</code>|
|default_host_name|text|Default hostname of the app Read-only|
|slot_swap_status_timestamp_utc_time|timestamp without time zone||
|slot_swap_status_source_slot_name|text|The source slot of the last swap operation|
|slot_swap_status_destination_slot_name|text|The destination slot of the last swap operation|
|key_vault_reference_identity|text|Identity to use for Key Vault Reference authentication|
|https_only|boolean|HttpsOnly: configures a web site to accept only https requests Issues redirect for http requests|
|redundancy_mode|text|Site redundancy mode Possible values include: 'RedundancyModeNone', 'RedundancyModeManual', 'RedundancyModeFailover', 'RedundancyModeActiveActive', 'RedundancyModeGeoRedundant'|
|in_progress_operation_id|uuid|Specifies an operation id if this site has a pending operation|
|storage_account_required|boolean|Checks if Customer provided storage account is required|
|identity_type|text|Type of managed service identity Possible values include: 'ManagedServiceIdentityTypeSystemAssigned', 'ManagedServiceIdentityTypeUserAssigned', 'ManagedServiceIdentityTypeSystemAssignedUserAssigned', 'ManagedServiceIdentityTypeNone'|
|identity_tenant_id|text|Tenant of managed service identity|
|identity_principal_id|text|Principal Id of managed service identity|
|identity_user_assigned_identities|jsonb|The list of user assigned identities associated with the resource The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/MicrosoftManagedIdentity/userAssignedIdentities/{identityName}|
|vnet_connection|jsonb|Describes the virtual network connection for the web app.|
|id|text|Resource Id|
|name|text|Resource Name|
|kind|text|Kind of resource|
|location|text|Resource Location|
|type|text|Resource type|
|tags|jsonb|Resource tags|
