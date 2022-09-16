
# Table: aws_ec2_security_group_ip_permissions
Describes a set of permissions for a security group rule.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_cq_id|uuid|Unique CloudQuery ID of aws_ec2_security_groups table (FK)|
|from_port|integer|The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number.|
|ip_protocol|text|The IP protocol name (tcp, udp, icmp, icmpv6) or number|
|to_port|integer|The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.|
|permission_type|text|egress or ingress|
