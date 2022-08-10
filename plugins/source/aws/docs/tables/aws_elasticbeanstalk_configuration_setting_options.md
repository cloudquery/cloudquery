
# Table: aws_elasticbeanstalk_configuration_setting_options
A specification identifying an individual configuration option along with its current value
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|configuration_setting_cq_id|uuid|Unique CloudQuery ID of aws_elasticbeanstalk_configuration_setting_options table (FK)|
|namespace|text|A unique namespace that identifies the option's associated AWS resource.|
|option_name|text|The name of the configuration option.|
|resource_name|text|A unique resource name for the option setting|
|value|text|The current value for the configuration option.|
