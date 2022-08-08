
# Table: azure_network_interfaces
Azure Network Interface
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|extended_location_name|text|The name of the extended location|
|extended_location_type|text|The type of the extended location|
|id|text|Resource ID.|
|location|text|Resource location.|
|name|text|Resource name.|
|dns_settings_applied_dns_servers|text[]|The servers that are part of the same availability set.|
|dns_settings_dns_servers|text[]|List of DNS servers IP addresses.|
|dns_settings_internal_dns_name_label|text|The internal dns name label.|
|dns_settings_internal_domain_name_suffix|text|The internal domain name suffix.|
|dns_settings_internal_fqdn|text|Fully qualified DNS name supporting internal communications between VMs in the same virtual network.|
|dscp_configuration_id|text|A reference to the dscp configuration to which the network interface is linked.|
|enable_accelerated_networking|boolean|If the network interface is accelerated networking enabled.|
|enable_ip_forwarding|boolean|Indicates whether IP forwarding is enabled on this network interface.|
|hosted_workloads|text[]|List of references to linked BareMetal resources.|
|mac_address|text|The MAC address of the network interface.|
|migration_phase|text|Migration phase of Network Interface resource.|
|network_security_group|text|The reference to the NetworkSecurityGroup resource.|
|nic_type|text|Type of Network Interface resource.|
|primary|boolean|Whether this is a primary network interface on a virtual machine.|
|private_endpoint|text|A reference to the private endpoint to which the network interface is linked.|
|private_link_service|jsonb|Privatelinkservice of the network interface resource.|
|provisioning_state|text|The provisioning state of the network interface resource.|
|resource_guid|text|The provisioning state of the network interface resource.|
|tap_configurations|jsonb|A list of TapConfigurations of the network interface.|
|virtual_machine_id|text|The reference to a virtual machine.|
|tags|jsonb|Resource tags.|
|type|text|Resource type.|
