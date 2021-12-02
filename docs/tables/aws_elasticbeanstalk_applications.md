
# Table: aws_elasticbeanstalk_applications
Describes the properties of an application.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the application.|
|name|text|The name of the application.|
|configuration_templates|text[]|The names of the configuration templates associated with this application.|
|date_created|timestamp without time zone|The date when the application was created.|
|date_updated|timestamp without time zone|The date when the application was last modified.|
|description|text|User-defined description of the application.|
|resource_lifecycle_config_service_role|text|The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a VersionLifecycleConfig for the application in one of the supporting calls (CreateApplication or UpdateApplicationResourceLifecycle)|
|max_age_rule_enabled|boolean|Specify true to apply the rule, or false to disable it.  This member is required.|
|max_age_rule_delete_source_from_s3|boolean|Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.|
|max_age_rule_max_age_in_days|integer|Specify the number of days to retain an application versions.|
|max_count_rule_enabled|boolean|Specify true to apply the rule, or false to disable it.  This member is required.|
|max_count_rule_delete_source_from_s3|boolean|Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.|
|max_count_rule_max_count|integer|Specify the maximum number of application versions to retain.|
|versions|text[]|The names of the versions for this application.|
