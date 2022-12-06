# Table: aws_elasticbeanstalk_environments

https://docs.aws.amazon.com/elasticbeanstalk/latest/APIReference/API_EnvironmentDescription.html

The composite primary key for this table is (**account_id**, **id**).

## Relations

The following tables depend on aws_elasticbeanstalk_environments:
  - [aws_elasticbeanstalk_configuration_settings](aws_elasticbeanstalk_configuration_settings.md)
  - [aws_elasticbeanstalk_configuration_options](aws_elasticbeanstalk_configuration_options.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|arn|String|
|region|String|
|tags|JSON|
|id (PK)|String|
|listeners|JSON|
|abortable_operation_in_progress|Bool|
|application_name|String|
|cname|String|
|date_created|Timestamp|
|date_updated|Timestamp|
|description|String|
|endpoint_url|String|
|environment_links|JSON|
|environment_name|String|
|health|String|
|health_status|String|
|operations_role|String|
|platform_arn|String|
|resources|JSON|
|solution_stack_name|String|
|status|String|
|template_name|String|
|tier|JSON|
|version_label|String|