# Table: aws_redshift_endpoint_access

This table shows data for Redshift Endpoint Access.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_EndpointAccess.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_redshift_clusters](aws_redshift_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|cluster_arn|String|
|address|String|
|cluster_identifier|String|
|endpoint_create_time|Timestamp|
|endpoint_name|String|
|endpoint_status|String|
|port|Int|
|resource_owner|String|
|subnet_group_name|String|
|vpc_endpoint|JSON|
|vpc_security_groups|JSON|