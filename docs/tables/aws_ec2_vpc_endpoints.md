
# Table: aws_ec2_vpc_endpoints
Describes a VPC endpoint.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|creation_timestamp|timestamp without time zone|The date and time that the VPC endpoint was created.|
|last_error_code|text|The error code for the VPC endpoint error.|
|last_error_message|text|The error message for the VPC endpoint error.|
|network_interface_ids|text[]|(Interface endpoint) One or more network interfaces for the endpoint.|
|owner_id|text|The ID of the AWS account that owns the VPC endpoint.|
|policy_document|text|The policy document associated with the endpoint, if applicable.|
|private_dns_enabled|boolean|(Interface endpoint) Indicates whether the VPC is associated with a private hosted zone.|
|requester_managed|boolean|Indicates whether the VPC endpoint is being managed by its service.|
|route_table_ids|text[]|(Gateway endpoint) One or more route tables associated with the endpoint.|
|service_name|text|The name of the service to which the endpoint is associated.|
|state|text|The state of the VPC endpoint.|
|subnet_ids|text[]|(Interface endpoint) One or more subnets in which the endpoint is located.|
|tags|jsonb|Any tags assigned to the VPC endpoint.|
|vpc_endpoint_id|text|The ID of the VPC endpoint.|
|vpc_endpoint_type|text|The type of endpoint.|
|vpc_id|text|The ID of the VPC to which the endpoint is associated.|
