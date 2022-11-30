# Table: datadog_downtimes



The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_name|String|
|id (PK)|Int|
|active|Bool|
|active_child|JSON|
|canceled|JSON|
|creator_id|Int|
|disabled|Bool|
|downtime_type|Int|
|end|JSON|
|message|String|
|monitor_id|JSON|
|monitor_tags|StringArray|
|mute_first_recovery_notification|Bool|
|parent_id|JSON|
|recurrence|JSON|
|scope|StringArray|
|start|Int|
|timezone|String|
|updater_id|JSON|