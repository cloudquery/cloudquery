# Table: datadog_monitors

This table shows data for Datadog Monitors.

The composite primary key for this table is (**account_name**, **id**).

## Relations

The following tables depend on datadog_monitors:
  - [datadog_monitor_downtimes](datadog_monitor_downtimes)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|creator|`json`|
|deleted|`timestamp[us, tz=UTC]`|
|id (PK)|`int64`|
|matching_downtimes|`json`|
|message|`utf8`|
|modified|`timestamp[us, tz=UTC]`|
|multi|`bool`|
|name|`utf8`|
|options|`json`|
|overall_state|`utf8`|
|priority|`int64`|
|query|`utf8`|
|restricted_roles|`list<item: utf8, nullable>`|
|state|`json`|
|tags|`list<item: utf8, nullable>`|
|type|`utf8`|
|additional_properties|`json`|