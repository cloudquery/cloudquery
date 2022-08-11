
# Table: azure_storage_account_network_rule_set_virtual_network_rules
VirtualNetworkRule virtual Network rule. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_cq_id|uuid|Unique ID of azure_storage_accounts table (FK)|
|virtual_network_resource_id|text|Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/MicrosoftNetwork/virtualNetworks/{vnetName}/subnets/{subnetName}|
|action|text|The action of virtual network rule.|
|state|text|Gets the state of virtual network rule Possible values include: 'StateProvisioning', 'StateDeprovisioning', 'StateSucceeded', 'StateFailed', 'StateNetworkSourceDeleted'|
