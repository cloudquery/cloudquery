# Table: aws_elasticbeanstalk_configuration_options

This table shows data for AWS Elastic Beanstalk Configuration Options.

https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationOptionDescription.html

The composite primary key for this table is (**environment_arn**, **name**, **application_arn**, **solution_stack_name**).

## Relations

This table depends on [aws_elasticbeanstalk_environments](aws_elasticbeanstalk_environments.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|environment_arn (PK)|`utf8`|
|change_severity|`utf8`|
|default_value|`utf8`|
|max_length|`int64`|
|max_value|`int64`|
|min_value|`int64`|
|name (PK)|`utf8`|
|namespace|`utf8`|
|regex|`json`|
|user_defined|`bool`|
|value_options|`list<item: utf8, nullable>`|
|value_type|`utf8`|
|application_arn (PK)|`utf8`|
|solution_stack_name (PK)|`utf8`|