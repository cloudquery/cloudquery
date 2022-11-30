# Table: datadog_synthetics



The primary key for this table is **public_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_name|String|
|public_id (PK)|String|
|config|JSON|
|locations|StringArray|
|message|String|
|monitor_id|Int|
|name|String|
|options|JSON|
|status|String|
|subtype|String|
|tags|StringArray|
|type|String|