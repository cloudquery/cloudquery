# Table: datadog_incident_attachments

This table shows data for Datadog Incident Attachments.

The composite primary key for this table is (**account_name**, **incident_id**, **id**).

## Relations

This table depends on [datadog_incidents](datadog_incidents).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|incident_id (PK)|`utf8`|
|attributes|`json`|
|id (PK)|`utf8`|
|relationships|`json`|
|type|`utf8`|
|additional_properties|`json`|