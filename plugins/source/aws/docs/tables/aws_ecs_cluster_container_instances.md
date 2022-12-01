# Table: aws_ecs_cluster_container_instances

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerInstance.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_ecs_clusters](aws_ecs_clusters.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|cluster_arn|String|
|agent_connected|Bool|
|agent_update_status|String|
|attachments|JSON|
|attributes|JSON|
|capacity_provider_name|String|
|container_instance_arn|String|
|ec2_instance_id|String|
|health_status|JSON|
|pending_tasks_count|Int|
|registered_at|Timestamp|
|registered_resources|JSON|
|remaining_resources|JSON|
|running_tasks_count|Int|
|status|String|
|status_reason|String|
|tags|JSON|
|version|Int|
|version_info|JSON|