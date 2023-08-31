# Table: aws_ecs_cluster_tasks

This table shows data for Amazon Elastic Container Service (ECS) Cluster Tasks.

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Task.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_ecs_clusters](aws_ecs_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|task_protection|`json`|
|attachments|`json`|
|attributes|`json`|
|availability_zone|`utf8`|
|capacity_provider_name|`utf8`|
|cluster_arn|`utf8`|
|connectivity|`utf8`|
|connectivity_at|`timestamp[us, tz=UTC]`|
|container_instance_arn|`utf8`|
|containers|`json`|
|cpu|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|desired_status|`utf8`|
|enable_execute_command|`bool`|
|ephemeral_storage|`json`|
|execution_stopped_at|`timestamp[us, tz=UTC]`|
|group|`utf8`|
|health_status|`utf8`|
|inference_accelerators|`json`|
|last_status|`utf8`|
|launch_type|`utf8`|
|memory|`utf8`|
|overrides|`json`|
|platform_family|`utf8`|
|platform_version|`utf8`|
|pull_started_at|`timestamp[us, tz=UTC]`|
|pull_stopped_at|`timestamp[us, tz=UTC]`|
|started_at|`timestamp[us, tz=UTC]`|
|started_by|`utf8`|
|stop_code|`utf8`|
|stopped_at|`timestamp[us, tz=UTC]`|
|stopped_reason|`utf8`|
|stopping_at|`timestamp[us, tz=UTC]`|
|task_arn|`utf8`|
|task_definition_arn|`utf8`|
|version|`int64`|