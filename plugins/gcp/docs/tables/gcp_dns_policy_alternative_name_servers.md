
# Table: gcp_dns_policy_alternative_name_servers

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|policy_cq_id|uuid|Unique ID of gcp_dns_policies table (FK)|
|policy_id|text||
|forwarding_path|text|Forwarding path for this TargetNameServer If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target  Possible values:   "default" - Cloud DNS makes forwarding decision based on IP address ranges; that is, RFC1918 addresses forward to the target through the VPC and non-RFC1918 addresses forward to the target through the internet   "private" - Cloud DNS always forwards to this target through the VPC|
|ipv4_address|text|IPv4 address to forward to|
|kind|text|The resource type|
