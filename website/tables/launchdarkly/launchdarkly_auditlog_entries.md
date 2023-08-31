# Table: launchdarkly_auditlog_entries

This table shows data for LaunchDarkly Audit Log Entries.

https://apidocs.launchdarkly.com/tag/Audit-log#operation/getAuditLogEntries

The primary key for this table is **id**.
It supports incremental syncs based on the **date** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|date (Incremental Key)|`timestamp[us, tz=UTC]`|
|id (PK)|`utf8`|
|account_id|`utf8`|
|accesses|`json`|
|kind|`utf8`|
|name|`utf8`|
|description|`utf8`|
|short_description|`utf8`|
|comment|`utf8`|
|subject|`json`|
|member|`json`|
|token|`json`|
|app|`json`|
|title_verb|`utf8`|
|title|`utf8`|
|target|`json`|
|parent|`json`|