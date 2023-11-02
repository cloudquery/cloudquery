# Table: aws_ec2_instance_connect_endpoints

This table shows data for Amazon Elastic Compute Cloud (EC2) Instance Connect Endpoints.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Ec2InstanceConnectEndpoint.html
The 'request_account_id' and 'request_region' columns are added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|availability_zone|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|dns_name|`utf8`|
|fips_dns_name|`utf8`|
|instance_connect_endpoint_arn|`utf8`|
|instance_connect_endpoint_id|`utf8`|
|network_interface_ids|`list<item: utf8, nullable>`|
|owner_id|`utf8`|
|preserve_client_ip|`bool`|
|security_group_ids|`list<item: utf8, nullable>`|
|state|`utf8`|
|state_message|`utf8`|
|subnet_id|`utf8`|
|vpc_id|`utf8`|