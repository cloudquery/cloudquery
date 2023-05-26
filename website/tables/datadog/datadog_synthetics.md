# Table: datadog_synthetics

This table shows data for Datadog Synthetics.

The composite primary key for this table is (**account_name**, **public_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_name (PK)|utf8|
|public_id (PK)|utf8|
|config|json|
|creator|json|
|locations|list<item: utf8, nullable>|
|message|utf8|
|monitor_id|int64|
|name|utf8|
|options|json|
|status|utf8|
|steps|json|
|subtype|utf8|
|tags|list<item: utf8, nullable>|
|type|utf8|
|additional_properties|json|