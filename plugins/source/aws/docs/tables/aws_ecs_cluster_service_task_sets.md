
# Table: aws_ecs_cluster_service_task_sets
Information about a set of Amazon ECS tasks in either an CodeDeploy or an EXTERNAL deployment
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_service_cq_id|uuid|Unique CloudQuery ID of aws_ecs_cluster_services table (FK)|
|capacity_provider_strategy|jsonb|The capacity provider strategy that are associated with the task set.|
|cluster_arn|text|The Amazon Resource Name (ARN) of the cluster that the service that hosts the task set exists in.|
|computed_desired_count|integer|The computed desired count for the task set|
|created_at|timestamp without time zone|The Unix timestamp for the time when the task set was created.|
|external_id|text|The external ID associated with the task set|
|id|text|The ID of the task set.|
|launch_type|text|The launch type the tasks in the task set are using|
|network_configuration_awsvpc_configuration_subnets|text[]|The IDs of the subnets associated with the task or service|
|network_configuration_awsvpc_configuration_assign_public_ip|text|Whether the task's elastic network interface receives a public IP address|
|network_configuration_awsvpc_configuration_security_groups|text[]|The IDs of the security groups associated with the task or service|
|pending_count|integer|The number of tasks in the task set that are in the PENDING status during a deployment|
|platform_family|text|The operating system that your tasks in the set are running on|
|platform_version|text|The Fargate platform version where the tasks in the task set are running|
|running_count|integer|The number of tasks in the task set that are in the RUNNING status during a deployment|
|scale_unit|text|The unit of measure for the scale value.|
|scale_value|float|The value, specified as a percent total of a service's desiredCount, to scale the task set|
|service_arn|text|The Amazon Resource Name (ARN) of the service the task set exists in.|
|stability_status|text|The stability status|
|stability_status_at|timestamp without time zone|The Unix timestamp for the time when the task set stability status was retrieved.|
|started_by|text|The tag specified when a task set is started|
|status|text|The status of the task set|
|tags|jsonb|The metadata that you apply to the task set to help you categorize and organize them|
|task_definition|text|The task definition that the task set is using.|
|arn|text|The Amazon Resource Name (ARN) of the task set.|
|updated_at|timestamp without time zone|The Unix timestamp for the time when the task set was last updated.|
