
# Table: aws_ec2_instance_network_interfaces
Describes a network interface.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_ec2_instances table (FK)|
|association_carrier_ip|text|The carrier IP address associated with the network interface.|
|association_ip_owner_id|text|The ID of the owner of the Elastic IP address.|
|association_public_dns_name|text|The public DNS name.|
|association_public_ip|text|The public IP address or Elastic IP address bound to the network interface.|
|attachment_attach_time|timestamp without time zone|The time stamp when the attachment initiated.|
|attachment_id|text|The ID of the network interface attachment.|
|attachment_delete_on_termination|boolean|Indicates whether the network interface is deleted when the instance is terminated.|
|attachment_device_index|integer|The index of the device on the instance for the network interface attachment.|
|attachment_network_card_index|integer|The index of the network card.|
|attachment_status|text|The attachment state.|
|description|text|The description.|
|interface_type|text|Describes the type of network interface|
|ipv4_prefixes|text[]|The IPv4 delegated prefixes that are assigned to the network interface.|
|ipv6_prefixes|text[]|The IPv6 delegated prefixes that are assigned to the network interface.|
|mac_address|text|The MAC address.|
|network_interface_id|text|The ID of the network interface.|
|owner_id|text|The ID of the Amazon Web Services account that created the network interface.|
|private_dns_name|text|The private DNS name.|
|private_ip_address|text|The IPv4 address of the network interface within the subnet.|
|source_dest_check|boolean|Indicates whether source/destination checking is enabled.|
|status|text|The status of the network interface.|
|subnet_id|text|The ID of the subnet.|
|vpc_id|text|The ID of the VPC.|
