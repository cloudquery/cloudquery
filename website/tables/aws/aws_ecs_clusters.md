# Table: aws_ecs_clusters

This table shows data for Amazon Elastic Container Service (ECS) Clusters.

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Cluster.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ecs_clusters:
  - [aws_ecs_cluster_container_instances](aws_ecs_cluster_container_instances)
  - [aws_ecs_cluster_services](aws_ecs_cluster_services)
  - [aws_ecs_cluster_tasks](aws_ecs_cluster_tasks)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|active_services_count|`int64`|
|attachments|`json`|
|attachments_status|`utf8`|
|capacity_providers|`list<item: utf8, nullable>`|
|cluster_arn|`utf8`|
|cluster_name|`utf8`|
|configuration|`json`|
|default_capacity_provider_strategy|`json`|
|pending_tasks_count|`int64`|
|registered_container_instances_count|`int64`|
|running_tasks_count|`int64`|
|service_connect_defaults|`json`|
|settings|`json`|
|statistics|`json`|
|status|`utf8`|