# Table: aws_redshift_endpoint_authorization

This table shows data for Redshift Endpoint Authorization.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_EndpointAuthorization.html

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
|allowed_all_vp_cs|`bool`|
|allowed_vp_cs|`list<item: utf8, nullable>`|
|authorize_time|`timestamp[us, tz=UTC]`|
|cluster_identifier|`utf8`|
|cluster_status|`utf8`|
|endpoint_count|`int64`|
|grantee|`utf8`|
|grantor|`utf8`|
|status|`utf8`|