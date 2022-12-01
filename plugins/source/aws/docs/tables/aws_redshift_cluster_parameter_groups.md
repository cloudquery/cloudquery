# Table: aws_redshift_cluster_parameter_groups

https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterParameterGroupStatus.html

The composite primary key for this table is (**cluster_arn**, **parameter_group_name**).

## Relations
This table depends on [aws_redshift_clusters](aws_redshift_clusters.md).

The following tables depend on aws_redshift_cluster_parameter_groups:
  - [aws_redshift_cluster_parameters](aws_redshift_cluster_parameters.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|cluster_arn (PK)|String|
|parameter_group_name (PK)|String|
|cluster_parameter_status_list|JSON|
|parameter_apply_status|String|