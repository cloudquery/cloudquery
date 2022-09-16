
# Table: aws_rds_cluster_parameters
This data type is used as a request parameter in the ModifyDBParameterGroup and ResetDBParameterGroup actions
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_parameter_group_cq_id|uuid|Unique CloudQuery ID of aws_rds_cluster_parameter_groups table (FK)|
|allowed_values|text|Specifies the valid range of values for the parameter.|
|apply_method|text|Indicates when to apply parameter updates.|
|apply_type|text|Specifies the engine specific parameters type.|
|data_type|text|Specifies the valid data type for the parameter.|
|description|text|Provides a description of the parameter.|
|is_modifiable|boolean|Indicates whether (true) or not (false) the parameter can be modified|
|minimum_engine_version|text|The earliest engine version to which the parameter can apply.|
|parameter_name|text|Specifies the name of the parameter.|
|parameter_value|text|Specifies the value of the parameter.|
|source|text|Indicates the source of the parameter value.|
|supported_engine_modes|text[]|The valid DB engine modes.|
