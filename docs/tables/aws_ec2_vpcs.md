
# Table: aws_ec2_vpcs
Describes a VPC.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|cidr_block|text|The primary IPv4 CIDR block for the VPC.|
|dhcp_options_id|text|The ID of the set of DHCP options you've associated with the VPC.|
|instance_tenancy|text|The allowed tenancy of instances launched into the VPC.|
|is_default|boolean|Indicates whether the VPC is the default VPC.|
|owner_id|text|The ID of the AWS account that owns the VPC.|
|state|text|The current state of the VPC.|
|tags|jsonb|Any tags assigned to the VPC.|
|id|text|The ID of the VPC.|
