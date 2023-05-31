# Table: gcp_compute_firewalls

This table shows data for GCP Compute Firewalls.

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|allowed|`json`|
|creation_timestamp|`utf8`|
|denied|`json`|
|description|`utf8`|
|destination_ranges|`list<item: utf8, nullable>`|
|direction|`utf8`|
|disabled|`bool`|
|id|`int64`|
|kind|`utf8`|
|log_config|`json`|
|name|`utf8`|
|network|`utf8`|
|priority|`int64`|
|self_link (PK)|`utf8`|
|source_ranges|`list<item: utf8, nullable>`|
|source_service_accounts|`list<item: utf8, nullable>`|
|source_tags|`list<item: utf8, nullable>`|
|target_service_accounts|`list<item: utf8, nullable>`|
|target_tags|`list<item: utf8, nullable>`|