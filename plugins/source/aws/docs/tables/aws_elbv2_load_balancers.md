
# Table: aws_elbv2_load_balancers
Information about a load balancer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|web_acl_arn|text|arn of associated web acl|
|tags|jsonb||
|canonical_hosted_zone_id|text|The ID of the Amazon Route 53 hosted zone associated with the load balancer.|
|created_time|timestamp without time zone|The date and time the load balancer was created.|
|customer_owned_ipv4_pool|text|[Application Load Balancers on Outposts] The ID of the customer-owned address pool.|
|dns_name|text|The public DNS name of the load balancer.|
|ip_address_type|text|The type of IP addresses used by the subnets for your load balancer|
|arn|text|The Amazon Resource Name (ARN) of the load balancer.|
|name|text|The name of the load balancer.|
|scheme|text|The nodes of an Internet-facing load balancer have public IP addresses|
|security_groups|text[]|The IDs of the security groups for the load balancer.|
|state_code|text|The state code|
|state_reason|text|A description of the state.|
|type|text|The type of load balancer.|
|vpc_id|text|The ID of the VPC for the load balancer.|
