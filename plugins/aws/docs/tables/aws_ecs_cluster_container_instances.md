
# Table: aws_ecs_cluster_container_instances
An EC2 instance that's running the Amazon ECS agent and has been registered with a cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_ecs_clusters table (FK)|
|agent_connected|boolean|This parameter returns true if the agent is connected to Amazon ECS|
|agent_update_status|text|The status of the most recent agent update|
|capacity_provider_name|text|The capacity provider that's associated with the container instance.|
|container_instance_arn|text|The Amazon Resource Name (ARN) of the container instance|
|ec2_instance_id|text|The ID of the container instance|
|health_status_overall_status|text|The overall health status of the container instance|
|pending_tasks_count|integer|The number of tasks on the container instance that are in the PENDING status.|
|registered_at|timestamp without time zone|The Unix timestamp for the time when the container instance was registered.|
|running_tasks_count|integer|The number of tasks on the container instance that are in the RUNNING status.|
|status|text|The status of the container instance|
|status_reason|text|The reason that the container instance reached its current status.|
|tags|jsonb|The metadata that you apply to the container instance to help you categorize and organize them|
|version|bigint|The version counter for the container instance|
|version_info_agent_hash|text|The Git commit hash for the Amazon ECS container agent build on the amazon-ecs-agent  (https://github.com/aws/amazon-ecs-agent/commits/master) GitHub repository.|
|version_info_agent_version|text|The version number of the Amazon ECS container agent.|
|version_info_docker_version|text|The Docker version that's running on the container instance.|
