# Table: aws_ecs_clusters

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Cluster.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ecs_clusters:
  - [aws_ecs_cluster_tasks](aws_ecs_cluster_tasks.md)
  - [aws_ecs_cluster_services](aws_ecs_cluster_services.md)
  - [aws_ecs_cluster_container_instances](aws_ecs_cluster_container_instances.md)

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
|active_services_count|Int|
|attachments|JSON|
|attachments_status|String|
|capacity_providers|StringArray|
|cluster_name|String|
|configuration|JSON|
|default_capacity_provider_strategy|JSON|
|pending_tasks_count|Int|
|registered_container_instances_count|Int|
|running_tasks_count|Int|
|service_connect_defaults|JSON|
|settings|JSON|
|statistics|JSON|
|status|String|
|tags|JSON|