
# Table: azure_network_security_group_default_security_rules
SecurityRule network security rule
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_cq_id|uuid|Unique CloudQuery ID of azure_network_security_groups table (FK)|
|description|text|A description for this rule Restricted to 140 chars|
|protocol|text|Network protocol this rule applies to Possible values include: 'SecurityRuleProtocolTCP', 'SecurityRuleProtocolUDP', 'SecurityRuleProtocolIcmp', 'SecurityRuleProtocolEsp', 'SecurityRuleProtocolAsterisk', 'SecurityRuleProtocolAh'|
|source_port_range|text|The source port or range Integer or range between 0 and 65535 Asterisk '*' can also be used to match all ports|
|destination_port_range|text|The destination port or range Integer or range between 0 and 65535 Asterisk '*' can also be used to match all ports|
|source_address_prefix|text|The CIDR or source IP range Asterisk '*' can also be used to match all source IPs Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used If this is an ingress rule, specifies where network traffic originates from|
|source_address_prefixes|text[]|The CIDR or source IP ranges|
|destination_address_prefix|text|The destination address prefix CIDR or destination IP range Asterisk '*' can also be used to match all source IPs Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used|
|destination_address_prefixes|text[]|The destination address prefixes CIDR or destination IP ranges|
|source_port_ranges|text[]|The source port ranges|
|destination_port_ranges|text[]|The destination port ranges|
|access|text|The network traffic is allowed or denied Possible values include: 'SecurityRuleAccessAllow', 'SecurityRuleAccessDeny'|
|priority|integer|The priority of the rule The value can be between 100 and 4096 The priority number must be unique for each rule in the collection The lower the priority number, the higher the priority of the rule|
|direction|text|The direction of the rule The direction specifies if rule will be evaluated on incoming or outgoing traffic Possible values include: 'SecurityRuleDirectionInbound', 'SecurityRuleDirectionOutbound'|
|provisioning_state|text|The provisioning state of the security rule resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'|
|name|text|The name of the resource that is unique within a resource group This name can be used to access the resource|
|etag|text|A unique read-only string that changes whenever the resource is updated|
|type|text|The type of the resource|
|id|text|Resource ID|
