
# Table: aws_directconnect_virtual_interface_bgp_peers
Information about a BGP peer. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_interface_id|uuid|Unique ID of aws_directconnect_virtual_interfaces table (FK)|
|address_family|text|The address family for the BGP peer.|
|amazon_address|text|The IP address assigned to the Amazon interface.|
|asn|integer|The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.|
|auth_key|text|The authentication key for BGP configuration|
|aws_device_v2|text|The Direct Connect endpoint on which the BGP peer terminates.|
|bgp_peer_id|text|The ID of the BGP peer.|
|bgp_peer_state|text|The state of the BGP peer|
|bgp_status|text|The status of the BGP peer|
|customer_address|text|The IP address assigned to the customer interface.|
