
# Table: azure_storage_account_network_rule_set_ip_rules
IPRule IP rule with specific IP or IP range in CIDR format. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|uuid|Unique ID of azure_storage_accounts table (FK)|
|ip_address_or_range|text|Specifies the IP or IP range in CIDR format Only IPV4 address is allowed|
|action|text|The action of IP ACL rule Possible values include: 'Allow'|
