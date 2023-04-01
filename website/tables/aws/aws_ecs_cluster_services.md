# Table: aws_ecs_cluster_services

This table shows data for Amazon Elastic Container Service (ECS) Cluster Services.

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Service.html

The composite primary key for this table is (**arn**, **cluster_arn**).

## Relations

This table depends on [aws_ecs_clusters](aws_ecs_clusters).

The following tables depend on aws_ecs_cluster_services:
  - [aws_ecs_cluster_task_sets](aws_ecs_cluster_task_sets)

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
|capacity_provider_strategy|JSON|
|cluster_arn (PK)|String|
|created_at|Timestamp|
|created_by|String|
|deployment_configuration|JSON|
|deployment_controller|JSON|
|deployments|JSON|
|desired_count|Int|
|enable_ecs_managed_tags|Bool|
|enable_execute_command|Bool|
|events|JSON|
|health_check_grace_period_seconds|Int|
|launch_type|String|
|load_balancers|JSON|
|network_configuration|JSON|
|pending_count|Int|
|placement_constraints|JSON|
|placement_strategy|JSON|
|platform_family|String|
|platform_version|String|
|propagate_tags|String|
|role_arn|String|
|running_count|Int|
|scheduling_strategy|String|
|service_arn|String|
|service_name|String|
|service_registries|JSON|
|status|String|
|task_definition|String|
|task_sets|JSON|