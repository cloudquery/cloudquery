
# Table: aws_ecs_task_definitions
The details of a task definition which describes the container and volume definitions of an Amazon Elastic Container Service task
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|The metadata that you apply to the service to help you categorize and organize them|
|compatibilities|text[]|The task launch types the task definition validated against during task definition registration|
|cpu|text|The number of cpu units used by the task|
|deregistered_at|timestamp without time zone|The Unix timestamp for when the task definition was deregistered.|
|execution_role_arn|text|The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf|
|family|text|The name of a family that this task definition is registered to|
|inference_accelerators|jsonb|The Elastic Inference accelerator associated with the task.|
|ipc_mode|text|The IPC resource namespace to use for the containers in the task|
|memory|text|The amount (in MiB) of memory used by the task|
|network_mode|text|The Docker networking mode to use for the containers in the task|
|pid_mode|text|The process namespace to use for the containers in the task|
|placement_constraints|jsonb|An array of placement constraint objects to use for tasks|
|proxy_configuration_container_name|text|The name of the container that will serve as the App Mesh proxy.|
|proxy_configuration_properties|jsonb|The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified as key-value pairs.  * IgnoredUID - (Required) The user ID (UID) of the proxy container as defined by the user parameter in a container definition|
|proxy_configuration_type|text|The proxy type|
|registered_at|timestamp without time zone|The Unix timestamp for when the task definition was registered.|
|registered_by|text|The principal that registered the task definition.|
|requires_attributes|jsonb|The container instance attributes required by your task|
|requires_compatibilities|text[]|The task launch types the task definition was validated against|
|revision|integer|The revision of the task in a particular family|
|status|text|The status of the task definition.|
|arn|text|The full Amazon Resource Name (ARN) of the task definition.|
|task_role_arn|text|The short name or full Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that grants containers in the task permission to call AWS APIs on your behalf|
