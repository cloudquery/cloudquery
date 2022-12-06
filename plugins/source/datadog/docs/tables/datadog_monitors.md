# Table: datadog_monitors



The composite primary key for this table is (**account_name**, **id**).

## Relations

The following tables depend on datadog_monitors:
  - [datadog_monitor_downtimes](datadog_monitor_downtimes.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_name (PK)|String|
|id (PK)|Int|
|deleted|Timestamp|
|priority|Int|
|created|Timestamp|
|creator|JSON|
|message|String|
|modified|Timestamp|
|multi|Bool|
|name|String|
|options|JSON|
|overall_state|String|
|query|String|
|restricted_roles|StringArray|
|state|JSON|
|tags|StringArray|
|type|String|