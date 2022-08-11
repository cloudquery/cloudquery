
# Table: azure_container_registry_network_rule_set_virtual_network_rules
VirtualNetworkRule virtual network rule
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|registry_cq_id|uuid|Unique CloudQuery ID of azure_container_registries table (FK)|
|action|text|The action of virtual network rule|
|virtual_network_id|text|Resource ID of a subnet|
