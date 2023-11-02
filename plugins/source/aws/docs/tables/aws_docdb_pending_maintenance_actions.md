# Table: aws_docdb_pending_maintenance_actions

This table shows data for Amazon DocumentDB Pending Maintenance Actions.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_PendingMaintenanceAction.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|pending_maintenance_action_details|`json`|
|resource_identifier|`utf8`|