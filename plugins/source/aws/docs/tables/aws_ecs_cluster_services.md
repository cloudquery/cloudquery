# Table: aws_ecs_cluster_services

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Service.html

The primary key for this table is **arn**.

## Relations
This table depends on [aws_ecs_clusters](aws_ecs_clusters.md).


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
|capacity_provider_strategy|JSON|
|cluster_arn|String|
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
|service_name|String|
|service_registries|JSON|
|status|String|
|tags|JSON|
|task_definition|String|
|task_sets|JSON|