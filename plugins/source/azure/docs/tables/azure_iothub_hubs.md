
# Table: azure_iothub_hubs
Azure IoT hub.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|etag|text|The Etag.|
|disable_local_auth|boolean|If true, SAS tokens with Iot hub scoped SAS keys cannot be used for authentication.|
|disable_device_sas|boolean|If true, all device(including Edge devices but excluding modules) scoped SAS keys cannot be used for authentication.|
|disable_module_sas|boolean|If true, all module scoped SAS keys cannot be used for authentication.|
|restrict_outbound_network_access|boolean|If true, egress from IotHub will be restricted to only the allowed FQDNs that are configured via allowedFqdnList.|
|allowed_fqdn_list|text[]|List of allowed FQDNs(Fully Qualified Domain Name) for egress from Iot Hub.|
|public_network_access|text|Whether requests from Public Network are allowed|
|network_rule_sets_default_action|text|Default Action for Network Rule Set|
|network_rule_sets_apply_to_built_in_event_hub_endpoint|boolean|If True, then Network Rule Set is also applied to BuiltIn EventHub EndPoint of IotHub|
|min_tls_version|text|Specifies the minimum TLS version to support for this hub|
|provisioning_state|text|The provisioning state.|
|state|text|The hub state.|
|host_name|text|The name of the host.|
|event_hub_endpoints|jsonb|The Event Hub-compatible endpoint properties|
|routing_fallback_route_name|text|The name of the route|
|routing_fallback_route_source|text|The source to which the routing rule is to be applied to|
|routing_fallback_route_condition|text|The condition which is evaluated in order to apply the fallback route|
|routing_fallback_route_endpoint_names|text[]|The list of endpoints to which the messages that satisfy the condition are routed to|
|routing_fallback_route_is_enabled|boolean|Used to specify whether the fallback route is enabled.|
|routing_enrichments|jsonb|The list of user-provided enrichments that the IoT hub applies to messages to be delivered to built-in and custom endpoints|
|storage_endpoints|jsonb|The list of Azure Storage endpoints where you can upload files|
|messaging_endpoints|jsonb|The messaging endpoint properties for the file upload notification queue.|
|enable_file_upload_notifications|boolean|If True, file upload notifications are enabled.|
|cloud_to_device_max_delivery_count|integer|The max delivery count for cloud-to-device messages in the device queue|
|cloud_to_device_default_ttl_as_iso8601|text|The default time to live for cloud-to-device messages in the device queue|
|cloud_to_device_feedback_lock_duration_as_iso8601|text|The lock duration for the feedback queue|
|cloud_to_device_feedback_ttl_as_iso8601|text|The period of time for which a message is available to consume before it is expired by the IoT hub|
|cloud_to_device_feedback_max_delivery_count|integer|The number of times the IoT hub attempts to deliver a message on the feedback queue|
|comments|text|IoT hub comments.|
|features|text|The capabilities and features enabled for the IoT hub|
|locations|jsonb|Primary and secondary location for iot hub|
|enable_data_residency|boolean|This property when set to true, will enable data residency, thus, disabling disaster recovery.|
|sku_name|text|The name of the SKU|
|sku_tier|text|The billing tier for the IoT hub|
|sku_capacity|bigint|The number of provisioned IoT Hub units|
|identity_principal_id|text|Principal Id|
|identity_tenant_id|text|Tenant Id|
|identity_type|text|The type of identity used for the resource|
|identity_user_assigned_identities|jsonb|The identities of assigned users|
|system_data_created_by|text|The identity that created the resource.|
|system_data_created_by_type|text|The type of identity that created the resource|
|system_data_created_at_time|timestamp without time zone|Created time|
|system_data_last_modified_by|text|The identity that last modified the resource.|
|system_data_last_modified_by_type|text|The type of identity that last modified the resource|
|system_data_last_modified_at_time|timestamp without time zone|Last modified time|
|id|text|The resource identifier.|
|name|text|The resource name.|
|type|text|The resource type.|
|location|text|The resource location.|
|tags|jsonb|The resource tags.|
