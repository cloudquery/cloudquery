
# Table: aws_ec2_byoip_cidrs
Information about an address range that is provisioned for use with your AWS resources through bring your own IP addresses (BYOIP).
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|cidr|text|The address range, in CIDR notation.|
|description|text|The description of the address range.|
|state|text|The state of the address pool.|
|status_message|text|Upon success, contains the ID of the address pool.|
