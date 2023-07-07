# Table: aws_ec2_vpc_endpoint_service_permissions

This table shows data for Amazon Elastic Compute Cloud (EC2) VPC Endpoint Service Permissions.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AllowedPrincipal.html

The composite primary key for this table is (**account_id**, **service_id**, **service_permission_id**).

## Relations

This table depends on [aws_ec2_vpc_endpoint_services](aws_ec2_vpc_endpoint_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|tags|`json`|
|principal|`utf8`|
|principal_type|`utf8`|
|service_id (PK)|`utf8`|
|service_permission_id (PK)|`utf8`|