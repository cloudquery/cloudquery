# Table: datadog_downtimes

This table shows data for Datadog Downtimes.

The composite primary key for this table is (**account_name**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|id (PK)|`int64`|
|active|`bool`|
|active_child|`json`|
|canceled|`json`|
|creator_id|`int64`|
|disabled|`bool`|
|downtime_type|`int64`|
|end|`json`|
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
|additional_properties|`json`|