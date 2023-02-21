# Table: launchdarkly_auditlog_entries

https://apidocs.launchdarkly.com/tag/Audit-log#operation/getAuditLogEntries

The primary key for this table is **id**.
It supports incremental syncs based on the **date** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|date (Incremental Key)|Timestamp|
|id (PK)|String|
|account_id|String|
|accesses|JSON|
|kind|String|
|name|String|
|description|String|
|short_description|String|
|comment|String|
|subject|JSON|
|member|JSON|
|token|JSON|
|app|JSON|
|title_verb|String|
|title|String|
|target|JSON|
|parent|JSON|