# Table: aws_ecs_cluster_container_instances

This table shows data for Amazon Elastic Container Service (ECS) Cluster Container Instances.

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerInstance.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ecs_clusters](aws_ecs_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn|`utf8`|
|tags|`json`|
|agent_connected|`bool`|
|agent_update_status|`utf8`|
|attachments|`json`|
|attributes|`json`|
|capacity_provider_name|`utf8`|
|container_instance_arn|`utf8`|
|ec2_instance_id|`utf8`|
|health_status|`json`|
|pending_tasks_count|`int64`|
|registered_at|`timestamp[us, tz=UTC]`|
|registered_resources|`json`|
|remaining_resources|`json`|
|running_tasks_count|`int64`|
|status|`utf8`|
|status_reason|`utf8`|
|version|`int64`|
|version_info|`json`|