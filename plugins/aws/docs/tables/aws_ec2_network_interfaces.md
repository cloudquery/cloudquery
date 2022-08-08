
# Table: aws_ec2_network_interfaces
Describes a network interface.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the egress-only internet gateway.|
|tags|jsonb|Any tags assigned to the network interface.|
|association_allocation_id|text|The allocation ID.|
|association_id|text|The association ID.|
|association_carrier_ip|text|The carrier IP address associated with the network interface|
|association_customer_owned_ip|text|The customer-owned IP address associated with the network interface.|
|association_ip_owner_id|text|The ID of the Elastic IP address owner.|
|association_public_dns_name|text|The public DNS name.|
|association_public_ip|text|The address of the Elastic IP address bound to the network interface.|
|attachment_attach_time|timestamp without time zone|The timestamp indicating when the attachment initiated.|
|attachment_id|text|The ID of the network interface attachment.|
|attachment_delete_on_termination|boolean|Indicates whether the network interface is deleted when the instance is terminated.|
|attachment_device_index|integer|The device index of the network interface attachment on the instance.|
|attachment_instance_id|text|The ID of the instance.|
|attachment_instance_owner_id|text|The Amazon Web Services account ID of the owner of the instance.|
|attachment_network_card_index|integer|The index of the network card.|
|attachment_status|text|The attachment state.|
|availability_zone|text|The Availability Zone.|
|deny_all_igw_traffic|boolean|Indicates whether a network interface with an IPv6 address is unreachable from the public internet|
|description|text|A description.|
|groups|jsonb|The tags assigned to the egress-only internet gateway.|
|interface_type|text|The type of network interface.|
|ipv4_prefixes|text[]|Describes an IPv4 prefix.|
|ipv6_address|text|The IPv6 globally unique address associated with the network interface.|
|ipv6_addresses|text[]|Describes an IPv6 address associated with a network interface.|
|ipv6_native|boolean|Indicates whether this is an IPv6 only network interface.|
|ipv6_prefixes|text[]|Describes the IPv6 prefix.|
|mac_address|text|The MAC address.|
|id|text|The ID of the network interface.|
|outpost_arn|text|The Amazon Resource Name (ARN) of the Outpost.|
|owner_id|text|The Amazon Web Services account ID of the owner of the network interface.|
|private_dns_name|text|The private DNS name.|
|private_ip_address|text|The IPv4 address of the network interface within the subnet.|
|requester_id|text|The alias or Amazon Web Services account ID of the principal or service that created the network interface.|
|requester_managed|boolean|Indicates whether the network interface is being managed by Amazon Web Services.|
|source_dest_check|boolean|Indicates whether source/destination checking is enabled.|
|status|text|The status of the network interface.|
|subnet_id|text|The ID of the subnet.|
|vpc_id|text|The ID of the VPC.|
