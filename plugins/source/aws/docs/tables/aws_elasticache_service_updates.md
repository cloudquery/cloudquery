# Table: aws_elasticache_service_updates

This table shows data for Elasticache Service Updates.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ServiceUpdate.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|auto_update_after_recommended_apply_by_date|`bool`|
|engine|`utf8`|
|engine_version|`utf8`|
|estimated_update_time|`utf8`|
|service_update_description|`utf8`|
|service_update_end_date|`timestamp[us, tz=UTC]`|
|service_update_name|`utf8`|
|service_update_recommended_apply_by_date|`timestamp[us, tz=UTC]`|
|service_update_release_date|`timestamp[us, tz=UTC]`|
|service_update_severity|`utf8`|
|service_update_status|`utf8`|
|service_update_type|`utf8`|