# Table: aws_ec2_vpc_endpoint_service_permissions

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AllowedPrincipal.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ec2_vpc_endpoint_services](aws_ec2_vpc_endpoint_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|principal|String|
|principal_type|String|
|service_id|String|
|service_permission_id|String|