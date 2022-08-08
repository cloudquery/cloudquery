
# Table: azure_container_registry_network_rule_set_ip_rules
IPRule IP rule with specific IP or IP range in CIDR format
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|registry_cq_id|uuid|Unique CloudQuery ID of azure_container_registries table (FK)|
|action|text|The action of IP ACL rule|
|ip_address_or_range|cidr|Specifies the IP or IP range in CIDR format|
