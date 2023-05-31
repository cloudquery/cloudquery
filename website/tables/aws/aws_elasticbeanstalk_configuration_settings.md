# Table: aws_elasticbeanstalk_configuration_settings

This table shows data for AWS Elastic Beanstalk Configuration Settings.

https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationSettingsDescription.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_elasticbeanstalk_environments](aws_elasticbeanstalk_environments).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|environment_id|`utf8`|
|application_name|`utf8`|
|date_created|`timestamp[us, tz=UTC]`|
|date_updated|`timestamp[us, tz=UTC]`|
|deployment_status|`utf8`|
|description|`utf8`|
|environment_name|`utf8`|
|option_settings|`json`|
|platform_arn|`utf8`|
|solution_stack_name|`utf8`|
|template_name|`utf8`|
|application_arn|`utf8`|