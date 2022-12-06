# Table: gcp_compute_firewalls



The primary key for this table is **self_link**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|self_link (PK)|String|
|allowed|JSON|
|creation_timestamp|String|
|denied|JSON|
|description|String|
|destination_ranges|StringArray|
|direction|String|
|disabled|Bool|
|id|Int|
|kind|String|
|log_config|JSON|
|name|String|
|network|String|
|priority|Int|
|source_ranges|StringArray|
|source_service_accounts|StringArray|
|source_tags|StringArray|
|target_service_accounts|StringArray|
|target_tags|StringArray|