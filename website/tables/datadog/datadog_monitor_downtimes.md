# Table: datadog_monitor_downtimes

This table shows data for Datadog Monitor Downtimes.

The composite primary key for this table is (**account_name**, **monitor_id**, **id**).

## Relations

This table depends on [datadog_monitors](datadog_monitors).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|monitor_id (PK)|`int64`|
|active|`bool`|
|active_child|`json`|
|canceled|`int64`|
|creator_id|`int64`|
|disabled|`bool`|
|downtime_type|`int64`|
|end|`int64`|
|id (PK)|`int64`|
|message|`utf8`|
|monitor_tags|`list<item: utf8, nullable>`|
|mute_first_recovery_notification|`bool`|
|notify_end_states|`list<item: utf8, nullable>`|
|notify_end_types|`list<item: utf8, nullable>`|
|parent_id|`int64`|
|recurrence|`json`|
|scope|`list<item: utf8, nullable>`|
|start|`int64`|
|timezone|`utf8`|
|updater_id|`int64`|
|additional_properties|`json`|