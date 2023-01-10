# Table: aws_docdb_pending_maintenance_actions

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_PendingMaintenanceAction.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|pending_maintenance_action_details|JSON|
|resource_identifier|String|