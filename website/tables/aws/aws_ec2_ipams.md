# Table: aws_ec2_ipams

This table shows data for Amazon Elastic Compute Cloud (EC2) Ipams.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Ipam.html

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
|default_resource_discovery_association_id|`utf8`|
|default_resource_discovery_id|`utf8`|
|description|`utf8`|
|ipam_arn|`utf8`|
|ipam_id|`utf8`|
|ipam_region|`utf8`|
|operating_regions|`json`|
|owner_id|`utf8`|
|private_default_scope_id|`utf8`|
|public_default_scope_id|`utf8`|
|resource_discovery_association_count|`int64`|
|scope_count|`int64`|
|state|`utf8`|