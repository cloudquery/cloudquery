
# Table: aws_apigateway_vpc_links
An API Gateway VPC link for a RestApi to access resources in an Amazon Virtual Private Cloud (VPC).
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|description|text|The description of the VPC link.|
|resource_id|text|The identifier of the VpcLink. It is used in an Integration to reference this VpcLink.|
|name|text|The name used to label and identify the VPC link.|
|status|text|The status of the VPC link. The valid values are AVAILABLE, PENDING, DELETING, or FAILED. Deploying an API will wait if the status is PENDING and will fail if the status is DELETING.|
|status_message|text|A description about the VPC link status.|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
|target_arns|text[]|The ARN of the network load balancer of the VPC targeted by the VPC link. The network load balancer must be owned by the same AWS account of the API owner.|
