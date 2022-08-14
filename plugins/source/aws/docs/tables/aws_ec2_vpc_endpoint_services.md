
# Table: aws_ec2_vpc_endpoint_services
Describes a VPC endpoint service.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|acceptance_required|boolean|Indicates whether VPC endpoint connection requests to the service must be accepted by the service owner.|
|availability_zones|text[]|The Availability Zones in which the service is available.|
|base_endpoint_dns_names|text[]|The DNS names for the service.|
|manages_vpc_endpoints|boolean|Indicates whether the service manages its VPC endpoints.|
|owner|text|The Amazon Web Services account ID of the service owner.|
|payer_responsibility|text|The payer responsibility.|
|private_dns_name|text|The private DNS name for the service.|
|private_dns_name_verification_state|text|The verification state of the VPC endpoint service.|
|private_dns_names|text[]|The private DNS names assigned to the VPC endpoint service.|
|id|text|The ID of the endpoint service.|
|service_name|text|The Amazon Resource Name (ARN) of the service.|
|service_type|text[]|The type of service.|
|tags|jsonb|Any tags assigned to the service.|
|vpc_endpoint_policy_supported|boolean|Indicates whether the service supports endpoint policies.|
