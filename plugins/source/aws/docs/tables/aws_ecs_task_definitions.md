# Table: aws_ecs_task_definitions

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskDefinition.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|compatibilities|StringArray|
|container_definitions|JSON|
|cpu|String|
|deregistered_at|Timestamp|
|ephemeral_storage|JSON|
|execution_role_arn|String|
|family|String|
|inference_accelerators|JSON|
|ipc_mode|String|
|memory|String|
|network_mode|String|
|pid_mode|String|
|placement_constraints|JSON|
|proxy_configuration|JSON|
|registered_at|Timestamp|
|registered_by|String|
|requires_attributes|JSON|
|requires_compatibilities|StringArray|
|revision|Int|
|runtime_platform|JSON|
|status|String|
|task_role_arn|String|
|volumes|JSON|