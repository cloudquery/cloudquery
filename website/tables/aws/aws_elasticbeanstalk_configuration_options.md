# Table: aws_elasticbeanstalk_configuration_options

This table shows data for AWS Elastic Beanstalk Configuration Options.

https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationOptionDescription.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_elasticbeanstalk_environments](aws_elasticbeanstalk_environments).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|environment_id|`utf8`|
|change_severity|`utf8`|
|default_value|`utf8`|
|max_length|`int64`|
|max_value|`int64`|
|min_value|`int64`|
|name|`utf8`|
|namespace|`utf8`|
|regex|`json`|
|user_defined|`bool`|
|value_options|`list<item: utf8, nullable>`|
|value_type|`utf8`|
|application_arn|`utf8`|