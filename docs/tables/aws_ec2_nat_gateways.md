
# Table: aws_ec2_nat_gateways
Describes a NAT gateway.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|id|text|The ID of the NAT gateway.|
|create_time|timestamp without time zone|The date and time the NAT gateway was created.|
|delete_time|timestamp without time zone|The date and time the NAT gateway was deleted, if applicable.|
|failure_code|text|If the NAT gateway could not be created, specifies the error code for the failure.|
|failure_message|text|If the NAT gateway could not be created, specifies the error message for the failure, that corresponds to the error code.|
|provisioned_bandwidth_provision_time|timestamp without time zone|Reserved.|
|provisioned_bandwidth_provisioned|text|Reserved.|
|provisioned_bandwidth_request_time|timestamp without time zone|Reserved.|
|provisioned_bandwidth_requested|text|Reserved.|
|provisioned_bandwidth_status|text|Reserved.|
|state|text|The state of the NAT gateway.|
|subnet_id|text|The ID of the subnet in which the NAT gateway is located.|
|tags|jsonb|The tags for the NAT gateway.|
|vpc_id|text|The ID of the VPC in which the NAT gateway is located.|
