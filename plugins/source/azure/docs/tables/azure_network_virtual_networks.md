# Table: azure_network_virtual_networks

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2#VirtualNetwork

The primary key for this table is **id**.

## Relations

The following tables depend on azure_network_virtual_networks:
  - [azure_network_virtual_network_gateways](azure_network_virtual_network_gateways.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|extended_location|JSON|
|id (PK)|String|
|location|String|
|address_space|JSON|
|bgp_communities|JSON|
|ddos_protection_plan|JSON|
|dhcp_options|JSON|
|enable_ddos_protection|Bool|
|enable_vm_protection|Bool|
|encryption|JSON|
|flow_timeout_in_minutes|Int|
|ip_allocations|JSON|
|subnets|JSON|
|virtual_network_peerings|JSON|
|provisioning_state|String|
|resource_guid|String|
|tags|JSON|
|etag|String|
|name|String|
|type|String|