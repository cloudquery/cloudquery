
# Table: aws_ec2_eips
Describes an Elastic IP address, or a carrier IP address.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|allocation_id|text|The ID representing the allocation of the address for use with EC2-VPC.|
|association_id|text|The ID representing the association of the address with an instance in a VPC.|
|carrier_ip|inet|The carrier IP address associated|
|customer_owned_ip|inet|The customer-owned IP address.|
|customer_owned_ipv4_pool|text|The ID of the customer-owned address pool.|
|domain|text|Indicates whether this Elastic IP address is for use with instances in EC2-Classic (standard) or instances in a VPC (vpc).|
|instance_id|text|The ID of the instance that the address is associated with (if any).|
|network_border_group|text|The name of the unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses.|
|network_interface_id|text|The ID of the network interface.|
|network_interface_owner_id|text|The ID of the AWS account that owns the network interface.|
|private_ip_address|inet|The private IP address associated with the Elastic IP address.|
|public_ip|inet|The Elastic IP address.|
|public_ipv4_pool|text|The ID of an address pool.|
|tags|jsonb|Any tags assigned to the Elastic IP address.|
