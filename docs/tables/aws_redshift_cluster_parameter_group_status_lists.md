
# Table: aws_redshift_cluster_parameter_group_status_lists
Describes the status of a parameter group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_parameter_group_cq_id|uuid|Unique CloudQuery ID of aws_redshift_cluster_parameter_groups table (FK)|
|parameter_apply_error_description|text|The error that prevented the parameter from being applied to the database.|
|parameter_apply_status|text|The status of the parameter that indicates whether the parameter is in sync with the database, waiting for a cluster reboot, or encountered an error when being applied.|
|parameter_name|text|The name of the parameter.|
