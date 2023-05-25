# Table: aws_redshift_cluster_parameter_groups

This table shows data for Redshift Cluster Parameter Groups.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterParameterGroupStatus.html

The composite primary key for this table is (**cluster_arn**, **parameter_group_name**).

## Relations

This table depends on [aws_redshift_clusters](aws_redshift_clusters).

The following tables depend on aws_redshift_cluster_parameter_groups:
  - [aws_redshift_cluster_parameters](aws_redshift_cluster_parameters)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|cluster_arn (PK)|utf8|
|parameter_group_name (PK)|utf8|
|cluster_parameter_status_list|json|
|parameter_apply_status|utf8|