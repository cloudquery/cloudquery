# Table: aws_redshift_endpoint_access

This table shows data for Redshift Endpoint Access.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_EndpointAccess.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_redshift_clusters](aws_redshift_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn|`utf8`|
|address|`utf8`|
|cluster_identifier|`utf8`|
|endpoint_create_time|`timestamp[us, tz=UTC]`|
|endpoint_name|`utf8`|
|endpoint_status|`utf8`|
|port|`int64`|
|resource_owner|`utf8`|
|subnet_group_name|`utf8`|
|vpc_endpoint|`json`|
|vpc_security_groups|`json`|