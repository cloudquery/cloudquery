# Table: aws_ec2_vpc_endpoint_service_configurations

This table shows data for Amazon Elastic Compute Cloud (EC2) VPC Endpoint Service Configurations.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ServiceConfiguration.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|acceptance_required|`bool`|
|availability_zones|`list<item: utf8, nullable>`|
|base_endpoint_dns_names|`list<item: utf8, nullable>`|
|gateway_load_balancer_arns|`list<item: utf8, nullable>`|
|manages_vpc_endpoints|`bool`|
|network_load_balancer_arns|`list<item: utf8, nullable>`|
|payer_responsibility|`utf8`|
|private_dns_name|`utf8`|
|private_dns_name_configuration|`json`|
|service_id|`utf8`|
|service_name|`utf8`|
|service_state|`utf8`|
|service_type|`json`|
|supported_ip_address_types|`list<item: utf8, nullable>`|