
# Table: aws_ecs_cluster_services
Details on a service within a cluster
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_ecs_clusters table (FK)|
|capacity_provider_strategy|jsonb|The capacity provider strategy the service uses|
|cluster_arn|text|The Amazon Resource Name (ARN) of the cluster that hosts the service.|
|created_at|timestamp without time zone|The Unix timestamp for the time when the service was created.|
|created_by|text|The principal that created the service.|
|deployment_configuration_deployment_circuit_breaker_enable|boolean|Determines whether to use the deployment circuit breaker logic for the service.  This member is required.|
|deployment_configuration_deployment_circuit_breaker_rollback|boolean|Determines whether to configure Amazon ECS to roll back the service if a service deployment fails|
|deployment_configuration_maximum_percent|integer|If a service is using the rolling update (ECS) deployment type, the maximum percent parameter represents an upper limit on the number of tasks in a service that are allowed in the RUNNING or PENDING state during a deployment, as a percentage of the desired number of tasks (rounded down to the nearest integer), and while any container instances are in the DRAINING state if the service contains tasks using the EC2 launch type|
|deployment_configuration_minimum_healthy_percent|integer|If a service is using the rolling update (ECS) deployment type, the minimum healthy percent represents a lower limit on the number of tasks in a service that must remain in the RUNNING state during a deployment, as a percentage of the desired number of tasks (rounded up to the nearest integer), and while any container instances are in the DRAINING state if the service contains tasks using the EC2 launch type|
|deployment_controller_type|text|The deployment controller type to use|
|desired_count|integer|The desired number of instantiations of the task definition to keep running on the service|
|enable_ecs_managed_tags|boolean|Determines whether to use Amazon ECS managed tags for the tasks in the service. For more information, see Tagging Your Amazon ECS Resources (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the Amazon Elastic Container Service Developer Guide.|
|enable_execute_command|boolean|Determines whether the execute command functionality is enabled for the service. If true, the execute command functionality is enabled for all containers in tasks as part of the service.|
|health_check_grace_period_seconds|integer|The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.|
|launch_type|text|The launch type the service is using|
|network_configuration_awsvpc_configuration_subnets|text[]|The IDs of the subnets associated with the task or service|
|network_configuration_awsvpc_configuration_assign_public_ip|text|Whether the task's elastic network interface receives a public IP address|
|network_configuration_awsvpc_configuration_security_groups|text[]|The IDs of the security groups associated with the task or service|
|pending_count|integer|The number of tasks in the cluster that are in the PENDING state.|
|placement_constraints|jsonb|The placement constraints for the tasks in the service.|
|placement_strategy|jsonb|The placement strategy that determines how tasks for the service are placed.|
|platform_family|text|The operating system that your tasks in the service run on|
|platform_version|text|The platform version to run your service on|
|propagate_tags|text|Determines whether to propagate the tags from the task definition or the service to the task|
|role_arn|text|The ARN of the IAM role that's associated with the service|
|running_count|integer|The number of tasks in the cluster that are in the RUNNING state.|
|scheduling_strategy|text|The scheduling strategy to use for the service|
|arn|text|The ARN that identifies the service|
|name|text|The name of your service|
|status|text|The status of the service|
|tags|jsonb|The metadata that you apply to the service to help you categorize and organize them|
|task_definition|text|The task definition to use for tasks in the service|
