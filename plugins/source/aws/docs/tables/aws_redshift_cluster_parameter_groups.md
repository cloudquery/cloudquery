# Table: aws_redshift_cluster_parameter_groups

This table shows data for Redshift Cluster Parameter Groups.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterParameterGroupStatus.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**cluster_arn**, **parameter_group_name**).
## Relations

This table depends on [aws_redshift_clusters](aws_redshift_clusters.md).

The following tables depend on aws_redshift_cluster_parameter_groups:
  - [aws_redshift_cluster_parameters](aws_redshift_cluster_parameters.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn|`utf8`|
|cluster_parameter_status_list|`json`|
|parameter_apply_status|`utf8`|
|parameter_group_name|`utf8`|