# Table: aws_ecs_cluster_services

This table shows data for Amazon Elastic Container Service (ECS) Cluster Services.

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Service.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**arn**, **cluster_arn**).
## Relations

This table depends on [aws_ecs_clusters](aws_ecs_clusters.md).

The following tables depend on aws_ecs_cluster_services:
  - [aws_ecs_cluster_task_sets](aws_ecs_cluster_task_sets.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|capacity_provider_strategy|`json`|
|cluster_arn|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|created_by|`utf8`|
|deployment_configuration|`json`|
|deployment_controller|`json`|
|deployments|`json`|
|desired_count|`int64`|
|enable_ecs_managed_tags|`bool`|
|enable_execute_command|`bool`|
|events|`json`|
|health_check_grace_period_seconds|`int64`|
|launch_type|`utf8`|
|load_balancers|`json`|
|network_configuration|`json`|
|pending_count|`int64`|
|placement_constraints|`json`|
|placement_strategy|`json`|
|platform_family|`utf8`|
|platform_version|`utf8`|
|propagate_tags|`utf8`|
|role_arn|`utf8`|
|running_count|`int64`|
|scheduling_strategy|`utf8`|
|service_arn|`utf8`|
|service_name|`utf8`|
|service_registries|`json`|
|status|`utf8`|
|task_definition|`utf8`|
|task_sets|`json`|