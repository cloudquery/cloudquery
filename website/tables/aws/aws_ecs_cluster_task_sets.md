# Table: aws_ecs_cluster_task_sets

This table shows data for Amazon Elastic Container Service (ECS) Cluster Task Sets.

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskSet.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_ecs_cluster_services](aws_ecs_cluster_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|capacity_provider_strategy|`json`|
|cluster_arn|`utf8`|
|computed_desired_count|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|external_id|`utf8`|
|id|`utf8`|
|launch_type|`utf8`|
|load_balancers|`json`|
|network_configuration|`json`|
|pending_count|`int64`|
|platform_family|`utf8`|
|platform_version|`utf8`|
|running_count|`int64`|
|scale|`json`|
|service_arn|`utf8`|
|service_registries|`json`|
|stability_status|`utf8`|
|stability_status_at|`timestamp[us, tz=UTC]`|
|started_by|`utf8`|
|status|`utf8`|
|task_definition|`utf8`|
|task_set_arn|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|