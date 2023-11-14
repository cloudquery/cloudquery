# Table: aws_ec2_vpc_endpoint_connections

This table shows data for Amazon Elastic Compute Cloud (EC2) VPC Endpoint Connections.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcEndpointConnection.html

The composite primary key for this table is (**account_id**, **region**, **vpc_endpoint_connection_id**, **vpc_endpoint_owner**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|creation_timestamp|`timestamp[us, tz=UTC]`|
|dns_entries|`json`|
|gateway_load_balancer_arns|`list<item: utf8, nullable>`|
|ip_address_type|`utf8`|
|network_load_balancer_arns|`list<item: utf8, nullable>`|
|service_id|`utf8`|
|vpc_endpoint_connection_id (PK)|`utf8`|
|vpc_endpoint_id|`utf8`|
|vpc_endpoint_owner (PK)|`utf8`|
|vpc_endpoint_state|`utf8`|