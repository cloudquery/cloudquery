
# Table: aws_ec2_vpc_endpoint_service_configurations
Describes a service configuration for a VPC endpoint service.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|acceptance_required|boolean|Indicates whether requests from other AWS accounts to create an endpoint to the service must first be accepted.|
|availability_zones|text[]|The Availability Zones in which the service is available.|
|base_endpoint_dns_names|text[]|The DNS names for the service.|
|gateway_load_balancer_arns|text[]|The Amazon Resource Names (ARNs) of the Gateway Load Balancers for the service.|
|manages_vpc_endpoints|boolean|Indicates whether the service manages its VPC endpoints.|
|network_load_balancer_arns|text[]|The Amazon Resource Names (ARNs) of the Network Load Balancers for the service.|
|payer_responsibility|text|The payer responsibility.|
|private_dns_name|text|The private DNS name for the service.|
|private_dns_name_configuration_name|text|The name of the record subdomain the service provider needs to create.|
|private_dns_name_configuration_state|text|The verification state of the VPC endpoint service.|
|private_dns_name_configuration_type|text|The endpoint service verification type, for example TXT.|
|private_dns_name_configuration_value|text|The value the service provider adds to the private DNS name domain record before verification.|
|service_id|text|The ID of the service.|
|service_name|text|The name of the service.|
|service_state|text|The service state.|
|service_type|text[]|The type of service.|
|tags|jsonb|Any tags assigned to the service.|
