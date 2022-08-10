
# Table: aws_elbv2_load_balancer_availability_zone_addresses
Information about a static IP address for a load balancer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_availability_zone_cq_id|uuid|Unique CloudQuery ID of aws_elbv2_load_balancer_availability_zones table (FK)|
|zone_name|text|The name of the Availability Zone..|
|allocation_id|text|[Network Load Balancers] The allocation ID of the Elastic IP address for an internal-facing load balancer.|
|ipv6_address|text|[Network Load Balancers] The IPv6 address.|
|ip_address|text|The static IP address.|
|private_ipv4_address|text|[Network Load Balancers] The private IPv4 address for an internal load balancer.|
