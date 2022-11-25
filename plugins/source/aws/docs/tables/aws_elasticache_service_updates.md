# Table: aws_elasticache_service_updates

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ServiceUpdate.html

The primary key for this table is **arn**.



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
|auto_update_after_recommended_apply_by_date|Bool|
|engine|String|
|engine_version|String|
|estimated_update_time|String|
|service_update_description|String|
|service_update_end_date|Timestamp|
|service_update_name|String|
|service_update_recommended_apply_by_date|Timestamp|
|service_update_release_date|Timestamp|
|service_update_severity|String|
|service_update_status|String|
|service_update_type|String|