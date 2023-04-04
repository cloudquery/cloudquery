# Table: aws_redshift_endpoint_authorization

This table shows data for Redshift Endpoint Authorization.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_EndpointAuthorization.html

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
|allowed_all_vp_cs|Bool|
|allowed_vp_cs|StringArray|
|authorize_time|Timestamp|
|cluster_identifier|String|
|cluster_status|String|
|endpoint_count|Int|
|grantee|String|
|grantor|String|
|status|String|