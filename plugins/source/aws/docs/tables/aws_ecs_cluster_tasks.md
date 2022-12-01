# Table: aws_ecs_cluster_tasks

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Task.html

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
|tags|JSON|
|task_protection|JSON|
|attachments|JSON|
|attributes|JSON|
|availability_zone|String|
|capacity_provider_name|String|
|cluster_arn|String|
|connectivity|String|
|connectivity_at|Timestamp|
|container_instance_arn|String|
|containers|JSON|
|cpu|String|
|created_at|Timestamp|
|desired_status|String|
|enable_execute_command|Bool|
|ephemeral_storage|JSON|
|execution_stopped_at|Timestamp|
|group|String|
|health_status|String|
|inference_accelerators|JSON|
|last_status|String|
|launch_type|String|
|memory|String|
|overrides|JSON|
|platform_family|String|
|platform_version|String|
|pull_started_at|Timestamp|
|pull_stopped_at|Timestamp|
|started_at|Timestamp|
|started_by|String|
|stop_code|String|
|stopped_at|Timestamp|
|stopped_reason|String|
|stopping_at|Timestamp|
|task_definition_arn|String|
|version|Int|