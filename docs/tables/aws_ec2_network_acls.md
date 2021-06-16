
# Table: aws_ec2_network_acls
Describes a network ACL.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|is_default|boolean|Indicates whether this is the default network ACL for the VPC.|
|network_acl_id|text|The ID of the network ACL.|
|owner_id|text|The ID of the AWS account that owns the network ACL.|
|tags|jsonb|Any tags assigned to the network ACL.|
|vpc_id|text|The ID of the VPC for the network ACL.|
