# Table: aws_ec2_vpc_endpoint_service_permissions

This table shows data for AWS Ec2 VPC Endpoint Service Permissions.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AllowedPrincipal.html

The composite primary key for this table is (**account_id**, **service_id**, **service_permission_id**).

## Relations

This table depends on [aws_ec2_vpc_endpoint_services](aws_ec2_vpc_endpoint_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|tags|JSON|
|principal|String|
|principal_type|String|
|service_id (PK)|String|
|service_permission_id (PK)|String|