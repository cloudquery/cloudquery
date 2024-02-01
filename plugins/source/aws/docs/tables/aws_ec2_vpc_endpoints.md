# Table: aws_ec2_vpc_endpoints

This table shows data for Amazon Elastic Compute Cloud (EC2) VPC Endpoints.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcEndpoint.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|creation_timestamp|`timestamp[us, tz=UTC]`|
|dns_entries|`json`|
|dns_options|`json`|
|groups|`json`|
|ip_address_type|`utf8`|
|last_error|`json`|
|network_interface_ids|`list<item: utf8, nullable>`|
|owner_id|`utf8`|
|policy_document|`utf8`|
|private_dns_enabled|`bool`|
|requester_managed|`bool`|
|route_table_ids|`list<item: utf8, nullable>`|
|service_name|`utf8`|
|state|`utf8`|
|subnet_ids|`list<item: utf8, nullable>`|
|vpc_endpoint_id|`utf8`|
|vpc_endpoint_type|`utf8`|
|vpc_id|`utf8`|