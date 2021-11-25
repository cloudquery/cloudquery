
# Table: aws_ecs_clusters
A regional grouping of one or more container instances on which you can run task requests
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|active_services_count|integer|The number of services that are running on the cluster in an ACTIVE state|
|attachments_status|text|The status of the capacity providers associated with the cluster|
|capacity_providers|text[]|The capacity providers associated with the cluster.|
|arn|text|The Amazon Resource Name (ARN) that identifies the cluster|
|name|text|A user-generated string that you use to identify your cluster.|
|execute_config_kms_key_id|text|Specify an AWS Key Management Service key ID to encrypt the data between the local client and the container.|
|execute_config_logs_cloud_watch_encryption_enabled|boolean|Whether or not to enable encryption on the CloudWatch logs|
|execute_config_log_cloud_watch_log_group_name|text|The name of the CloudWatch log group to send logs to|
|execute_config_log_s3_bucket_name|text|The name of the S3 bucket to send logs to|
|execute_config_log_s3_encryption_enabled|boolean|Whether or not to enable encryption on the CloudWatch logs|
|execute_config_log_s3_key_prefix|text|An optional folder in the S3 bucket to place logs in.|
|execute_config_logging|text|The log setting to use for redirecting logs for your execute command results. The following log settings are available.  * NONE: The execute command session is not logged.  * DEFAULT: The awslogs configuration in the task definition is used|
|default_capacity_provider_strategy|jsonb|The default capacity provider strategy for the cluster|
|pending_tasks_count|integer|The number of tasks in the cluster that are in the PENDING state.|
|registered_container_instances_count|integer|The number of container instances registered into the cluster|
|running_tasks_count|integer|The number of tasks in the cluster that are in the RUNNING state.|
|settings|jsonb|The settings for the cluster|
|statistics|jsonb|Additional information about your clusters that are separated by launch type, including:  * runningEC2TasksCount  * RunningFargateTasksCount  * pendingEC2TasksCount  * pendingFargateTasksCount  * activeEC2ServiceCount  * activeFargateServiceCount  * drainingEC2ServiceCount  * drainingFargateServiceCount|
|status|text|The status of the cluster|
|tags|jsonb|The metadata that you apply to the cluster to help you categorize and organize them|
