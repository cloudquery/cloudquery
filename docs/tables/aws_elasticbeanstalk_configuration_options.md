
# Table: aws_elasticbeanstalk_configuration_options
Describes the possible values for a configuration option.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|environment_cq_id|uuid|Unique CloudQuery ID of aws_elasticbeanstalk_environments table (FK)|
|application_arn|text|The arn of the associated application.|
|name|text|The name of the configuration option.|
|namespace|text|A unique namespace identifying the option's associated AWS resource.|
|change_severity|text|An indication of which action is required if the value for this configuration option changes: 						* NoInterruption : There is no interruption to the environment or application availability. 						* RestartEnvironment : The environment is entirely restarted, all AWS resources are deleted and recreated, and the environment is unavailable during the process. 						* RestartApplicationServer : The environment is available the entire time|
|default_value|text|The default value for this configuration option.|
|max_length|integer|If specified, the configuration option must be a string value no longer than this value.|
|max_value|integer|If specified, the configuration option must be a numeric value less than this value.|
|min_value|integer|If specified, the configuration option must be a numeric value greater than this value.|
|regex_label|text|A unique name representing this regular expression.|
|regex_pattern|text|The regular expression pattern that a string configuration option value with this restriction must match.|
|user_defined|boolean|An indication of whether the user defined this configuration option:  * true : This configuration option was defined by the user|
|value_options|text[]|If specified, values for the configuration option are selected from this list.|
|value_type|text|An indication of which type of values this option has and whether it is allowable to select one or more than one of the possible values:  * Scalar : Values for this option are a single selection from the possible values, or an unformatted string, or numeric value governed by the MIN/MAX/Regex constraints.  * List : Values for this option are multiple selections from the possible values.  * Boolean : Values for this option are either true or false .  * Json : Values for this option are a JSON representation of a ConfigDocument.|
