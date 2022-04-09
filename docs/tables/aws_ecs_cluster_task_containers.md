
# Table: aws_ecs_cluster_task_containers
A Docker container that's part of a task.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_task_cq_id|uuid|Unique CloudQuery ID of aws_ecs_cluster_tasks table (FK)|
|container_arn|text|The Amazon Resource Name (ARN) of the container.|
|cpu|text|The number of CPU units set for the container|
|exit_code|integer|The exit code returned from the container.|
|gpu_ids|text[]|The IDs of each GPU assigned to the container.|
|health_status|text|The health status of the container|
|image|text|The image used for the container.|
|image_digest|text|The container image manifest digest|
|last_status|text|The last known status of the container.|
|managed_agents|jsonb|The details of any Amazon ECS managed agents associated with the container.|
|memory|text|The hard limit (in MiB) of memory set for the container.|
|memory_reservation|text|The soft limit (in MiB) of memory set for the container.|
|name|text|The name of the container.|
|network_bindings|jsonb|The network bindings associated with the container.|
|network_interfaces|jsonb|The network interfaces associated with the container.|
|reason|text|A short (255 max characters) human-readable string to provide additional details about a running or stopped container.|
|runtime_id|text|The ID of the Docker container.|
|task_arn|text|The ARN of the task.|
