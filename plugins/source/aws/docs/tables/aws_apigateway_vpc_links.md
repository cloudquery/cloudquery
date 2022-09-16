
# Table: aws_apigateway_vpc_links
An API Gateway VPC link for a RestApi to access resources in an Amazon Virtual Private Cloud (VPC)
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource|
|region|text|The AWS Region of the resource|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|description|text|The description of the VPC link|
|id|text|The identifier of the VpcLink|
|name|text|The name used to label and identify the VPC link|
|status|text|The status of the VPC link|
|status_message|text|A description about the VPC link status|
|tags|jsonb|The collection of tags|
|target_arns|text[]|The ARN of the network load balancer of the VPC targeted by the VPC link|
