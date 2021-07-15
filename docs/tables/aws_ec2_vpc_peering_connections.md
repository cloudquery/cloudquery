
# Table: aws_ec2_vpc_peering_connections
Describes a VPC peering connection.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|accepter_cidr_block|text|The IPv4 CIDR block for the VPC.|
|accepter_cidr_block_set|text[]|Information about the IPv4 CIDR blocks for the VPC.|
|accepter_ipv6_cidr_block_set|text[]|The IPv6 CIDR block for the VPC.|
|accepter_owner_id|text|The AWS account ID of the VPC owner.|
|accepter_allow_dns_resolution_from_remote_vpc|boolean|Indicates whether a local VPC can resolve public DNS hostnames to private IP addresses when queried from instances in a peer VPC.|
|accepter_allow_egress_local_classic_link_to_remote_vpc|boolean|Indicates whether a local ClassicLink connection can communicate with the peer VPC over the VPC peering connection.|
|accepter_allow_egress_local_vpc_to_remote_classic_link|boolean|Indicates whether a local VPC can communicate with a ClassicLink connection in the peer VPC over the VPC peering connection.|
|accepter_vpc_region|text|The Region in which the VPC is located.|
|accepter_vpc_id|text|The ID of the VPC.|
|expiration_time|timestamp without time zone|The time that an unaccepted VPC peering connection will expire.|
|requester_cidr_block|text|The IPv4 CIDR block for the VPC.|
|requester_cidr_block_set|text[]|Information about the IPv4 CIDR blocks for the VPC.|
|requester_ipv6_cidr_block_set|text[]|The IPv6 CIDR block for the VPC.|
|requester_owner_id|text|The AWS account ID of the VPC owner.|
|requester_allow_dns_resolution_from_remote_vpc|boolean|Indicates whether a local VPC can resolve public DNS hostnames to private IP addresses when queried from instances in a peer VPC.|
|requester_allow_egress_local_classic_link_to_remote_vpc|boolean|Indicates whether a local ClassicLink connection can communicate with the peer VPC over the VPC peering connection.|
|requester_allow_egress_local_vpc_to_remote_classic_link|boolean|Indicates whether a local VPC can communicate with a ClassicLink connection in the peer VPC over the VPC peering connection.|
|requester_vpc_region|text|The Region in which the VPC is located.|
|requester_vpc_id|text|The ID of the VPC.|
|status_code|text|The status of the VPC peering connection.|
|status_message|text|A message that provides more information about the status, if applicable.|
|tags|jsonb|Any tags assigned to the resource.|
|id|text|The ID of the VPC peering connection.|
