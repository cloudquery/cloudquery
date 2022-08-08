
# Table: aws_lightsail_database_parameters
Describes the parameters of a database
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_databases table (FK)|
|allowed_values|text|Specifies the valid range of values for the parameter|
|apply_method|text|Indicates when parameter updates are applied|
|apply_type|text|Specifies the engine-specific parameter type|
|data_type|text|Specifies the valid data type for the parameter|
|description|text|Provides a description of the parameter|
|is_modifiable|boolean|A Boolean value indicating whether the parameter can be modified|
|name|text|Specifies the name of the parameter|
|value|text|Specifies the value of the parameter|
