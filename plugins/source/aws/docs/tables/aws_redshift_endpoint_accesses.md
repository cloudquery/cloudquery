# Table: aws_redshift_endpoint_accesses

This table shows data for Redshift Endpoint Accesses.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_EndpointAccess.html

The composite primary key for this table is (**cluster_arn**, **address**, **cluster_identifier**, **endpoint_name**).

## Relations

This table depends on [aws_redshift_clusters](aws_redshift_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn (PK)|`utf8`|
|address (PK)|`utf8`|
|cluster_identifier (PK)|`utf8`|
|endpoint_create_time|`timestamp[us, tz=UTC]`|
|endpoint_name (PK)|`utf8`|
|endpoint_status|`utf8`|
|port|`int64`|
|resource_owner|`utf8`|
|subnet_group_name|`utf8`|
|vpc_endpoint|`json`|
|vpc_security_groups|`json`|