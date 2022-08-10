
# Table: aws_elasticbeanstalk_configuration_settings
Describes the settings for a configuration set.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|environment_cq_id|uuid|Unique CloudQuery ID of aws_elasticbeanstalk_environments table (FK)|
|application_name|text|The name of the application associated with this configuration set.|
|application_arn|text|The arn of the associated application.|
|date_created|timestamp without time zone|The date (in UTC time) when this configuration set was created.|
|date_updated|timestamp without time zone|The date (in UTC time) when this configuration set was last modified.|
|deployment_status|text|If this configuration set is associated with an environment, the DeploymentStatus parameter indicates the deployment status of this configuration set:  * null: This configuration is not associated with a running environment.  * pending: This is a draft configuration that is not deployed to the associated environment but is in the process of deploying.  * deployed: This is the configuration that is currently deployed to the associated running environment.  * failed: This is a draft configuration that failed to successfully deploy.|
|description|text|Describes this configuration set.|
|environment_name|text|If not null, the name of the environment for this configuration set.|
|platform_arn|text|The ARN of the platform version.|
|solution_stack_name|text|The name of the solution stack this configuration set uses.|
|template_name|text|If not null, the name of the configuration template for this configuration set.|
