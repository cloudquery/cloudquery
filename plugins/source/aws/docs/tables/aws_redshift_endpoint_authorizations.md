# Table: aws_redshift_endpoint_authorizations

This table shows data for Redshift Endpoint Authorizations.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_EndpointAuthorization.html

The composite primary key for this table is (**cluster_arn**, **cluster_identifier**, **grantee**, **grantor**).

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
|allowed_all_vp_cs|`bool`|
|allowed_vp_cs|`list<item: utf8, nullable>`|
|authorize_time|`timestamp[us, tz=UTC]`|
|cluster_identifier (PK)|`utf8`|
|cluster_status|`utf8`|
|endpoint_count|`int64`|
|grantee (PK)|`utf8`|
|grantor (PK)|`utf8`|
|status|`utf8`|