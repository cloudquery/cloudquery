
# Table: aws_elasticache_parameter_group_parameters
Describes an individual setting that controls some aspect of ElastiCache behavior.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|parameter_group_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_parameter_groups table (FK)|
|allowed_values|text|The valid range of values for the parameter.|
|change_type|text|Indicates whether a change to the parameter is applied immediately or requires a reboot for the change to be applied|
|data_type|text|The valid data type for the parameter.|
|description|text|A description of the parameter.|
|is_modifiable|boolean|Indicates whether (true) or not (false) the parameter can be modified|
|minimum_engine_version|text|The earliest cache engine version to which the parameter can apply.|
|parameter_name|text|The name of the parameter.|
|parameter_value|text|The value of the parameter.|
|source|text|The source of the parameter.|
