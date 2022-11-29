# Table: azure_iothub_devices

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-07-02/devices#IotHubDescription

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|etag|String|
|properties_authorization_policies|JSON|
|properties_disable_local_auth|Bool|
|properties_disable_device_sas|Bool|
|properties_disable_module_sas|Bool|
|properties_restrict_outbound_network_access|Bool|
|properties_allowed_fqdn_list|StringArray|
|properties_public_network_access|String|
|properties_ip_filter_rules|JSON|
|properties_network_rule_sets|JSON|
|properties_min_tls_version|String|
|properties_private_endpoint_connections|JSON|
|properties_provisioning_state|String|
|properties_state|String|
|properties_host_name|String|
|properties_event_hub_endpoints|JSON|
|properties_routing|JSON|
|properties_storage_endpoints|JSON|
|properties_messaging_endpoints|JSON|
|properties_enable_file_upload_notifications|Bool|
|properties_cloud_to_device|JSON|
|properties_comments|String|
|properties_features|String|
|properties_locations|JSON|
|properties_enable_data_residency|Bool|
|sku|JSON|
|identity|JSON|
|system_data|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|