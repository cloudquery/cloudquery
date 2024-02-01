# Table: aws_ecs_task_definitions

This table shows data for Amazon Elastic Container Service (ECS) Task Definitions.

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskDefinition.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|compatibilities|`list<item: utf8, nullable>`|
|container_definitions|`json`|
|cpu|`utf8`|
|deregistered_at|`timestamp[us, tz=UTC]`|
|ephemeral_storage|`json`|
|execution_role_arn|`utf8`|
|family|`utf8`|
|inference_accelerators|`json`|
|ipc_mode|`utf8`|
|memory|`utf8`|
|network_mode|`utf8`|
|pid_mode|`utf8`|
|placement_constraints|`json`|
|proxy_configuration|`json`|
|registered_at|`timestamp[us, tz=UTC]`|
|registered_by|`utf8`|
|requires_attributes|`json`|
|requires_compatibilities|`list<item: utf8, nullable>`|
|revision|`int64`|
|runtime_platform|`json`|
|status|`utf8`|
|task_definition_arn|`utf8`|
|task_role_arn|`utf8`|
|volumes|`json`|