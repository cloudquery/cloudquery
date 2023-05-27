# Table: datadog_monitor_downtimes

This table shows data for Datadog Monitor Downtimes.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [datadog_monitors](datadog_monitors).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_name|`utf8`|
|active|`bool`|
|active_child|`json`|
|canceled|`json`|
|creator_id|`int64`|
|disabled|`bool`|
|downtime_type|`int64`|
|end|`json`|
|id|`int64`|
|message|`utf8`|
|monitor_id|`json`|
|monitor_tags|`list<item: utf8, nullable>`|
|mute_first_recovery_notification|`bool`|
|parent_id|`json`|
|recurrence|`json`|
|scope|`list<item: utf8, nullable>`|
|start|`int64`|
|timezone|`utf8`|
|updater_id|`json`|