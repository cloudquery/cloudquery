# Table: aws_directconnect_virtual_interfaces

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_VirtualInterface.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|id|String|
|tags|JSON|
|address_family|String|
|amazon_address|String|
|amazon_side_asn|Int|
|asn|Int|
|auth_key|String|
|aws_device_v2|String|
|aws_logical_device_id|String|
|bgp_peers|JSON|
|connection_id|String|
|customer_address|String|
|customer_router_config|String|
|direct_connect_gateway_id|String|
|jumbo_frame_capable|Bool|
|location|String|
|mtu|Int|
|owner_account|String|
|route_filter_prefixes|JSON|
|site_link_enabled|Bool|
|virtual_gateway_id|String|
|virtual_interface_id|String|
|virtual_interface_name|String|
|virtual_interface_state|String|
|virtual_interface_type|String|
|vlan|Int|