
# Table: azure_iothub_hub_network_rule_sets_ip_rules
NetworkRuleSetIPRule IP Rule to be applied as part of Network Rule Set
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hub_cq_id|uuid|Unique CloudQuery ID of azure_iothub_hubs table (FK)|
|filter_name|text|Name of the IP filter rule.|
|action|text|IP Filter Action|
|ip_mask|text|A string that contains the IP address range in CIDR notation for the rule.|
