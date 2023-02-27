# Table: aws_elasticache_update_actions

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_UpdateAction.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|cache_cluster_id|String|
|cache_node_update_status|JSON|
|engine|String|
|estimated_update_time|String|
|node_group_update_status|JSON|
|nodes_updated|String|
|replication_group_id|String|
|service_update_name|String|
|service_update_recommended_apply_by_date|Timestamp|
|service_update_release_date|Timestamp|
|service_update_severity|String|
|service_update_status|String|
|service_update_type|String|
|sla_met|String|
|update_action_available_date|Timestamp|
|update_action_status|String|
|update_action_status_modified_date|Timestamp|