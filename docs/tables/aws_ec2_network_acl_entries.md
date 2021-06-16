
# Table: aws_ec2_network_acl_entries
Describes an entry in a network ACL.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|network_acl_id|uuid|Unique ID of aws_ec2_network_acls table (FK)|
|cidr_block|text|The IPv4 network range to allow or deny, in CIDR notation.|
|egress|boolean|Indicates whether the rule is an egress rule (applied to traffic leaving the subnet).|
|icmp_type_code|integer|The ICMP code.|
|icmp_type_code_type|integer|The ICMP type.|
|ipv6_cidr_block|text|The IPv6 network range to allow or deny, in CIDR notation.|
|port_range_from|integer|The first port in the range.|
|port_range_to|integer|The last port in the range.|
|protocol|text|The protocol number.|
|rule_action|text|Indicates whether to allow or deny the traffic that matches the rule.|
|rule_number|integer|The rule number for the entry.|
