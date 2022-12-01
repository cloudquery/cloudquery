# Table: datadog_incidents



The composite primary key for this table is (**account_name**, **id**).

## Relations

The following tables depend on datadog_incidents:
  - [datadog_incident_attachments](datadog_incident_attachments.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_name (PK)|String|
|id (PK)|String|
|attributes|JSON|
|relationships|JSON|
|type|String|