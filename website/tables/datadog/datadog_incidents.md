# Table: datadog_incidents

This table shows data for Datadog Incidents.

The composite primary key for this table is (**account_name**, **id**).

## Relations

The following tables depend on datadog_incidents:
  - [datadog_incident_attachments](datadog_incident_attachments)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|id (PK)|`utf8`|
|attributes|`json`|
|relationships|`json`|
|type|`utf8`|
|additional_properties|`json`|