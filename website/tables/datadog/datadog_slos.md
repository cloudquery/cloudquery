# Table: datadog_slos

This table shows data for Datadog Slos.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name|`utf8`|
|created_at|`int64`|
|creator|`json`|
|description|`utf8`|
|groups|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|modified_at|`int64`|
|monitor_ids|`list<item: int64, nullable>`|
|monitor_tags|`list<item: utf8, nullable>`|
|name|`utf8`|
|query|`json`|
|tags|`list<item: utf8, nullable>`|
|target_threshold|`float64`|
|thresholds|`json`|
|timeframe|`utf8`|
|type|`utf8`|
|warning_threshold|`float64`|
|additional_properties|`json`|