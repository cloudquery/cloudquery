
# Table: aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces
Describes a network interface.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_endpoint_vpc_endpoint_cq_id|uuid|Unique CloudQuery ID of aws_redshift_cluster_endpoint_vpc_endpoints table (FK)|
|availability_zone|text|The Availability Zone.|
|network_interface_id|text|The network interface identifier.|
|private_ip_address|text|The IPv4 address of the network interface within the subnet.|
|subnet_id|text|The subnet identifier.|
