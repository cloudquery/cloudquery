
# Table: azure_network_public_ip_addresses
PublicIPAddress public IP address resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|extended_location_name|text|The name of the extended location|
|extended_location_type|text|The type of the extended location|
|sku_name|text|Name of a public IP address SKU Possible values include: 'PublicIPAddressSkuNameBasic', 'PublicIPAddressSkuNameStandard'|
|sku_tier|text|Tier of a public IP address SKU Possible values include: 'PublicIPAddressSkuTierRegional', 'PublicIPAddressSkuTierGlobal'|
|public_ip_allocation_method|text|The public IP address allocation method Possible values include: 'Static', 'Dynamic'|
|public_ip_address_version|text|The public IP address version Possible values include: 'IPv4', 'IPv6'|
|private_ip_address|text|The private IP address of the IP configuration|
|private_ip_allocation_method|text|The private IP address allocation method Possible values include: 'Static', 'Dynamic'|
|subnet|jsonb|The reference to the subnet resource|
|public_ip_address|jsonb|The reference to the public IP resource|
|dns_settings_domain_name_label|text|The domain name label The concatenation of the domain name label and the regionalized DNS zone make up the fully qualified domain name associated with the public IP address If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system|
|dns_settings_fqdn|text|The Fully Qualified Domain Name of the A DNS record associated with the public IP This is the concatenation of the domainNameLabel and the regionalized DNS zone|
|dns_settings_reverse_fqdn|text|The reverse FQDN A user-visible, fully qualified domain name that resolves to this public IP address If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addrarpa domain to the reverse FQDN|
|ddos_settings_ddos_custom_policy_id|text|Resource ID|
|ddos_settings_protection_coverage|text|The DDoS protection policy customizability of the public IP Only standard coverage will have the ability to be customized Possible values include: 'DdosSettingsProtectionCoverageBasic', 'DdosSettingsProtectionCoverageStandard'|
|ddos_settings_protected_ip|boolean|Enables DDoS protection on the public IP|
|ip_address|text|The IP address associated with the public IP address resource|
|public_ip_prefix_id|text|Resource ID|
|idle_timeout_in_minutes|integer|The idle timeout of the public IP address|
|resource_guid|text|The resource GUID property of the public IP address resource|
|provisioning_state|text|The provisioning state of the public IP address resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'|
|etag|text|A unique read-only string that changes whenever the resource is updated|
|zones|text[]|A list of availability zones denoting the IP allocated for the resource needs to come from|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
