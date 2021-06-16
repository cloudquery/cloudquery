
# Table: aws_redshift_cluster_parameter_groups
Describes the status of a parameter group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid|Unique ID of aws_redshift_clusters table (FK)|
|parameter_apply_status|text|The status of parameter updates.|
|parameter_group_name|text|The name of the cluster parameter group.|
