# Table: datadog_hosts



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
|aliases|StringArray|
|apps|StringArray|
|aws_name|String|
|host_name|String|
|is_muted|Bool|
|last_reported_time|Int|
|meta|JSON|
|metrics|JSON|
|mute_timeout|Int|
|name|String|
|sources|StringArray|
|tags_by_source|JSON|
|up|Bool|