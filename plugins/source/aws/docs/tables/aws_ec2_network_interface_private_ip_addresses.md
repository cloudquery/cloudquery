
# Table: aws_ec2_network_interface_private_ip_addresses
Describes the private IPv4 address of a network interface.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|network_interface_cq_id|uuid|Unique CloudQuery ID of aws_ec2_network_interfaces table (FK)|
|association_allocation_id|text|The allocation ID.|
|association_id|text|The association ID.|
|association_carrier_ip|text|The carrier IP address associated with the network interface|
|association_customer_owned_ip|text|The customer-owned IP address associated with the network interface.|
|association_ip_owner_id|text|The ID of the Elastic IP address owner.|
|association_public_dns_name|text|The public DNS name.|
|association_public_ip|text|The address of the Elastic IP address bound to the network interface.|
|primary|boolean|Indicates whether this IPv4 address is the primary private IPv4 address of the network interface.|
|private_dns_name|text|The private DNS name.|
|private_ip_address|text|The private IPv4 address.|
