# Table: aws_ec2_vpc_endpoint_service_permissions

This table shows data for Amazon Elastic Compute Cloud (EC2) VPC Endpoint Service Permissions.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AllowedPrincipal.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **service_id**, **service_permission_id**).
## Relations

This table depends on [aws_ec2_vpc_endpoint_services](aws_ec2_vpc_endpoint_services.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|principal|`utf8`|
|principal_type|`utf8`|
|service_id|`utf8`|
|service_permission_id|`utf8`|