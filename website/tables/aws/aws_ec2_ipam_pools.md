# Table: aws_ec2_ipam_pools

This table shows data for Amazon Elastic Compute Cloud (EC2) Amazon VPC IP Address Manager (IPAM) Pools.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IpamPool.html

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
|address_family|`utf8`|
|allocation_default_netmask_length|`int64`|
|allocation_max_netmask_length|`int64`|
|allocation_min_netmask_length|`int64`|
|allocation_resource_tags|`json`|
|auto_import|`bool`|
|aws_service|`utf8`|
|description|`utf8`|
|ipam_arn|`utf8`|
|ipam_pool_arn|`utf8`|
|ipam_pool_id|`utf8`|
|ipam_region|`utf8`|
|ipam_scope_arn|`utf8`|
|ipam_scope_type|`utf8`|
|locale|`utf8`|
|owner_id|`utf8`|
|pool_depth|`int64`|
|public_ip_source|`utf8`|
|publicly_advertisable|`bool`|
|source_ipam_pool_id|`utf8`|
|state|`utf8`|
|state_message|`utf8`|