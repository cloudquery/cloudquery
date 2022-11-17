# Table: azure_iot_hubs

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub#Description

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
|sku|JSON|
|etag|String|
|identity|JSON|
|allowed_fqdn_list|StringArray|
|authorization_policies|JSON|
|cloud_to_device|JSON|
|comments|String|
|disable_device_sas|Bool|
|disable_local_auth|Bool|
|disable_module_sas|Bool|
|enable_data_residency|Bool|
|enable_file_upload_notifications|Bool|
|event_hub_endpoints|JSON|
|features|String|
|ip_filter_rules|JSON|
|messaging_endpoints|JSON|
|min_tls_version|String|
|network_rule_sets|JSON|
|private_endpoint_connections|JSON|
|public_network_access|String|
|restrict_outbound_network_access|Bool|
|routing|JSON|
|storage_endpoints|JSON|
|host_name|String|
|locations|JSON|
|provisioning_state|String|
|state|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|