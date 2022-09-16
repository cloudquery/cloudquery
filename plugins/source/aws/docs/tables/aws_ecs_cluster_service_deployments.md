
# Table: aws_ecs_cluster_service_deployments
The details of an Amazon ECS service deployment
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_service_cq_id|uuid|Unique CloudQuery ID of aws_ecs_cluster_services table (FK)|
|capacity_provider_strategy|jsonb|The capacity provider strategy that the deployment is using.|
|created_at|timestamp without time zone|The Unix timestamp for the time when the service deployment was created.|
|desired_count|integer|The most recent desired count of tasks that was specified for the service to deploy or maintain.|
|failed_tasks|integer|The number of consecutively failed tasks in the deployment|
|id|text|The ID of the deployment.|
|launch_type|text|The launch type the tasks in the service are using|
|network_configuration_awsvpc_configuration_subnets|text[]|The IDs of the subnets associated with the task or service|
|network_configuration_awsvpc_configuration_assign_public_ip|text|Whether the task's elastic network interface receives a public IP address|
|network_configuration_awsvpc_configuration_security_groups|text[]|The IDs of the security groups associated with the task or service|
|pending_count|integer|The number of tasks in the deployment that are in the PENDING status.|
|platform_family|text|The operating system that your tasks in the service, or tasks are running on|
|platform_version|text|The platform version that your tasks in the service run on|
|rollout_state|text|The rolloutState of a service is only returned for services that use the rolling update (ECS) deployment type that aren't behind a Classic Load Balancer|
|rollout_state_reason|text|A description of the rollout state of a deployment.|
|running_count|integer|The number of tasks in the deployment that are in the RUNNING status.|
|status|text|The status of the deployment|
|task_definition|text|The most recent task definition that was specified for the tasks in the service to use.|
|updated_at|timestamp without time zone|The Unix timestamp for the time when the service deployment was last updated.|
