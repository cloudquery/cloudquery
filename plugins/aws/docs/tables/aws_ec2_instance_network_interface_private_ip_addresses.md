
# Table: aws_ec2_instance_network_interface_private_ip_addresses
Describes a private IPv4 address.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_network_interface_id|uuid|Unique CloudQuery ID of aws_ec2_instance_network_interfaces table (FK)|
|association_carrier_ip|text|The carrier IP address associated with the network interface.|
|association_ip_owner_id|text|The ID of the owner of the Elastic IP address.|
|association_public_dns_name|text|The public DNS name.|
|association_public_ip|text|The public IP address or Elastic IP address bound to the network interface.|
|is_primary|boolean|Indicates whether this IPv4 address is the primary private IP address of the network interface.|
|private_dns_name|text|The private IPv4 DNS name.|
|private_ip_address|text|The private IPv4 address of the network interface.|
