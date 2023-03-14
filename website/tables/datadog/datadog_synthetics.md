# Table: datadog_synthetics

This table shows data for Datadog Synthetics.

The composite primary key for this table is (**account_name**, **public_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_name (PK)|String|
|public_id (PK)|String|
|config|JSON|
|creator|JSON|
|locations|StringArray|
|message|String|
|monitor_id|Int|
|name|String|
|options|JSON|
|status|String|
|steps|JSON|
|subtype|String|
|tags|StringArray|
|type|String|
|additional_properties|JSON|