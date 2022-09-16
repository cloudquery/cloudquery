
# Table: aws_ec2_nat_gateway_addresses
Describes the IP addresses and network interface associated with a NAT gateway.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|nat_gateway_cq_id|uuid|Unique CloudQuery ID of aws_ec2_nat_gateways table (FK)|
|allocation_id|text|The allocation ID of the Elastic IP address that's associated with the NAT gateway.|
|network_interface_id|text|The ID of the network interface associated with the NAT gateway.|
|private_ip|text|The private IP address associated with the Elastic IP address.|
|public_ip|text|The Elastic IP address associated with the NAT gateway.|
