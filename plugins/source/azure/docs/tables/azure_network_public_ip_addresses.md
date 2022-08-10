
# Table: azure_network_public_ip_addresses
PublicIPAddress public IP address resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|extended_location_name|text|The name of the extended location.|
|extended_location_type|text|The type of the extended location.|
|sku_name|text|Name of a public IP address SKU|
|sku_tier|text|Tier of a public IP address SKU|
|public_ip_allocation_method|text|The public IP address allocation method|
|public_ip_address_version|text|The public IP address version|
|ip_configuration|jsonb|The IP configuration associated with the public IP address.|
|dns_settings_domain_name_label|text|The domain name label.|
|dns_settings_fqdn|text|The Fully Qualified Domain Name of the A DNS record associated with the public IP.|
|dns_settings_reverse_fqdn|text|The reverse FQDN.|
|ddos_settings_ddos_custom_policy_id|text|Resource ID.|
|ddos_settings_protection_coverage|text|The DDoS protection policy customizability of the public IP|
|ddos_settings_protected_ip|boolean|Enables DDoS protection on the public IP.|
|ip_tags|jsonb|The list of tags associated with the public IP address.|
|ip_address|inet|The IP address associated with the public IP address resource.|
|public_ip_prefix_id|text|Resource ID.|
|idle_timeout_in_minutes|integer|The idle timeout of the public IP address.|
|resource_guid|text|The resource GUID property of the public IP address resource.|
|provisioning_state|text|The provisioning state of the public IP address resource|
|service_public_ip_address|jsonb|The service public IP address of the public IP address resource.|
|nat_gateway|jsonb|The NatGateway for the Public IP address.|
|migration_phase|text|Migration phase of Public IP Address|
|linked_public_ip_address|jsonb|The linked public IP address of the public IP address resource.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|zones|text[]|A list of availability zones denoting the IP allocated for the resource needs to come from.|
|id|text|Resource ID.|
|name|text|Resource name.|
|type|text|Resource type.|
|location|text|Resource location.|
|tags|jsonb|Resource tags.|
