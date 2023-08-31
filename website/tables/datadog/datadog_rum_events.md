# Table: datadog_rum_events

This table shows data for Datadog Real User Monitoring (RUM) Events.

The composite primary key for this table is (**account_name**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|attributes|`json`|
|id (PK)|`utf8`|
|type|`utf8`|
|additional_properties|`json`|