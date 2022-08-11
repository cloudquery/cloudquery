
# Table: aws_ecs_cluster_tasks
Details on a task in a cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_ecs_clusters table (FK)|
|attributes|jsonb|The attributes of the task|
|availability_zone|text|The Availability Zone for the task.|
|capacity_provider_name|text|The capacity provider that's associated with the task.|
|cluster_arn|text|The ARN of the cluster that hosts the task.|
|connectivity|text|The connectivity status of a task.|
|connectivity_at|timestamp without time zone|The Unix timestamp for the time when the task last went into CONNECTED status.|
|container_instance_arn|text|The ARN of the container instances that host the task.|
|cpu|text|The number of CPU units used by the task as expressed in a task definition|
|created_at|timestamp without time zone|The Unix timestamp for the time when the task was created|
|desired_status|text|The desired status of the task|
|enable_execute_command|boolean|Determines whether execute command functionality is enabled for this task|
|ephemeral_storage_size_in_gib|integer|The total amount, in GiB, of ephemeral storage to set for the task|
|execution_stopped_at|timestamp without time zone|The Unix timestamp for the time when the task execution stopped.|
|group|text|The name of the task group that's associated with the task.|
|health_status|text|The health status for the task|
|inference_accelerators|jsonb|The Elastic Inference accelerator that's associated with the task.|
|last_status|text|The last known status for the task|
|launch_type|text|The infrastructure where your task runs on|
|memory|text|The amount of memory (in MiB) that the task uses as expressed in a task definition|
|overrides|jsonb|One or more container overrides.|
|platform_family|text|The operating system that your tasks are running on|
|platform_version|text|The platform version where your task runs on|
|pull_started_at|timestamp without time zone|The Unix timestamp for the time when the container image pull began.|
|pull_stopped_at|timestamp without time zone|The Unix timestamp for the time when the container image pull completed.|
|started_at|timestamp without time zone|The Unix timestamp for the time when the task started|
|started_by|text|The tag specified when a task is started|
|stop_code|text|The stop code indicating why a task was stopped|
|stopped_at|timestamp without time zone|The Unix timestamp for the time when the task was stopped|
|stopped_reason|text|The reason that the task was stopped.|
|stopping_at|timestamp without time zone|The Unix timestamp for the time when the task stops|
|tags|jsonb|The metadata that you apply to the task to help you categorize and organize the task|
|arn|text|The Amazon Resource Name (ARN) of the task.|
|task_definition_arn|text|The ARN of the task definition that creates the task.|
|version|bigint|The version counter for the task|
