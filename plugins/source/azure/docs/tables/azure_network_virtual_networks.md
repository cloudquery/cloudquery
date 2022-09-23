# Table: azure_network_virtual_networks


The primary key for this table is **id**.

## Relations
The following tables depend on `azure_network_virtual_networks`:
  - [`azure_network_virtual_network_gateways`](azure_network_virtual_network_gateways.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|extended_location|JSON|
|address_space|JSON|
|dhcp_options|JSON|
|subnets|JSON|
|virtual_network_peerings|JSON|
|resource_guid|String|
|provisioning_state|String|
|enable_ddos_protection|Bool|
|enable_vm_protection|Bool|
|ddos_protection_plan|JSON|
|bgp_communities|JSON|
|ip_allocations|JSON|
|etag|String|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|