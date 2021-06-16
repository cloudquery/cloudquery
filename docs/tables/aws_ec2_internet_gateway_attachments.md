
# Table: aws_ec2_internet_gateway_attachments
Describes the attachment of a VPC to an internet gateway or an egress-only internet gateway.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|internet_gateway_id|uuid|Unique ID of aws_ec2_internet_gateways table (FK)|
|state|text|The current state of the attachment.|
|vpc_id|text|The ID of the VPC.|
