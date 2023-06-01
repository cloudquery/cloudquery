# Table: datadog_incidents

This table shows data for Datadog Incidents.

The primary key for this table is **id**.

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
|account_name|`utf8`|
|attributes|`json`|
|id (PK)|`utf8`|
|relationships|`json`|
|type|`utf8`|
|additional_properties|`json`|