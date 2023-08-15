# Table: aws_ec2_vpc_endpoint_services

This table shows data for Amazon Elastic Compute Cloud (EC2) VPC Endpoint Services.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ServiceDetail.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ec2_vpc_endpoint_services:
  - [aws_ec2_vpc_endpoint_service_permissions](aws_ec2_vpc_endpoint_service_permissions)

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
|manages_vpc_endpoints|`bool`|
|owner|`utf8`|
|payer_responsibility|`utf8`|
|private_dns_name|`utf8`|
|private_dns_name_verification_state|`utf8`|
|private_dns_names|`json`|
|service_id|`utf8`|
|service_name|`utf8`|
|service_type|`json`|
|supported_ip_address_types|`list<item: utf8, nullable>`|
|vpc_endpoint_policy_supported|`bool`|