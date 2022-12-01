# Table: aws_elasticbeanstalk_configuration_settings

https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationSettingsDescription.html

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
|application_name|String|
|date_created|Timestamp|
|date_updated|Timestamp|
|deployment_status|String|
|description|String|
|environment_name|String|
|option_settings|JSON|
|platform_arn|String|
|solution_stack_name|String|
|template_name|String|
|application_arn|String|