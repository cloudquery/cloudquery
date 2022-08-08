
# Table: aws_redshift_cluster_parameters
Describes a parameter in a cluster parameter group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_parameter_group_cq_id|uuid|Unique CloudQuery ID of aws_redshift_cluster_parameter_groups table (FK)|
|allowed_values|text|The valid range of values for the parameter.|
|apply_type|text|Specifies how to apply the WLM configuration parameter|
|data_type|text|The data type of the parameter.|
|description|text|A description of the parameter.|
|is_modifiable|boolean|If true, the parameter can be modified|
|minimum_engine_version|text|The earliest engine version to which the parameter can apply.|
|parameter_name|text|The name of the parameter.|
|parameter_value|text|The value of the parameter|
|source|text|The source of the parameter value, such as "engine-default" or "user".|
