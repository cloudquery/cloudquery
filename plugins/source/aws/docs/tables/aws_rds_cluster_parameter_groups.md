# Table: aws_rds_cluster_parameter_groups


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_rds_cluster_parameter_groups`:
  - [`aws_rds_cluster_parameter_group_parameters`](aws_rds_cluster_parameter_group_parameters.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|db_cluster_parameter_group_name|String|
|db_parameter_group_family|String|
|description|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|