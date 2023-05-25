# Table: datadog_incident_attachments

This table shows data for Datadog Incident Attachments.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [datadog_incidents](datadog_incidents).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id (PK)|uuid|
|_cq_parent_id|uuid|
|account_name|utf8|
|attributes|extension_type<storage=binary>|
|id|utf8|
|relationships|extension_type<storage=binary>|
|type|utf8|