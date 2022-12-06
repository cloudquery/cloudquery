# Table: aws_elasticbeanstalk_configuration_options

https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationOptionDescription.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_elasticbeanstalk_environments](aws_elasticbeanstalk_environments.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|environment_id|String|
|change_severity|String|
|default_value|String|
|max_length|Int|
|max_value|Int|
|min_value|Int|
|name|String|
|namespace|String|
|regex|JSON|
|user_defined|Bool|
|value_options|StringArray|
|value_type|String|
|application_arn|String|