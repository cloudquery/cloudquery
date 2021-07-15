
# Table: aws_apigatewayv2_vpc_links
Represents a VPC link.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|id|text|The ID of the VPC link.|
|name|text|The name of the VPC link.|
|security_group_ids|text[]|A list of security group IDs for the VPC link.|
|subnet_ids|text[]|A list of subnet IDs to include in the VPC link.|
|vpc_link_id|text|The ID of the VPC link. (original field name)|
|created_date|timestamp without time zone|The timestamp when the VPC link was created.|
|tags|jsonb|Tags for the VPC link.|
|vpc_link_status|text|The status of the VPC link.|
|vpc_link_status_message|text|A message summarizing the cause of the status of the VPC link.|
|vpc_link_version|text|The version of the VPC link.|
